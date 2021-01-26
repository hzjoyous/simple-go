package other

import (
	"fmt"
	"os"
	"simple-go/command/console"
	"strings"
)

func init() {
	c := console.Console{Signature: "makeCommand", Description: "this is a cmd", Handle: makeCommand}
	commandList[c.Signature] = c
}

func makeCommand() {
	commandName := "defaultCommandName"
	if len(os.Args) > 2 {
		commandName = os.Args[2]
	}

	commandTemplate := `
type {commandName} struct {
	console
}

func init() {
	var command console
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

	fmt.Println(strings.Replace(commandTemplate, "{commandName}", commandName, -1))
}
