package study

import (
	"fmt"
	"simple-go/command/util"
	"simple-go/command/console"
)

func init() {
	c := console.Console{Signature: "cmdFunc", Description: "this is a template", Handle: cmdFunc}
	commandList[c.Signature] = c
}

func cmdFunc() {

	fmt.Println("多值返回")
	a, b := func1()
	fmt.Println(a, b)

	startTime := util.GetMicroTime()
	cmdFuncCounter = 1
	fmt.Println(fibonacci(46))
	endTime := util.GetMicroTime()
	fmt.Println(float64(endTime-startTime) / 1000.0)
	fmt.Println("计数器统计方法执行次数", cmdFuncCounter)
	// 求第46位斐波那契数列将调用方法5942430146次
}

var (
	cmdFuncCounter int
)

func fibonacci(n int) (res int) {
	cmdFuncCounter++
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

func func1() (a string, b int) {
	a = "str"
	b = 1
	return
}
