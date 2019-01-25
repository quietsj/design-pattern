package strategy

import "fmt"

type Duck struct {
	flyBehavior FlyBehavior
	quackBehavior QuackBehavior
}

func newDuck(flyBehavior FlyBehavior, quackBehavior QuackBehavior) Duck {
	t := new(Duck)
	t.flyBehavior = flyBehavior
	t.quackBehavior = quackBehavior
	return *t
}

func (t *Duck)setFlyBehavior(fb FlyBehavior)  {
	t.flyBehavior = fb
}

func (t *Duck)setQuackBehavior(qb QuackBehavior)  {
	t.quackBehavior = qb
}

func (t *Duck)display()  {

}

func (t *Duck)performFly()  {
	t.flyBehavior.fly()
}

func (t *Duck)performQuack()  {
	t.quackBehavior.quack()
}

func (t *Duck)swim()  {
	fmt.Println("All ducks float, even decoys!")
}