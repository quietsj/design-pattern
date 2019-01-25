package strategy

import (
	"fmt"
	"testing"
)

func TestStrategy(t *testing.T)  {
	var mallard = newMallardDuck()
	var redHead = newRedHeadDuck()
	var rubber  = newRubberDuck()
	var decoy   = newDecoyDuck()
	var model   = newModelDuck()
	mallard.display()
	mallard.performFly()
	mallard.performQuack()
	redHead.display()
	redHead.performFly()
	redHead.performQuack()
	rubber.display()
	rubber.performFly()
	rubber.performQuack()
	decoy.display()
	decoy.performFly()
	decoy.performQuack()
	model.display()
	model.performFly()
	fmt.Println("Model duck add rocket power!")
	model.setFlyBehavior(new(FlyRocketPowered))
	model.performFly()
	model.performQuack()
}