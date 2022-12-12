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

type flResponse struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	FileID      string `json:"fileId"`
	DownloadURL string `json:"downloadUrl"`
	PreviewURL  string `json:"previewUrl"`
}

type fileListResponse struct {
	Files []flResponse `json:"files"`
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
	var fls []flResponse
	for _, st := range fileList {
		fls = append(fls, flResponse{
			Name:        st.FileName,
			Type:        st.FileType,
			FileID:      st.FileID,
			DownloadURL: st.DownloadURL,
			PreviewURL:  st.PreviewURL})
	}
	c.JSON(http.StatusOK, fileListResponse{Files: fls})
}

func (f *filesRoutes) copyFileToService(c *gin.Context) {
	c.JSON(http.StatusOK, "ok!")
}
