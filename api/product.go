package api

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/oddinnovate/a4go/db/sqlc"
	"github.com/oddinnovate/a4go/token"
	"github.com/oddinnovate/a4go/util"
)

type createProductRequest struct {
	Name        string                  `form:"name"`
	Price       int64                   `form:"price"`
	Description string                  `form:"description"`
	Images      []*multipart.FileHeader `form:"images"`
}

func (server *Server) addProduct(ctx *gin.Context) {
	var cf = server.Config
	var req createProductRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Create an OSSClient instance.
	// Set yourEndpoint to the endpoint of the region in which the bucket is located. For example, if the bucket is located in the China (Hangzhou) region, set yourEndpoint to https://oss-cn-hangzhou.aliyuncs.com. Specify the endpoint based on your business requirements.
	// Security risks may arise if you use the AccessKey pair of an Alibaba Cloud account to access OSS because the account has permissions on all API operations. We recommend that you use a RAM user to call API operations or perform routine operations and maintenance. To create a RAM user, log on to the RAM console.
	client, err := oss.New(cf.OSSEndpoint, cf.OssAccessID, cf.OSSAccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// Specify the name of the bucket. Example: examplebucket.
	bucket, err := client.Bucket(cf.OSSBucket)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	var imgs_name []string
	var imgs_url []string

	for _, image := range req.Images {
		file, err := image.Open()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		imgs := "product-images/" + util.RandomProduct() + image.Filename
		imgs_name = append(imgs_name, imgs)

		imgsUrl := cf.OSSBucket + "." + cf.OSSEndpoint + "/" + imgs
		imgs_url = append(imgs_url, imgsUrl)

		log.Println(imgs_name)
		log.Println(imgs_url)

		// ctx.SaveUploadedFile(image, imgs)
		err = bucket.PutObject(imgs, file)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}

		defer file.Close()

	}

	authPayload := ctx.MustGet(AuthorizationPayloadKey).(*token.Payload)
	arg := db.CreateProductParams{
		Owner:       authPayload.Username,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Imgs:        imgs_name,
		ImgsUrl:     imgs_url,
	}

	product, err := server.Store.CreateProduct(ctx, arg)
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

	ctx.JSON(http.StatusOK, product)
}

type getProductRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getProduct(ctx *gin.Context) {
	var req getProductRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	product, err := server.Store.GetProduct(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(AuthorizationPayloadKey).(*token.Payload)
	if product.Owner != authPayload.Username {
		err := errors.New("product does not belong to the logged in user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, product)
}

type listProductsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1,max=10"`
}

func (server *Server) listProducts(ctx *gin.Context) {
	var req listProductsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(AuthorizationPayloadKey).(*token.Payload)
	arg := db.ListProductsParams{
		Owner:  authPayload.Username,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	products, err := server.Store.ListProducts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, products)
}
