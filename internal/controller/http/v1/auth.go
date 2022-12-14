package v1

import (
	"bez/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

type authRoutes struct {
	googleAPI usecase.GoogleAPI
	driveAPI  usecase.DriveAPI

	userUseCase usecase.User
}

func newAuthRoutes(handler *gin.Engine, googleAPI usecase.GoogleAPI, driveAPI usecase.DriveAPI, userUseCase usecase.User) {

	a := authRoutes{googleAPI: googleAPI, driveAPI: driveAPI, userUseCase: userUseCase}

	handler.GET("/auth", a.getAuthLink)
	handler.POST("/auth", a.addClientToken)
}

type authLinkResponse struct {
	URL string `json:"url"`
}

func (a *authRoutes) getAuthLink(c *gin.Context) {
	c.JSON(http.StatusOK, authLinkResponse{URL: a.googleAPI.CreateRegLink()})
}

type clientTokenRequest struct {
	URL string `json:"url"`
}

type clientTokenResponse struct {
	DisplayName string `json:"displayName"`
	Picture     string `json:"picture"`
	Email       string `json:"email"`
}

func (a *authRoutes) addClientToken(c *gin.Context) {
	var cl clientTokenRequest
	if err := c.ShouldBindJSON(&cl); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	match := regexp.MustCompile("code=(.*)&")
	res := match.FindStringSubmatch(cl.URL)
	if len(res) != 2 {
		errorResponse(c, http.StatusBadRequest, "cannot parse url")
		return
	}
	token, err := a.googleAPI.CreateUserToken(c.Request.Context(), res[1])
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	clientHTTP, err := a.googleAPI.CreateClient(c.Request.Context(), token)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	defer clientHTTP.CloseIdleConnections()
	dr, err := a.driveAPI.UserDrive(c.Request.Context(), clientHTTP)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	nm, err := a.driveAPI.GetPersonalInfo(dr)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = a.userUseCase.CreateUser(c.Request.Context(), nm.Email, token.RefreshToken)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, clientTokenResponse{
		DisplayName: nm.DisplayName,
		Picture:     nm.Picture,
		Email:       nm.Email,
	})
}
