package command

import "fmt"

type cmd struct {
	ConsoleInterface
}

func init() {
	var command ConsoleInterface
	command = new(cmd)
	commandList[command.GetSignature()] = command
}

func (cmd cmd) GetSignature() string {
	return "cmd"
}

func (cmd cmd) GetDescription() string {
	return "this is a Description"
}

func (cmd cmd) Handle(){
	fmt.Println("this is a cmd")
}
