// Package weather provides tools to forecast weather conditions for a given city.
package weather

var (
	// CurrentCondition describes the current weather condition of a city.
	CurrentCondition string
	// CurrentLocation represents the city the forecast is about.
	CurrentLocation string
)

// Forecast returns a string indicating the current weather condition for a specified city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
