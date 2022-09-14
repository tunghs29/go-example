package controller

import (
	"Test/common"
	"Test/component"
	"Test/database"
	"Test/modules/product/models"
	"Test/modules/product/responsitory"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GET/v1/products

func ListProduct(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		collection := database.GetCollection(
			appCtx.GetMainDBConnection(),
			models.Product{}.GetProductDatabase(),
			models.Product{}.GetProductCollection(),
		)

		result, err := responsitory.
			NewListProductResp(collection).
			ListProduct(c.Request.Context())
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))

		fmt.Println("hello")
	}
}
