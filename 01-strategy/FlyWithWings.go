package strategy

import "fmt"

type FlyWithWings struct {}

func (t *FlyWithWings)fly()  {
fmt.Println("I'am flying	!")
}