# design-pattern
Implementation of design pattern go language

## content
+ [strategy pattern](#strategy-pattern)
+ [observer pattern](#observer-pattern)

## strategy pattern
+ **define**: The algorithmic family is defined and
	encapsulated separately so that they can be replaced
	each other. This pattern makes algorithmic changes
	independent of the objects using the algorithmic.
+ **type**: behavioral pattern
+ **logical class diagram**:
	```
	FlyBehavior -> interface
		|-ways:
			|-fly()
		|-implement:
			|-FlyWithWings
			|-FlyNoWay
			|-FlyRocketPowered
	
	QuackBehavior -> interface
		|-ways:
			|-quack()
		|-implement:
			|-Quack
			|-Squeak
			|-MuteQuack
	
	Duck -> struct
		|-members:
			|-flyBehavior FlyBehavior
			|-quackBehavior QuackBehavior
		|-ways
			|-setFlyBehavior()
			|-setQuackBehavior()
			|-display()
			|-performFly()
			|-performQuack()
			|-swim()
		|-subclasses
			|-MallardDuck
			|-RedHeadDuck
			|-RubberDuck
			|-DecoyDuck
			|-ModelDuck
			
	MallardDuck -> struct
    	|-overwrite
    		|-display()
    		
	RedHeadDuck -> struct
    	|-overwrite
    		|-display()
    		
	RubberDuck -> struct
    	|-overwrite
    		|-display()
    		
	DecoyDuck -> struct
    	|-overwrite
    		|-display()
    		
	ModelDuck -> struct
    	|-overwrite
    		|-display()
	```
+ **the structure of strategy pattern**:
	+ **encapsulation class**: Secondary encapsulation of the
	the strategy is to avoid direct invocation of the
	strategy by high-level modules.
	+ **abstract strategy**: Usually, it't an interface.
	When there are repetitive logic in each implementation
	class, abstract classes are used to encapsulate this
	part of common code.
	+ **concrete strategy**: Concrete strategy roles are
	usually assumed by a group of classes encapsulating
	algorithms, which can be freely replaced as needed.
+ **strategy pattern code implementation**:
	+ **[01-strategy](./01-strategy)**
	+ **test result**
		```
		$ go test -v
		=== RUN   TestStrategy
		I'm a real Mallard duck!
		I'am flying     !
		Quack!
		I'm a real Red Headed duck!
		I'am flying     !
		Quack!
		I'm a real Rubber duck!
		I can't fly!
		Squeak!
		I'm a real Decoy duck!
		I can't fly!
		Silence!
		I'm a model duck!
		I can't fly!
		Model duck add rocket power!
		I'm flying with a rocket!
		Quack!
		--- PASS: TestStrategy (0.00s)
		PASS
		ok      design-pattern/01-strategy      0.004s

		```
+ **main advantages of strategy pattern**:
	+ Strategy classes are free switch between them because
	the strategy classes are implemented from the same abstraction.
	+ Easy to extend, adding a new strategy is very easy for
	strategy pattern, and can be extended basically without
	changing the original code.
	+ Avoid using multiple conditions. If you don't use the
	strategy pattern, you must use conditional statements
	to join all algorithms and use conditional judgement to
	decide which algorithm to use. Using multiple conditional
	judgement is very difficult to maintain.
+ **main disadvantages of strategy pattern**:
	+ Maintaining the various strategy classes imposes
	additional overhead on development.
	+ All the strategy classes must be exposed to the client
	(the caller)because it is up to client decide which
	strategies to use, so the client should know what strategies
	are and the differences between them, otherwise the
	consequences can be severs.
+ **applicable scenario**:
	+ The main logic of several class is the same, with only
	minor differences in the algorithm and behavior of partial
	logic.
	+ There are several similar behaviors, or algorithms
	that clients need to dynamically decide which to use,
	and these algorithms can bo encapsulated for clients
	to call using strategy patter.

## observer pattern
+ **define**: A one-to-many dependency between objects is
	defined so that when an object changes state, all its
	dependencies are notified and updated automatically.
