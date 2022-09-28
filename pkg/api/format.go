package api

import (
	"fmt"
	"net/http"
)

// Форматирование объектов
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
		"\nTemp: %.f°C \nWind: %.f kph",
		current.Temp_c,
		current.Wind_kph)
}
