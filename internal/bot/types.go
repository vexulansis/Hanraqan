package bot

// Объекты чата
type User struct {
	Id         int    `json:"id"`
	Is_bot     bool   `json:"is_bot"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Username   string `json:"username"`
}
type Chat struct {
	Id   int    `json:"id"`
	Type string `json:"type"`
}
type Message struct {
	Message_id   int                  `json:"message_id"`
	From         User                 `json:"from"`
	Chat         Chat                 `json:"chat"`
	Text         string               `json:"text"`
	Location     Location             `json:"location"`
	Reply_markup InlineKeyboardMarkup `json:"reply_markup"`
}
type BotMessage struct {
	Message_id   int                  `json:"message_id"`
	Chat_id      int                  `json:"chat_id"`
	Text         string               `json:"text"`
	Reply_markup InlineKeyboardMarkup `json:"reply_markup"`
}
type BotRequest struct {
	Message_id   int                 `json:"message_id"`
	Chat_id      int                 `json:"chat_id"`
	Text         string              `json:"text"`
	Reply_markup ReplyKeyboardMarkup `json:"reply_markup"`
}

// Объекты обновлений
type Update struct {
	Update_id      int           `json:"update_id"`
	Message        Message       `json:"message"`
	Callback_query CallbackQuery `json:"callback_query"`
}
type RestResponse struct {
	Result []Update `json:"result"`
}
type CallbackQuery struct {
	Id      string  `json:"id"`
	From    User    `json:"from"`
	Message Message `json:"message"`
	Data    string  `json:"data"`
}

// Объекты геопозиции
type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

// Объекты кастомных клавиатур
type InlineKeyboardMarkup struct {
	Inline_keyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}
type InlineKeyboardButton struct {
	Text          string `json:"text"`
	Callback_data string `json:"callback_data"`
}
type ReplyKeyboardMarkup struct {
	Reply_keyboard    [][]KeyboardButton `json:"keyboard"`
	Resize_keyboard   bool               `json:"resize_keyboard"`
	One_time_keyboard bool               `json:"one_time_keyboard"`
}
type KeyboardButton struct {
	Text             string `json:"text"`
	Request_location bool   `json:"request_location"`
}

// Объекты API
type Response struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}
type APILocation struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
	Name      string  `json:"name"`
	Region    string  `json:"region"`
	Country   string  `json:"country"`
}
type Current struct {
	Temp_c   float64 `json:"temp_c"`
	Temp_f   float64 `json:"temp_f"`
	Wind_kph float64 `json:"wind_kph"`
	Wind_mph float64 `json:"wind_mph"`
}
