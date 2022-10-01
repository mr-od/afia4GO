package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/oddinnovate/a4go/db/sqlc"
	"github.com/oddinnovate/a4go/token"
	"github.com/oddinnovate/a4go/util"
	uuid "github.com/satori/go.uuid"
)

type ChatRoomRequest struct {
	Name string `json:"name"`
}

// type SeverC struct {
// 	*api.Server
// }

func (cs *Server) CreateRoom(ctx *gin.Context) {
	var req ChatRoomRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(AuthorizationPayloadKey).(*token.Payload)
	arg := db.CreateRoomParams{
		Name:     req.Name,
		Owner:    authPayload.Username,
		PublicID: uuid.NewV4().String(),
	}

	room, err := cs.Store.CreateRoom(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, util.ErrorResponse(err))
				return

			}
		}
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, room)

}

type listRoomsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1,max=10"`
}

func (cs *Server) ListRooms(ctx *gin.Context) {
	var req listRoomsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// authpayload := ctx.MustGet(AuthorizationPayloadKey).(*token.Payload)
	arg := db.ListRoomsParams{
		// Owner:  authpayload.Username,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	rooms, err := cs.Store.ListRooms(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, rooms)
}
