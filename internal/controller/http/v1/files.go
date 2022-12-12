package v1

import (
	"bez/internal/entity"
	"bez/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type filesRoutes struct {
	srv   usecase.Service
	fl    usecase.File
	ggl   usecase.GoogleAPI
	drive usecase.DriveAPI
}

func newFilesRoutes(handler *gin.Engine, srv usecase.Service, f usecase.File, ggl usecase.GoogleAPI, drive usecase.DriveAPI) {

	fl := filesRoutes{srv: srv, fl: f, ggl: ggl, drive: drive}

	handler.GET("/files", fl.getFileList)
	handler.POST("/files", fl.copyFileToService)
}

type fileListResponse struct {
	Files []entity.FileTorrent `json:"files"`
}

func (f *filesRoutes) getFileList(c *gin.Context) {
	srv, err := f.srv.GetServices(c.Request.Context())
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var fileList []entity.FileTorrent
	for _, serv := range srv {
		files, err := f.fl.GetFileList(c.Request.Context(), serv.Email)
		if err != nil {
			errorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		fileList = append(fileList, files...)
	}
}

func (f *filesRoutes) copyFileToService(c *gin.Context) {
	c.JSON(http.StatusOK, "ok!")
}
