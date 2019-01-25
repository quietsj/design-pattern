package strategy

import "fmt"

type ModelDuck struct {
	Duck
}

func newModelDuck() ModelDuck {
	t := new(ModelDuck)
	t.Duck = newDuck(new(FlyNoWay), new(Quack))
	return *t
}

func (t *ModelDuck)display()  {
	fmt.Println("I'm a model duck!")
}