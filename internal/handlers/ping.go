package handlers

import output "github.com/OlegLaban/weather_service/internal/DTO/Output"

func (h *handlers) Ping() (output.PingDTO, error) {
	return *output.NewPing(), nil
}
