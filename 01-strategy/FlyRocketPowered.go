package strategy

import "fmt"

type FlyRocketPowered struct {}

func (t *FlyRocketPowered)fly()  {
	fmt.Println("I'm flying with a rocket!")
}