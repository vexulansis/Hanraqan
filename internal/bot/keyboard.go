package bot

func createInlineKeyboard(rows ...[]InlineKeyboardButton) *InlineKeyboardMarkup {
	keyboard := new(InlineKeyboardMarkup)
	keyboard.Inline_keyboard = append(keyboard.Inline_keyboard, rows...)
	return keyboard
}
func createInlineKeyboardRow(buttons ...InlineKeyboardButton) []InlineKeyboardButton {
	var row []InlineKeyboardButton
	row = append(row, buttons...)
	return row
}
func createInlineKeyboardButton(text string, callback_data string) *InlineKeyboardButton {
	button := new(InlineKeyboardButton)
	button.Text = text
	button.Callback_data = callback_data
	return button
}
func createReplyKeyboard(rows ...[]KeyboardButton) *ReplyKeyboardMarkup {
	keyboard := new(ReplyKeyboardMarkup)
	keyboard.Reply_keyboard = append(keyboard.Reply_keyboard, rows...)
	keyboard.One_time_keyboard = true
	keyboard.Resize_keyboard = true
	return keyboard
}
func createReplyKeyboardRow(buttons ...KeyboardButton) []KeyboardButton {
	var row []KeyboardButton
	row = append(row, buttons...)
	return row
}
func createReplyKeyboardButton(text string, request_location bool) *KeyboardButton {
	button := new(KeyboardButton)
	button.Text = text
	button.Request_location = request_location
	return button
}
