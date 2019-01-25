package strategy

import "fmt"

type RubberDuck struct {
	Duck
}

func newRubberDuck() RubberDuck {
	t := new(RubberDuck)
	t.Duck = newDuck(new(FlyNoWay), new(Squeak))
	return *t
}

func (t *RubberDuck)display()  {
	fmt.Println("I'm a real Rubber duck!")
}