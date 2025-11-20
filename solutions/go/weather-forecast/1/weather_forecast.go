// Package weather provides tools to forecast weather condition in current locations.
package weather

var (
    // CurrentCondition represent certain weather condition at the moment.
	CurrentCondition string
    // CurrentLocation show the location for CurrentCondition.
	CurrentLocation  string
)

// Forecast return a string in format "Location - current weather condition: Condition".
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
