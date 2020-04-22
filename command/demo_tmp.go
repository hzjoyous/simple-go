package command

import (
	"fmt"
	"time"
)

type demoTmp struct {
	ConsoleInterface
}

func init() {
	var command ConsoleInterface
	command = new(demoTmp)
	commandList[command.GetSignature()] = command
}

func (demoTmp demoTmp) GetSignature() string {
	return "demoTmp"
}

func (demoTmp demoTmp) GetDescription() string {
	return "this is a Description"
}

func (demoTmp demoTmp) Handle() {

	var (
		strValue   string
		intValue   int
		boolValue  bool
		floatValue float64
	)
	fmt.Println(strValue, intValue, boolValue, floatValue)

	funcObj := func(n int) {
		for {
			fmt.Print("func", n)
			time.Sleep(1*time.Second)
		}
	}
	go funcObj(1)
	go funcObj(2)
	go funcObj(3)
	time.Sleep(3 * time.Second)
	funcObj2 := demoTmpGetFunc()
	funcObj2(1)
	//fmt.Println()
}

func demoTmpGetFunc() func(n int){
	return func(n int){
		fmt.Println(n)
	}
}