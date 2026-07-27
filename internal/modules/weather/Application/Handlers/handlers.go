package handlers

import (
	config "github.com/OlegLaban/weather_service/internal/modules/weather/Domain/Interfaces/Config"
	storages "github.com/OlegLaban/weather_service/internal/modules/weather/Domain/Interfaces/Storages"
	services "github.com/OlegLaban/weather_service/internal/modules/weather/Domain/Services"
)

type handlers struct {
	ws storages.WeatherSource
	c  config.Config
	tc services.TempConverter
}
