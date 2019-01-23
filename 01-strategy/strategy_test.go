package strategy

import (
	"fmt"
	"testing"
)

func TestStrategy(t *testing.T)  {
	var mallard = MallardDuck{Duck{new(FlyWithWings), new(Quack)}}
	var redHead = RedHeadDuck{Duck{new(FlyWithWings), new(Quack)}}
	var rubber  = RubberDuck{Duck{new(FlyNoWay), new(Squeak)}}
	var decoy   = DecoyDuck{Duck{new(FlyNoWay),new(MuteQuack)}}
	var model   = ModelDuck{Duck{new(FlyNoWay), new(Quack)}}
	mallard.display()
	mallard.performFly()
	mallard.performQuack()
	fmt.Println("--------------------")
	redHead.display()
	redHead.performFly()
	redHead.performQuack()
	fmt.Println("--------------------")
	rubber.display()
	rubber.performFly()
	rubber.performQuack()
	fmt.Println("--------------------")
	decoy.display()
	decoy.performFly()
	decoy.performQuack()
	fmt.Println("--------------------")
	model.display()
	model.performFly()
	fmt.Println("Model duck add rocket power!")
	model.setFlyBehavior(new(FlyRocketPowered))
	model.performFly()
	model.performQuack()
}

/*
command: go test -v
=== RUN   TestStrategy
I'm a real Mallard duck!
I'am flying     !
Quack!
--------------------
I'm a real Red Headed duck!
I'am flying     !
Quack!
--------------------
I'm a real Rubber duck!
I can't fly!
Squeak!
--------------------
I'm a real Decoy duck!
I can't fly!
Silence!
--------------------
I'm a model duck!
I can't fly!
Model duck add rocket power!
I'm flying with a rocket!
Quack!
--- PASS: TestStrategy (0.00s)
PASS
ok      design-pattern/01-strategy      0.004s

 */