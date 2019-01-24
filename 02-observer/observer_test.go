package observer

import "testing"

func TestObserver(t *testing.T){
	weatherData := newWeatherData()
	currentConditionDisplay := newCurrentConditionsDisplay(weatherData)
	statisticDisplay := newStatisticsDisplay(weatherData)
	forecastDisplay := newForecastDisplay(weatherData)
	weatherData.setMeasurements(80,65,30.4)
	currentConditionDisplay.display()
	statisticDisplay.display()
	forecastDisplay.display()
	weatherData.setMeasurements(82,70,29.2)
	weatherData.setMeasurements(78,90,29.2)
	currentConditionDisplay.display()
	statisticDisplay.display()
	forecastDisplay.display()
}