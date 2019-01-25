package strategy

import "fmt"

type MuteQuack struct {}

func (t *MuteQuack)quack()  {
	fmt.Println("Silence!")
}