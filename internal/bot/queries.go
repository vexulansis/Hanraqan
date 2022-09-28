package bot

import (
	h "Hunraqan/internal/helpers"
	"database/sql"
)

func existUser(update *Update, db *sql.DB) bool {
	query := `SELECT count(id) FROM users WHERE id=$1 LIMIT 1;`
	count := 0
	rows, err := db.Query(query, update.Message.From.Id)
	if err != nil {
		h.ErrorLog("existUser/db.Query")
	}
	for rows.Next() {
		rows.Scan(&count)
	}
	return count != 0
}
func insertUser(update Update, db *sql.DB) {
	if existUser(&update, db) {
		h.InfoLog("User already exists")
	} else {
		query := `INSERT INTO users VALUES ($1,$2,$3,$4);`
		_, err := db.Exec(query, update.Message.From.Id, update.Message.From.Username, update.Message.Location.Latitude, update.Message.Location.Longitude)
		if err != nil {
			h.ErrorLog("insertUser/db.Exec")
		}
	}
}
func updateLocation(update Update, db *sql.DB) {
	query :=
		`UPDATE users
		SET latitude=$1, longitude=$2 
		WHERE id=$3;`
	_, err := db.Exec(query, update.Message.Location.Latitude, update.Message.Location.Longitude, update.Message.From.Id)
	if err != nil {
		h.ErrorLog("updateLocation/db.Exec")
	}
}
func getLocation(update Update, db *sql.DB) (float64, float64) {
	var latitude, longitude float64
	query := `SELECT latitude FROM users WHERE id=$1 LIMIT 1;`
	rows, err := db.Query(query, update.Callback_query.From.Id)
	if err != nil {
		h.ErrorLog("getLocation/db.Query")
	}
	for rows.Next() {
		rows.Scan(&latitude)
	}
	query = `SELECT longitude FROM users WHERE id=$1 LIMIT 1;`
	rows, err = db.Query(query, update.Callback_query.From.Id)
	if err != nil {
		h.ErrorLog("getLocation/db.Query")
	}
	for rows.Next() {
		rows.Scan(&longitude)
	}
	return latitude, longitude
}
