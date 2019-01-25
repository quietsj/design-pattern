package observer

import "fmt"

type HeatIndexDisplay struct {
	heatIndex float64
	weatherData *WeatherData
}

func newHeatIndexDisplay(weatherData *WeatherData) *HeatIndexDisplay {
	t := new(HeatIndexDisplay)
	t.weatherData = weatherData
	weatherData.registerObserver(t)
	return t
}

func (t *HeatIndexDisplay)update(temperature float64, humidity float64, pressure float64)  {
	t.heatIndex = t.computeHeatIndex(temperature, humidity)
	t.display()
}

func (t *HeatIndexDisplay)computeHeatIndex(temperature float64, humidity float64) float64 {
	index := (float64)((16.923 + (0.185212 * temperature) + (5.37941 * humidity) -
		(0.100254 * temperature * humidity)+ (0.00941695 * (temperature * temperature)) +
		(0.00728898 * (humidity * humidity)) +
		(0.000345372 * (temperature * temperature * humidity)) -
		(0.000814971 * (temperature * humidity * humidity)) +
		(0.0000102102 * (temperature * temperature * humidity * humidity)) -
		(0.000038646 * (temperature * temperature * temperature)) +
		(0.0000291583 * (humidity * humidity * humidity)) +
		(0.00000142721 * (temperature * temperature * temperature * humidity)) +
		(0.000000197483 * (temperature * humidity * humidity * humidity)) -
		(0.0000000218429 * (temperature * temperature * temperature * humidity * humidity)) +
		0.000000000843296 * (temperature * temperature * humidity * humidity * humidity)) -
		(0.0000000000481975 * (temperature * temperature * temperature * humidity * humidity * humidity)))
	return index
}

func (t *HeatIndexDisplay)display()  {
	fmt.Println("Heat index is", t.heatIndex)
}