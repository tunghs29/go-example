package controller

import (
	"Test/common"
	"Test/component"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
)

type FileJson struct {
	Title       string                `json:"name" binding:"required"`
	Description string                `json:"description"`
	fileData    *multipart.FileHeader `form:"file"`
}

func ImportMasterData(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			panic(err)
		}

		//folder := c.DefaultPostForm("folder", "img")
		file, err := fileHeader.Open()
		if err != nil {
			panic(err)
		}
		defer file.Close()

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(gin.H{"Hello": "Hello"}))
	}
}
