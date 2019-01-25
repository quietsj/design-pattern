package strategy

import "fmt"

type RedHeadDuck struct {
	Duck
}

func newRedHeadDuck() RedHeadDuck {
	t := new(RedHeadDuck)
	t.Duck = newDuck(new(FlyWithWings), new(Quack))
	return *t
}

func (t *RedHeadDuck)display()  {
	fmt.Println("I'm a real Red Headed duck!")
}