package observer

import "fmt"

type StatisticsDisplay struct {
	maxTemperature float64
	minTemperature float64
	temperatureSum float64
	numReadings int
	weatherData *WeatherData
}

func newStatisticsDisplay(weatherData *WeatherData) *StatisticsDisplay {
	t := new(StatisticsDisplay)
	t.minTemperature = 200
	t.weatherData = weatherData
	weatherData.registerObserver(t)
	return t
}

func (t *StatisticsDisplay)update(temperature float64, humidity float64, pressure float64)  {
	t.temperatureSum += temperature
	t.numReadings++
	if temperature > t.maxTemperature{
		t.maxTemperature = temperature
	}
	if temperature < t.minTemperature{
		t.minTemperature = temperature
	}
	t.display()
}

func (t *StatisticsDisplay)display()  {
	fmt.Printf("Avg/Max/Min temperature = %.2f/%.2f/%.2f\n", t.temperatureSum / float64(t.numReadings),
		t.maxTemperature, t.minTemperature)
}

