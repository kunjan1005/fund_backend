package user

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func New(r *mux.Router) {
	//create singup router in user
	r.HandleFunc("/user/singnup", func(w http.ResponseWriter, req *http.Request) {
		var body TypeSingup
		json.NewDecoder(req.Body).Decode(&body)
		Singup(body)
	}).Methods("POST")
}
