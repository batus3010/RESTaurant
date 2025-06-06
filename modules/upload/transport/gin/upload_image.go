package uploadGin

import (
	"RESTaurant_v2/components/appctx"
	uploadBiz "RESTaurant_v2/modules/upload/biz"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadImage(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// use when client want to save images into a folder, otherwise it's just image
		folder := c.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		defer file.Close()

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//imgStore := uploadStorage.NewSQLStore(db)
		biz := uploadBiz.NewUploadBiz(appCtx.UploadProvider())
		img, err := biz.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": img})
	}
}
