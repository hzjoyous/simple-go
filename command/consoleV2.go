package command

import "fmt"

type consoleV2 struct {
	ConsoleInterface
}

func init() {
	var command ConsoleInterface
	command = new(consoleV2)
	commandList[command.GetSignature()] = command
}

func (consoleV2 consoleV2) GetSignature() string {
	return "consoleV2"
}

func (consoleV2 consoleV2) GetDescription() string {
	return "this is a Description"
}

func (consoleV2 consoleV2) Handle(){
	fmt.Println("this is a consoleV2")
}


type consoleV2Data struct {
	name string
	desc string
}

type consoleV2TestInterface interface {

}

type consoleV2Test struct {

}
