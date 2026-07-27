package models

type TemperatureType int

const (
	Celcius TemperatureType = iota
	Fahrenheit
)

func (tt TemperatureType) String() string {
	switch tt {
	case Celcius:
		return "C"
	case Fahrenheit:
		return "F"
	default:
		return "Unknown"
	}
}
