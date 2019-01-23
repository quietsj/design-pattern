# design-pattern
Implementation of design pattern go language

## content
+ [strategy pattern](#strategy-pattern)
+ [next]()


## strategy pattern
+ **定义**：定义了算法族，分别封装起来，让它们之间可以相互替换，
此模式让算法的变化独立于使用算法的对象。
+ **类型**：行为类模式
+ **逻辑类图**：
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
+ **策略模式的结构**:
	+ **封装类**：对策略进行二次封装，目的是避免高层模块对策略的直接调用。
	+ **抽象策略**：通常情况下为一个接口，当各个实现类中存在着重复的逻辑时，则使用抽象类来封装这部分公共的代码。
	+ **具体策略**：具体策略角色通常由一组封装了算法的类来担任，这些类之间可以根据需要自由替换。
+ **策略模式的主要优点**：
	+ 策略类之间可以自由切换，由于策略类实现自同一个抽象，所以他们之间可以自由切换。
	+ 易于扩展，增加一个新的策略对策略模式来说非常容易，基本上可以在不改变原有代码的基础上进行扩展。
	+ 避免使用多重条件，如果不使用策略模式，对于所有的算法，必须使用条件语句进行连接，通过条件判断来决定使用哪一种算法，使用多重条件判断是非常不容易维护的。
+ **策略模式的主要缺点**：
	+ 维护各个策略类会给开发带来额外开销。
	+ 必须对客户端（调用者）暴露所有的策略类，因为使用哪种策略是由客户端来决定的，因此，客户端应该知道有什么策略，并且了解各种策略之间的区别，否则，后果很严重。
+ **适用场景**
	+ 几个类的主要逻辑相同，只在部分逻辑的算法和行为上稍有区别的情况。
	+ 有几种相似的行为，或者说算法，客户端需要动态地决定使用哪一种，那么可以使用策略模式，将这些算法封装起来供客户端调用。