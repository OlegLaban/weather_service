package services

import (
	"errors"

	internal_errors "github.com/OlegLaban/weather_service/internal/modules/weather/Domain/Errors"
	models "github.com/OlegLaban/weather_service/internal/modules/weather/Domain/Models"
)

type TempConverter struct{}

func (tc *TempConverter) Convert(actual models.Temp, to models.TemperatureType) (models.Temp, error) {
	if actual.Type == to {
		return actual, nil
	}

	switch to {
	case models.Celcius:
		return tc.celsiusToFahrenheit(actual), nil
	case models.Fahrenheit:
		return tc.fahrenheitToCelsius(actual), nil
	default:
		return models.Temp{}, errors.Join(internal_errors.ErrCannotConvertTemp, errors.New("unknown temp type"))
	}
}

func (tc *TempConverter) celsiusToFahrenheit(actual models.Temp) models.Temp {
	convertedValue := actual.Value*1.8 + 32

	return models.Temp{
		Type:  models.Fahrenheit,
		Value: convertedValue,
	}
}

func (tc *TempConverter) fahrenheitToCelsius(actual models.Temp) models.Temp {
	convertedValue := (actual.Value - 32) / 1.8

	return models.Temp{
		Type:  models.Celcius,
		Value: convertedValue,
	}
}
