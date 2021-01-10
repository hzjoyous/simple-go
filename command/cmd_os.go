package command

import (
	"fmt"
	"os"
	"runtime"
)

func init() {
	var command ConsoleInterface
	command = new(cmdOs)
	commandList[command.GetSignature()] = command
}

type cmdOs struct {
	ConsoleInterface
}

func (cmdOs cmdOs) GetSignature()string{
	return "cmdOs"
}

func (cmdOs cmdOs) GetDescription()string{
	return "this is cmdOs"
}

func (cmdOs cmdOs) Handle(){
	hostname ,_ := os.Hostname()
	fmt.Println(hostname)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}

