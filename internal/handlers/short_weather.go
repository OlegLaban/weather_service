package handlers

import output "github.com/OlegLaban/weather_service/internal/DTO/Output"

func (h *handlers) GetShortWeatherInfo() (output.ShortWeatherInfoDTO, error) {
	return output.ShortWeatherInfoDTO{}, nil
}
