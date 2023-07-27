// Package weather has everything related to the weather.
package weather

// CurrentCondition represents the condition for a given place.
var CurrentCondition string

// CurrentLocation is the current city.
var CurrentLocation string

// Forecast returns the weather condition for a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
