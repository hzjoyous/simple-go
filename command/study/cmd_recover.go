package study

import (
	"fmt"
	"dog/command/console"
)

func init() {
	c := console.Console{Signature: "recoverTest", Description: "cmdC", Handle: recoverTest}
	commandList[c.Signature] = c
}

var panicNumber int

func recoverTest() {
	panicNumber = 0
	for {
		if panicNumber > 3 {
			fmt.Println("出错三次，退出")
			break
		}
		recoverAction()
		panicNumber += 1
	}
}
func recoverAction() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出了错：", err)
		}
	}()
	myPanic()
	fmt.Printf("这里应该执行不到！")
}

func myPanic() {
	var x = 30
	var y = 0
	for i := 0; i < 5; i++ {

		//if i == 3 {
		//	panic("我就是一个大错误！")
		//}

		c := x / y
		fmt.Println(c)
	}
}
