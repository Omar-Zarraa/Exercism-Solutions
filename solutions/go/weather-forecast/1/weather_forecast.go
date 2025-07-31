// Package weather provides the forecast for the country of Goblinocus
// it contains 2 variables CurrentCondition and CurrentLocation
// and one function Forecast.
package weather

// CurrentCondition represents the current weather condition.
var CurrentCondition string

// CurrentLocation represents the current city.
var CurrentLocation string

// Forecast takes two arguments, city and condition, and returns a string that describes the current weather condition in the city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
