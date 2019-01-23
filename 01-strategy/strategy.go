package strategy

import (
	"fmt"
)



//封装飞行行为，飞行行为算法族
type FlyBehavior interface {
	fly()
}
//实现鸭子飞行行为
type FlyWithWings struct {}
func (t *FlyWithWings)fly()  {
	fmt.Println("I'am flying	!")
}
//什么事都不做，不会飞
type FlyNoWay struct {}
func (t *FlyNoWay)fly()  {
	fmt.Println("I can't fly!")
}
//火箭动力飞行
type FlyRocketPowered struct {}
func (t *FlyRocketPowered)fly()  {
	fmt.Println("I'm flying with a rocket!")
}



//封装呱呱叫行为，呱呱叫行为算法族
type QuackBehavior interface {
	quack()
}
//实现鸭子呱呱叫
type Quack struct {}
func (t *Quack)quack()  {
	fmt.Println("Quack!")
}
//橡皮鸭吱吱叫
type Squeak struct {}
func (t *Squeak)quack()  {
	fmt.Println("Squeak!")
}
//什么事都不做，不会叫
type MuteQuack struct {}
func (t *MuteQuack)quack()  {
	fmt.Println("Silence!")
}



//客户：鸭子超类
type Duck struct {
	flyBehavior FlyBehavior
	quackBehavior QuackBehavior
}
//设置飞行行为
func (t *Duck)setFlyBehavior(fb FlyBehavior)  {
	t.flyBehavior = fb
}
//设置呱呱叫行为
func (t *Duck)setQuackBehavior(qb QuackBehavior)  {
	t.quackBehavior = qb
}
//展示
func (t *Duck)display()  {

}
//表演飞行
func (t *Duck)performFly()  {
	t.flyBehavior.fly()
}
//表演呱呱叫
func (t *Duck)performQuack()  {
	t.quackBehavior.quack()
}
//游泳
func (t *Duck)swim()  {
	fmt.Println("All ducks float, even decoys!")
}



//绿头鸭
type MallardDuck struct {
	Duck
}
//重写display方法
func (t *MallardDuck)display()  {
	fmt.Println("I'm a real Mallard duck!")
}



//红头鸭
type RedHeadDuck struct {
	Duck
}
//重写display方法
func (t *RedHeadDuck)display()  {
	fmt.Println("I'm a real Red Headed duck!")
}



//橡皮鸭
type RubberDuck struct {
	Duck
}
//重写display方法
func (t *RubberDuck)display()  {
	fmt.Println("I'm a real Rubber duck!")
}



//诱饵鸭
type DecoyDuck struct {
	Duck
}
//重写display方法
func (t *DecoyDuck)display()  {
	fmt.Println("I'm a real Decoy duck!")
}



//模型鸭
type ModelDuck struct {
	Duck
}
//重写display方法
func (t *ModelDuck)display()  {
	fmt.Println("I'm a model duck!")
}