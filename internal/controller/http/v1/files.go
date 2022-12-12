package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type filesRoutes struct {
}

func newFilesRoutes(handler *gin.Engine) {

	fl := filesRoutes{}

	handler.GET("/files", fl.getFileList)
	handler.POST("/files", fl.copyFileToService)
}

func (f *filesRoutes) getFileList(c *gin.Context) {
	c.JSON(http.StatusOK, "ok!")
}

func (f *filesRoutes) copyFileToService(c *gin.Context) {
	c.JSON(http.StatusOK, "ok!")
}
