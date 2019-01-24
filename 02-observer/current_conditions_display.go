package observer

import "fmt"

type CurrentConditionsDisplay struct {
	temperature float64
	humidity float64
	weatherData *WeatherData
}

func newCurrentConditionsDisplay(weatherData *WeatherData) *CurrentConditionsDisplay {
	t := new(CurrentConditionsDisplay)
	t.weatherData = weatherData
	weatherData.registerObserver(t)
	return t
}

func (t *CurrentConditionsDisplay)update(temperature float64, humidity float64, pressure float64)  {
	t.temperature = temperature
	t.humidity = humidity
	
}

func (t *CurrentConditionsDisplay)display()  {
	fmt.Printf("Current conditions: %.2f F degrees and %.2f%s\n", t.temperature, t.humidity, "% humidity")
}