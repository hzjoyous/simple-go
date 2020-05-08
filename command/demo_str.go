package command

import "fmt"

type demoStr struct {
	ConsoleInterface
}

func init() {
	var command ConsoleInterface
	command = new(demoStr)
	commandList[command.GetSignature()] = command
}

func (demoStr demoStr) GetSignature() string {
	return "demoStr"
}

func (demoStr demoStr) GetDescription() string {
	return "this is a Description"
}

func (demoStr demoStr) Handle() {
	fmt.Println("this is a demoStr")
}
