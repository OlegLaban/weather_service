package handlers

import (
	"errors"

	domain_errors "github.com/OlegLaban/weather_service/internal/modules/weather/Domain/Errors"
	models "github.com/OlegLaban/weather_service/internal/modules/weather/Domain/Models"
)

func (h *handlers) GetWeather(getDto models.GetWeatherDto) (models.WeatherInfoDTO, error) {
	loc := h.c.GetLocation()
	actualTemp := h.ws.Get(getDto.Time, loc.Lat, loc.Long)

	convertedTemp, err := h.tc.Convert(actualTemp.Temp, getDto.TempType)

	if err != nil {
		return actualTemp, errors.Join(domain_errors.ErrCantGetWeatherFromStorage, err)
	}

	actualTemp.Temp = convertedTemp

	return actualTemp, nil
}
