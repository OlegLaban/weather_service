package models

type GetWeatherDto struct {
	Time     int
	TempType TemperatureType
}

func New(timestamp int, tempType TemperatureType) GetWeatherDto {
	return GetWeatherDto{Time: timestamp, TempType: tempType}
}
