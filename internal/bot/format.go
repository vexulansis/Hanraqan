package bot

import (
	"fmt"
	"net/http"
)

// Форматирование объектов
func UpdateFormat(update Update) string {
	return fmt.Sprintf(
		"\n\tUpdate_id: %d \nMessage: %sCallbackQuery: %s \n",
		update.Update_id,
		MessageFormat(update.Message),
		Callback_queryFormat(update.Callback_query))
}
func MessageFormat(message Message) string {
	return fmt.Sprintf(
		"\n\tMessage_id: %d \nFrom: %sChat: %s\tText: %s \nLocation: %s \n",
		message.Message_id,
		UserFormat(message.From),
		ChatFormat(message.Chat),
		message.Text,
		LocationFormat(message.Location))
}
func Callback_queryFormat(callback_query CallbackQuery) string {
	return fmt.Sprintf(
		"\n\tId: %s \nFrom: %sMessage: %s \nData: %s \n",
		callback_query.Id,
		UserFormat(callback_query.From),
		MessageFormat(callback_query.Message),
		callback_query.Data)
}
func UserFormat(user User) string {
	return fmt.Sprintf(
		"\n\tId: %d \n\tIs_bot: %t \n\tFirst_name: %s \n\tLast_name: %s \n\tUsername: %s \n",
		user.Id,
		user.Is_bot,
		user.First_name,
		user.Last_name,
		user.Username)
}
func ChatFormat(chat Chat) string {
	return fmt.Sprintf(
		"\n\tId: %d \n\tType: %s \n",
		chat.Id,
		chat.Type)
}

func TgLocationFormat(location Location) string {
	return fmt.Sprintf(
		"\n\tLatitude: %.6f \n\tLongitude: %.6f",
		location.Latitude,
		location.Longitude)
}
func HttpResponseFormat(response http.Response) string {
	return fmt.Sprintf(
		"\nStatus: %s\nStatusCode: %d \nProto: %s\nProtoMajor: %d \nProtoMinor: %d",
		response.Status,
		response.StatusCode,
		response.Proto,
		response.ProtoMajor,
		response.ProtoMinor)
}
func ResponseFormat(response Response) string {
	return fmt.Sprintf(
		"\nLocation: %sCurrent: %s \n",
		LocationFormat(response.Location),
		CurrentFormat(response.Current))
}
func LocationFormat(location Location) string {
	return fmt.Sprintf(
		"\nLatitude: %.2f \nLongitude: %.2f \n",
		location.Latitude,
		location.Longitude)
}
func CurrentFormat(current Current) string {
	return fmt.Sprintf(
		"\nTemp: %.2f \nWind: %.2f",
		current.Temp_c,
		current.Wind_kph)
}
