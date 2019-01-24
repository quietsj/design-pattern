package observer

type Observer interface {
	update(temperature float64, humidity float64, pressure float64)
}