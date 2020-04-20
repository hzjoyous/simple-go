package command

import "fmt"

func init() {
	var command ConsoleInterface
	command = new(demoDefer)
	commandList[command.GetSignature()] = command
}

type demoDefer struct {
	ConsoleInterface
}

func (demoDefer demoDefer) GetSignature()string{
	return "demoDefer"
}

func (demoDefer demoDefer) GetDescription()string{
	return "this is description"
}

func (demoDefer demoDefer) Handle(){
	defer func() {fmt.Println("a")}()
	defer func() {fmt.Println("b")}()
	defer func() {fmt.Println("c")}()
}

