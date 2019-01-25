package strategy

import "fmt"

type Quack struct {}

func (t *Quack)quack()  {
	fmt.Println("Quack!")
}