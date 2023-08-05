package humidex

import "math"

// Humidex returns the humidex for a given temperature in °C and relative humidity in %.
//
// The calculations are based on: https://en.wikipedia.org/wiki/Humidex
func Humidex(airTemperature float64, relativeHumidity float64) float64 {
	return airTemperature + 0.5555*(6.11*math.Exp(5417.753*((1/273.16)-(1/(273.15+DrewPoint(airTemperature, relativeHumidity)))))-10)
}
