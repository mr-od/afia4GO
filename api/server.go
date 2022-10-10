package api

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/websocket"
	db "github.com/oddinnovate/a4go/db/sqlc"
	"github.com/oddinnovate/a4go/token"
	"github.com/oddinnovate/a4go/util"
)

type Server struct {
	Config     util.Config
	Store      db.Store
	TokenMaker token.Maker
	Router     *gin.Engine
	hub        *Hub
	gnats      *Gnats
	// chatService *
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		Config:     config,
		Store:      store,
		TokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", util.ValidCurrency)
	}
	server.setupRouter()

	return server, err
}

func (server *Server) setupRouter() {

	router := gin.Default()

	// User Route
	router.POST("/api/v1/users", server.createUser)
	router.POST("/api/v1/users/login", server.loginUser)
	router.POST("/api/v1/tokens/refresh_token", server.renewAccessToken)

	authRoutes := router.Group("/").Use(AuthMiddleware(server.TokenMaker))

	// Accounts Route
	authRoutes.POST("api/v1/accounts", server.createAccount)
	authRoutes.GET("api/v1/accounts/:id", server.getAccount)
	authRoutes.GET("api/v1/accounts", server.listAccounts)

	// Products Route
	authRoutes.POST("api/v1/products", server.addProduct)
	authRoutes.GET("api/v1/products/:id", server.getProduct)
	authRoutes.GET("api/v1/products", server.listProducts)

	// Order Route
	authRoutes.POST("api/v1/orders/place", server.placeOrder)

	// Transfers Route
	authRoutes.POST("api/v1/transfers", server.createTransfer)

	// Chat Route
	authRoutes.POST("api/v1/room", server.CreateRoomHandler)
	authRoutes.GET("api/v1/rooms", server.ListRoomsHandler)

	server.Router = router

}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// serveWs handles websocket requests from the peer.
func (s *Server) serveWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	u, err := authSocket(conn, s)
	if err != nil {
		if err1 := conn.WriteJSON(Message{
			Type: MessageTypeUnauthorizedErr,
			Body: ServerErrorMessage{
				Message: "Auth error occurred",
			},
		}); err != nil {
			log.Errorf("Got err while sending auth message: %v", err1)
		}
		if err2 := conn.Close(); err2 != nil {
			log.Errorf("Got err while closing socket during auth: %v", err2)
		}
		return
	}

	m := Message{
		Type: MessageTypeAuthAck,
		Body: "Authentication success",
	}
	if err := conn.WriteJSON(&m); err != nil {
		log.Errorf("Err while authack %v", err)
		if err2 := conn.Close(); err2 != nil {
			log.Errorf("Got err while closing socket during auth: %v", err2)
		}
		return
	}

	client := newClient(s.hub, s.gnats, conn, u)
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}

type AuthRequest struct {
	Token string `binding:"required"`
}

func authSocket(conn *websocket.Conn, s *Server) (*token.Payload, error) {
	if err := conn.SetReadDeadline(time.Now().Add(10 * time.Second)); err != nil {
		log.Errorf("Error occurred while setting write deadline during auth: %v", err)
		return nil, err
	}

	var r AuthRequest
	if err := conn.ReadJSON(&r); err != nil {
		return nil, err
	}

	j, err := s.TokenMaker.VerifyToken(r.Token)
	if err != nil {
		return nil, err
	}

	return j, nil
}
