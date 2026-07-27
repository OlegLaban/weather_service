package models

type WeatherInfoDTO struct {
	Temp     Temp
	Provider Provider
	Location Location
	Wind     Wind
}
