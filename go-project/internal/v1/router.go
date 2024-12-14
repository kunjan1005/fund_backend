package v1

import (
	"backend/internal/v1/middleware"
	"backend/internal/v1/user"

	"github.com/gorilla/mux"
)

func New(r *mux.Router) {
	//register controller of router
	v1Router := r.PathPrefix("/v1").Subrouter()
	user.New(v1Router)

	v1Router.Use(middleware.New)
}
