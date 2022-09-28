package api

import (
	h "Hunraqan/internal/helpers"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Получение метеорологических данных
func GetWeather(latitude float64, longitude float64) string {
	weatherAPI := h.GetEnvStr("WEATHER_API")
	key := h.GetEnvStr("WEATHER_API_KEY")
	q := fmt.Sprintf("%.4f,%.4f", latitude, longitude)
	resp, err := http.Get(weatherAPI + "/current.json" + "?key=" + key + "&q=" + q)
	h.InfoLog(HttpResponseFormat(*resp))
	if err != nil {
		h.ErrorLog("GetWeather/http.Get")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		h.ErrorLog("GetWeather/io.ReadAll")
	}
	response := new(Response)
	err = json.Unmarshal(body, &response)
	if err != nil {
		h.ErrorLog("GetWeather/json.Unmarshall")
	}
	h.InfoLog(ResponseFormat(*response))
	return CurrentFormat(response.Current)
}
