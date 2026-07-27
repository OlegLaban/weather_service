package storages

import models "github.com/OlegLaban/weather_service/internal/modules/weather/Domain/Models"

type WeatherSource interface {
	Get(time int, lat float64, long float64) models.WeatherInfoDTO
}
