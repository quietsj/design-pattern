package strategy

import "fmt"

type DecoyDuck struct {
	Duck
}

func newDecoyDuck() DecoyDuck {
	t := new(DecoyDuck)
	t.Duck = newDuck(new(FlyNoWay), new(MuteQuack))
	return *t
}

func (t *DecoyDuck)display()  {
	fmt.Println("I'm a real Decoy duck!")
}