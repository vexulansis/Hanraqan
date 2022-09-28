package api

type User struct {
	Id       int      `json:"id"`
	Location Location `json:"location"`
}
type Response struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}
type Location struct {
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

