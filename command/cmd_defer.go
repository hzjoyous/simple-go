package command

import "fmt"

func init() {
	var command ConsoleInterface
	command = new(cmdDefer)
	commandList[command.GetSignature()] = command
}

type cmdDefer struct {
	ConsoleInterface
}

func (cmdDefer cmdDefer) GetSignature()string{
	return "cmdDefer"
}

func (cmdDefer cmdDefer) GetDescription()string{
	return "this is description"
}

func (cmdDefer cmdDefer) Handle(){
	defer func() {fmt.Println("a")}()
	defer func() {fmt.Println("b")}()
	defer func() {fmt.Println("c")}()
}

