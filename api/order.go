package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/oddinnovate/a4go/db/sqlc"
	"github.com/oddinnovate/a4go/token"
)

// type orderRequest struct {
// 	DeliveryFee int64  `json:"delivery_fee" binding:"required"`
// 	Subtotal    int64  `json:"subtotal" binding:"required"`
// 	Total       int64  `json:"total" binding:"required"`
// 	Status      string `json:"status" binding:"required"`
// }

// type orderItemsRequest struct {
// 	OrderID   int64  `json:"order_id" binding:"required"`
// 	ProductID int64  `json:"product_id" binding:"required"`
// 	Quantity  int64  `json:"quantity" binding:"required"`
// 	Status    string `json:"status" binding:"required"`
// }

// type placeOrderRequest struct {
// 	Order    orderRequest        `json:"order"`
// 	Products []orderItemsRequest `json:"products"`
// }

type placeOrderRequest struct {
	Orders   Order     `json:"order_details"`
	Products []Product `json:"aos"`
}

type Order struct {
	DeliveryFee int64 `json:"delivery_fee"`
	Subtotal    int64 `json:"subtotal"`
	Total       int64 `json:"total"`
	// Status      string `json:"status"`
}

type Product struct {
	OrderID   int64 `json:"order_id"`
	ProductID int64 `json:"product_id"`
	Quantity  int64 `json:"quantity"`
	// Status    string `json:"status"`
}

func (server *Server) placeOrder(ctx *gin.Context) {
	// var req Order
	var mr placeOrderRequest
	authPayload := ctx.MustGet(AuthorizationPayloadKey).(*token.Payload)

	if err := ctx.ShouldBindJSON(&mr); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateOrderParams{
		Owner:       authPayload.Username,
		Status:      "pending",
		DeliveryFee: mr.Orders.DeliveryFee,
		Subtotal:    mr.Orders.Subtotal,
		Total:       mr.Orders.Total,
	}

	mainOrder, err := server.Store.CreateOrder(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return

			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	for _, oi := range mr.Products {
		orderItems := Product{
			// OrderID:   mainOrder.ID,
			ProductID: oi.ProductID,
			Quantity:  oi.Quantity,
		}
		ctx.JSON(http.StatusCreated, orderItems)
		arg2 := db.CreateOrderItemParams{
			Owner:     authPayload.Username,
			Status:    "pending",
			OrderID:   mainOrder.ID,
			ProductID: oi.ProductID,
			Quantity:  oi.Quantity,
		}
		oip, err := server.Store.CreateOrderItem(ctx, arg2)
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok {
				switch pqErr.Code.Name() {
				case "foreign_key_violation", "unique_violation":
					ctx.JSON(http.StatusForbidden, errorResponse(err))
					return

				}
			}
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, oip)
	}

	ctx.JSON(http.StatusOK, mainOrder)

}
