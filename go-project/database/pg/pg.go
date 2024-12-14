package database

import (
	"backend/config"
	"database/sql"
	"fmt"
	"log"
	"log/slog"

	_ "github.com/lib/pq"
)

var PgService *sql.DB

func MakeConnection(cfg *config.Config) {
	// connStr := fmt.Sprintf("user=%s  password=%s  dbname=%s sslmode=%s", )
	connStr := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", cfg.Database.User_name, cfg.Database.Password, cfg.Database.User_name)

	//Make connection to database
	fmt.Println(connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		slog.Error(err.Error())
		log.Fatal("Database not connected")
	}
	//Close connection once the function is executed
	PgService = db
	SyncTable(db)
}
