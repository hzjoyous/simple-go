package command

import "fmt"

type demo struct {
	ConsoleInterface
}

func init() {
	var command ConsoleInterface
	command = new(demo)
	commandList[command.GetSignature()] = command
}

func (demo demo) GetSignature() string {
	return "demo"
}

func (demo demo) GetDescription() string {
	return "this is a Description"
}

func (demo demo) Handle(){
	fmt.Println("this is a demo")
}
