package config

import models "github.com/OlegLaban/weather_service/internal/modules/weather/Domain/Models"

type Config interface {
	GetLocation() models.Location
}
