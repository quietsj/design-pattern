package strategy

import (
	"fmt"
)



type FlyBehavior interface {
	fly()
}
type FlyWithWings struct {}
func (t *FlyWithWings)fly()  {
	fmt.Println("I'am flying	!")
}
type FlyNoWay struct {}
func (t *FlyNoWay)fly()  {
	fmt.Println("I can't fly!")
}
type FlyRocketPowered struct {}
func (t *FlyRocketPowered)fly()  {
	fmt.Println("I'm flying with a rocket!")
}



type QuackBehavior interface {
	quack()
}
type Quack struct {}
func (t *Quack)quack()  {
	fmt.Println("Quack!")
}
type Squeak struct {}
func (t *Squeak)quack()  {
	fmt.Println("Squeak!")
}
type MuteQuack struct {}
func (t *MuteQuack)quack()  {
	fmt.Println("Silence!")
}



type Duck struct {
	flyBehavior FlyBehavior
	quackBehavior QuackBehavior
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



type MallardDuck struct {
	Duck
}
func (t *MallardDuck)display()  {
	fmt.Println("I'm a real Mallard duck!")
}



type RedHeadDuck struct {
	Duck
}
func (t *RedHeadDuck)display()  {
	fmt.Println("I'm a real Red Headed duck!")
}



type RubberDuck struct {
	Duck
}
func (t *RubberDuck)display()  {
	fmt.Println("I'm a real Rubber duck!")
}



type DecoyDuck struct {
	Duck
}
func (t *DecoyDuck)display()  {
	fmt.Println("I'm a real Decoy duck!")
}



type ModelDuck struct {
	Duck
}
func (t *ModelDuck)display()  {
	fmt.Println("I'm a model duck!")
}