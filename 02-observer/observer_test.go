package observer

import (
	"fmt"
	"reflect"
	"testing"
)

func getVariableType(v Observer) reflect.Type  {
	return reflect.TypeOf(v)
}

func TestObserver(t *testing.T){
	weatherData := newWeatherData()
	var currentConditionDisplay = newCurrentConditionsDisplay(weatherData)
	var statisticDisplay = newStatisticsDisplay(weatherData)
	var forecastDisplay = newForecastDisplay(weatherData)
	var heatIndexDisplay = newHeatIndexDisplay(weatherData)
	weatherData.setMeasurements(80,65,30.4)
	weatherData.setMeasurements(82,70,29.2)
	weatherData.setMeasurements(78,90,29.2)
	fmt.Println("Observers:", 
		getVariableType(currentConditionDisplay), getVariableType(statisticDisplay), 
		getVariableType(forecastDisplay), getVariableType(heatIndexDisplay))
}