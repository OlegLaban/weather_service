package models

type Temp struct {
	Type  TemperatureType
	Value float64
}

func (t Temp) Convert(tt TemperatureType) {

}
