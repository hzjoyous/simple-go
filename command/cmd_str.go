package command

import "fmt"

type cmdStr struct {
	ConsoleInterface
}

func init() {
	var command ConsoleInterface
	command = new(cmdStr)
	commandList[command.GetSignature()] = command
}

func (cmdStr cmdStr) GetSignature() string {
	return "cmdStr"
}

func (cmdStr cmdStr) GetDescription() string {
	return "this is a Description"
}

func (cmdStr cmdStr) Handle() {
	fmt.Println("this is a cmdStr")
}
