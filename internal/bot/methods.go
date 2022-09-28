package bot

import (
	h "Hunraqan/internal/helpers"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

// Методы взаимодействия с Telegram Bot API
// func sendMessage() {

// }
func getUpdates(offset int) []Update {
	botURL := h.GetEnvStr("BOT_URL")
	resp, err := http.Get(botURL + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		h.ErrorLog("getUpdates/http.Get")
		return nil
	}
	defer resp.Body.Close()
	h.InfoLog("getting updates...")
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		h.ErrorLog("getUpdates/io.ReadAll")
		return nil
	}
	restResponse := new(RestResponse)
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		h.ErrorLog("getUpdates/json.Unmarshall")
		return nil
	}
	return restResponse.Result
}
func sendMessage(botMessage BotMessage) {
	botURL := h.GetEnvStr("BOT_URL")
	buf, err := json.Marshal(botMessage)
	if err != nil {
		h.ErrorLog("sendMessage/json.Marshal")
	}
	resp, err := http.Post(botURL+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	h.InfoLog(HttpResponseFormat(*resp))
	if err != nil {
		h.ErrorLog("sendMessage/http.Post")
	}
}
func sendRequest(botRequest BotRequest) {
	botURL := h.GetEnvStr("BOT_URL")
	buf, err := json.Marshal(botRequest)
	if err != nil {
		h.ErrorLog("sendRequest/json.Marshal")
	}
	resp, err := http.Post(botURL+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	h.InfoLog(HttpResponseFormat(*resp))
	if err != nil {
		h.ErrorLog("sendRequest/http.Post")
	}
}
func editMessageReplyMarkup(botMessage BotMessage) {
	botURL := h.GetEnvStr("BOT_URL")
	buf, err := json.Marshal(botMessage)
	if err != nil {
		h.ErrorLog("editMessageReplyMarkup/json.Marshal")
	}
	resp, err := http.Post(botURL+"/editMessageReplyMarkup", "application/json", bytes.NewBuffer(buf))
	h.InfoLog(HttpResponseFormat(*resp))
	if err != nil {
		h.ErrorLog("editMessageReplyMarkup/http.Post")
	}
}
func editMessageText(botMessage BotMessage) {
	botURL := h.GetEnvStr("BOT_URL")
	buf, err := json.Marshal(botMessage)
	if err != nil {
		h.ErrorLog("editMessageText/json.Marshal")
	}
	resp, err := http.Post(botURL+"/editMessageText", "application/json", bytes.NewBuffer(buf))
	h.InfoLog(HttpResponseFormat(*resp))
	if err != nil {
		h.ErrorLog("editMessageText/http.Post")
	}
}
