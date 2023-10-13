package main

import (
	"fmt"
	"github.com/be-edu/weather-3/v2/weather"
)

// Imports omitted
func main() {
	fmt.Printf("Weather package version: %s\n", weather.GetVersion())
	makePrediction(51.509865, -0.118092, "London")
}
func makePrediction(lat float64, long float64, locationName string) error {
	pred, err := weather.PredictAtCoords(lat, long)
	if err != nil {
		return err
	}
	fmt.Printf("The weather prediction for %s is: %v\n", locationName, pred.ToString())
	return nil
}
