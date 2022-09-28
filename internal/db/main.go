package db

import (
	h "Hunraqan/internal/helpers"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// подключение к БД

func Connect() *sql.DB {
	host := h.GetEnvStr("HOST")
	dbPort := h.GetEnvStr("DB_PORT")
	user := h.GetEnvStr("USER")
	password := h.GetEnvStr("PASSWORD")
	dbName := h.GetEnvStr("DB_NAME")
	dbURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, dbPort, user, password, dbName)
	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		h.ErrorLog("Connect/sql.Open")
	}
	h.InfoLog("Connection succeed")
	return db
}
