package observer

import "fmt"

type ForecastDisplay struct {
	currentPressure float64
	lastPressure float64
	weatherData *WeatherData
}

func newForecastDisplay(weatherData *WeatherData) *ForecastDisplay {
	t := new(ForecastDisplay)
	t.currentPressure = 29.92
	t.weatherData = weatherData
	weatherData.registerObserver(t)
	return t
}

func (t *ForecastDisplay)update(temperature float64, humidity float64, pressure float64)  {
	t.lastPressure = t.currentPressure
	t.currentPressure = pressure
	t.display()
}

func (t *ForecastDisplay)display()  {
	fmt.Printf("Forecast: ")
	if t.currentPressure > t.lastPressure{
		fmt.Println("Improving weather on the way!")
	} else if t.currentPressure == t.lastPressure{
		fmt.Println("More of the same!")
	}else {
		fmt.Println("Watch out for cooler, rainy weather!")
	}
}