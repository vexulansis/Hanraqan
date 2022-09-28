package bot

import (
	db "Hunraqan/internal/db"
	h "Hunraqan/internal/helpers"
)

// Запуск бота
func Start() {
	db := db.Connect()
	offset := 0
	for {
		updates := getUpdates(offset)
		for _, update := range updates {
			offset = update.Update_id + 1
			h.InfoLog(UpdateFormat(update))
			respond(update, db)
		}
	}
}
