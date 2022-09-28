package bot

import (
	"Hunraqan/pkg/api"
	"database/sql"
)

func respond(update Update, db *sql.DB) {
	insertUser(update, db)
	if update.Message.Text == "/start" {
		sendMessage(onStart(update, db))
	} else if update.Message.Location.Latitude != 0 && update.Message.Location.Longitude != 0 {
		onLocation(update, db)
	} else if update.Callback_query.Data == "/update" {
		onUpdate(update, db)
	} else {
		editMessageReplyMarkup(onCallback(update, db))
		editMessageText(onCallback(update, db))
	}
}
func onStart(update Update, db *sql.DB) BotMessage {
	main := createInlineKeyboard(
		createInlineKeyboardRow(*createInlineKeyboardButton("Настройки", "/options")))
	startMessage := new(BotMessage)
	startMessage.Chat_id = update.Message.Chat.Id
	startMessage.Text = api.GetWeather(getLocation(update, db))
	startMessage.Reply_markup = *main
	return *startMessage
}
func onLocation(update Update, db *sql.DB) {
	updateLocation(update, db)
}
func onUpdate(update Update, db *sql.DB) {
	request := createReplyKeyboard(createReplyKeyboardRow(*createReplyKeyboardButton("Обновить локацию", true)))
	requestMessage := new(BotRequest)
	requestMessage.Chat_id = update.Callback_query.Message.Chat.Id
	requestMessage.Text = "Обновить локацию?"
	requestMessage.Reply_markup = *request
	sendRequest(*requestMessage)
}
func onCallback(update Update, db *sql.DB) BotMessage {
	main := createInlineKeyboard(
		createInlineKeyboardRow(*createInlineKeyboardButton("Настройки", "/options")))
	options := createInlineKeyboard(
		createInlineKeyboardRow(
			*createInlineKeyboardButton("Конфигурация", "/config"),
			*createInlineKeyboardButton("Локация", "/location")),
		createInlineKeyboardRow(
			*createInlineKeyboardButton("Назад", "/main")))
	location := createInlineKeyboard(
		createInlineKeyboardRow(
			*createInlineKeyboardButton("Показать", "/show"),
			*createInlineKeyboardButton("Обновить", "/update")),
		createInlineKeyboardRow(
			*createInlineKeyboardButton("Назад", "/options")))
	show := createInlineKeyboard(createInlineKeyboardRow(*createInlineKeyboardButton("Назад", "/location")))
	currentWeather := api.GetWeather(getLocation(update, db))
	editedMessage := new(BotMessage)
	editedMessage.Chat_id = update.Callback_query.Message.Chat.Id
	editedMessage.Message_id = update.Callback_query.Message.Message_id
	editedMessage.Text = currentWeather
	switch update.Callback_query.Data {
	case "/main":
		editedMessage.Reply_markup = *main
	case "/options":
		editedMessage.Reply_markup = *options
		// case "/config":
		// 	editedMessage.Reply_markup = *config
		// 	editMessageReplyMarkup(*editedMessage)
	case "/location":
		editedMessage.Reply_markup = *location
	case "/show":
		location := new(Location)
		latitude, longitude := getLocation(update, db)
		location.Latitude = latitude
		location.Longitude = longitude
		editedMessage.Text = LocationFormat(*location)
		editedMessage.Reply_markup = *show
	}
	return *editedMessage
}
