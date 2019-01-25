package strategy

import "fmt"

type Squeak struct {}

func (t *Squeak)quack()  {
	fmt.Println("Squeak!")
}