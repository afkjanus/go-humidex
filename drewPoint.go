package humidex

import "math"

const (
	a = 6.1121 // mbar
	b = 18.678
	c = 257.14 // °C
)

// DrewPoint returns the drew point for a given temperature in °C and relative humidity in %.
//
// The calculations are based on: https://en.wikipedia.org/wiki/Dew_point
func DrewPoint(airTemperature float64, relativeHumidity float64) float64 {
	pa := pA(airTemperature, relativeHumidity)
	return (c * math.Log(pa/a)) / (b - math.Log(pa/a))
}

func pA(airTemperature float64, relativeHumidity float64) float64 {
	return a * math.Exp(gamma(airTemperature, relativeHumidity))
}

func gamma(airTemperature float64, relativeHumidity float64) float64 {
	return math.Log(relativeHumidity/100) + ((b * airTemperature) / (c + airTemperature))
}
