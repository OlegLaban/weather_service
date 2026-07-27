package output

type ShortWeatherInfoDTO struct {
	Temp          float64 `json:"temp"`
	WindSpeed     string  `json:"wind_speed"`
	WindDirection string  `json:"wind_direction"`
}
