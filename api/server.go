package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/oddinnovate/a4go/db/sqlc"
	"github.com/oddinnovate/a4go/token"
	"github.com/oddinnovate/a4go/util"
)

type Server struct {
	Config     util.Config
	Store      db.Store
	TokenMaker token.Maker
	Router     *gin.Engine
	// chatService *chat.SeverC
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
	authRoutes.POST("api/v1/room", server.CreateRoom)
	authRoutes.GET("api/v1/rooms", server.ListRooms)

	server.Router = router

}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/oddinnovate/a4go/api/v1"
// 	db "github.com/oddinnovate/a4go/db/sqlc"
// 	"github.com/oddinnovate/a4go/token"
// 	"github.com/oddinnovate/a4go/util"
// )

// type ChatRoomDTO struct {
// 	ID   string `json:"id"`
// 	Name string `json:"name"`
// }

// type SeverC struct {
// 	*api.Server
// }

// func (se *SeverC) CreateChat(ctx *gin.Context) {
// 	var req ChatRoomDTO
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
// 		return
// 	}

// 	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
// 	arg := db.CreateMessageParams{}

// }
