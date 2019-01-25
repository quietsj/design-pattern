package observer

type WeatherData struct {
	observers []Observer
	temperature float64
	humidity float64
	pressure float64
}

func newWeatherData() *WeatherData {
	t := new(WeatherData)
	t.observers = make([]Observer, 0)
	return t
}

func (t *WeatherData)registerObserver(o Observer)  {
	t.observers = append(t.observers, o)
}

func (t *WeatherData)removeObserver(o Observer)  {
	for i, v := range t.observers{
		if v == o{
			t.observers = append(t.observers[:i], t.observers[i+1:]...)
			break
		}
	}
}

func (t *WeatherData)notifyObservers()  {
	for _, v := range t.observers{
		v.update(t.temperature, t.humidity, t.pressure)
	}
}

func (t *WeatherData)measurementsChanged()  {
	t.notifyObservers()
}

func (t *WeatherData)setMeasurements(temperature float64, humidity float64, pressure float64)  {
	t.temperature = temperature
	t.humidity = humidity
	t.pressure = pressure
	t.measurementsChanged()
}

func (t *WeatherData)getTemperature() float64 {
	return t.temperature
}

func (t *WeatherData)getHumidity() float64 {
	return t.humidity
}

func (t *WeatherData)getPressure() float64 {
	return t.pressure
}