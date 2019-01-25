package observer

import (
	"fmt"
	"testing"
)

func TestObserver(t *testing.T){
	weatherData := newWeatherData()
	var currentConditionDisplay = newCurrentConditionsDisplay(weatherData)
	var statisticDisplay = newStatisticsDisplay(weatherData)
	var forecastDisplay = newForecastDisplay(weatherData)
	weatherData.setMeasurements(80,65,30.4)
	weatherData.setMeasurements(82,70,29.2)
	weatherData.setMeasurements(78,90,29.2)
	fmt.Println("Observers:", currentConditionDisplay, statisticDisplay, forecastDisplay)
}