package command

import "fmt"

func init() {
	var command ConsoleInterface
	command = new(demoFunc)
	commandList[command.GetSignature()] = command
}

type demoFunc struct {
	ConsoleInterface
}

func (console demoFunc) GetSignature() string {
	return "demoFunc"
}

func (console demoFunc) GetDescription() string {
	return "this is description"
}

func (console demoFunc) Handle() {
	a,b := func1()
	fmt.Println(a,b)

	startTime := GetMicroTime()
	fmt.Println(fibonacci(46))
	endTime := GetMicroTime()
	fmt.Println(float64(endTime - startTime) / 1000.0)
}


func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}


func func1()(a string,b int){
	a = "str"
	b = 1
	return
}