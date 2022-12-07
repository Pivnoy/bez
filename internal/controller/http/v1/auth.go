package v1

import (
	"bez/internal/usecase"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type authRoutes struct {
	googleAPI usecase.GoogleAPI
}

func newAuthRoutes(router *chi.Mux, googleAPI usecase.GoogleAPI) {

	a := authRoutes{googleAPI: googleAPI}

	router.Get("/auth", a.getAuthLink)
	router.Post("/auth", a.addClientToken)
}

func (a *authRoutes) getAuthLink(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	str := a.googleAPI.CreateRegLink()
	fmt.Println(str)
	w.Write([]byte(str))
}

// about info
// displayName, picture, email
func (a *authRoutes) addClientToken(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Pidarasy blyat 4")
}
