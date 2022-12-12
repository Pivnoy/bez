package v1

import (
	"bez/internal/usecase"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, googleAPI usecase.GoogleAPI, driveAPI usecase.DriveAPI, us usecase.User, sr usecase.Service, fl usecase.File) {
	newAuthRoutes(handler, googleAPI, driveAPI, us)
	newFilesRoutes(handler, sr, fl, googleAPI, driveAPI)
}
