package main

import (
	"backend/config"
	database "backend/database/pg"
	v1 "backend/internal/v1"
	"log"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

// Main file for start the project
func main() {
	//Set config file
	cfg := config.MustLoad()
	//database connection
	database.MakeConnection(cfg)
	//register router
	router := mux.NewRouter()
	v1.New(router)
	//register main router files
	hostAddress := cfg.Host + cfg.Port
	server := http.Server{
		Addr:    hostAddress,
		Handler: router,
	}
	slog.Info("Server is running on ", slog.String("server", hostAddress))

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Server is not start", err)
	}

}
