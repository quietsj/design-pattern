package strategy

import "fmt"

type MallardDuck struct {
	Duck
}

func newMallardDuck() MallardDuck {
	t := new(MallardDuck)
	t.Duck = newDuck(new(FlyWithWings), new(Quack))
	return *t
}

func (t *MallardDuck)display()  {
	fmt.Println("I'm a real Mallard duck!")
}