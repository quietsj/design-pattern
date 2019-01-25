package strategy

import "fmt"

type FlyNoWay struct {}

func (t *FlyNoWay)fly()  {
	fmt.Println("I can't fly!")
}