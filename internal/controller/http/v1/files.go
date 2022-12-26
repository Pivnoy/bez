package v1

import (
	"bez/internal/entity"
	"bez/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"net/http"
	"regexp"
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

type fileCopyRequest struct {
	URL string `json:"url"`
}

func (f *filesRoutes) copyFileToService(c *gin.Context) {
	var fc fileCopyRequest
	if err := c.ShouldBindJSON(&fc); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	match := regexp.MustCompile("^https:\\/\\/.*/d/(.*)/.*$")
	res := match.FindStringSubmatch(fc.URL)
	if len(res) != 2 {
		errorResponse(c, http.StatusBadRequest, "cannot parse url")
		return
	}
	srv, err := f.srv.GetServices(c.Request.Context())
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	cl, err := f.ggl.CreateClient(c.Request.Context(), &oauth2.Token{RefreshToken: srv[0].RefreshToken, TokenType: "Bearer"})
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	dr, err := f.drive.UserDrive(c.Request.Context(), cl)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	fl, err := f.drive.CopyFile(dr, res[1])
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	personal, err := f.drive.GetPersonalInfo(dr)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = f.fl.StoreFile(c.Request.Context(), entity.FileTorrent{
		ID:          uuid.New(),
		FileName:    fl.Name,
		FileType:    fl.MimeType,
		FileID:      fl.Id,
		Count:       0,
		OwnerEmail:  personal.Email,
		DownloadURL: fl.WebContentLink,
		PreviewURL:  fl.WebViewLink,
	})
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

type touchFileRequest struct {
	FileID string `json:"fileId"`
}
