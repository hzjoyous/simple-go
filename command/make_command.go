package command

import (
	"fmt"
	"os"
	"strings"
)

type makeCommand struct {
	ConsoleInterface
}

func init() {
	var command ConsoleInterface
	command = new(makeCommand)
	commandList[command.GetSignature()] = command
}

func (makeCommand makeCommand) GetSignature() string {
	return "makeCommand"
}

func (makeCommand makeCommand) GetDescription() string {
	return "this is a Description"
}

func (makeCommand makeCommand) Handle() {
	commandName := "defaultCommandName"
	if len(os.Args)>2 {
		commandName =os.Args[2]
	}

	commandTemplate := `
type {commandName} struct {
	ConsoleInterface
}

func init() {
	var command ConsoleInterface
	command = new({commandName})
	commandList[command.GetSignature()] = {commandName}
}

func ({commandName} {commandName}) GetSignature() string {
	return "makeCommand"
}

func ({commandName} {commandName}) GetDescription() string {
	return "this is a Description"
}

func ({commandName} {commandName}) Handle(){
	fmt.Println("this is a {commandName}")
}
`

	fmt.Println(strings.Replace(commandTemplate,"{commandName}",commandName,-1))
}
