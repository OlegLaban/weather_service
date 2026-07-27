package contracts

import output "github.com/OlegLaban/weather_service/internal/DTO/Output"

type Handlers interface {
	Ping() (output.PingDTO, error)
	GetShortWeatherInfo() (output.ShortWeatherInfoDTO, error)
}
