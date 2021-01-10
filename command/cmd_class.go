package command

import "fmt"

type cmdClass struct {
	ConsoleInterface
}

func init() {
	var command ConsoleInterface
	command = new(cmdClass)
	commandList[command.GetSignature()] = command
}

func (cmdClass cmdClass) GetSignature() string {
	return "cmdClass"
}

func (cmdClass cmdClass) GetDescription() string {
	return "this is a Description"
}

func (cmdClass cmdClass) Handle() {

	console := baseConsole{Signature: "signature", Description: "description"}
	consoleList[console.GetSignature()] = console
	consolecmd := cmdConsole{baseConsole{Signature: "asd", Description: "sada"}}
	consoleList[consolecmd.GetSignature()] = consolecmd
	fmt.Println(consolecmd.GetSignature())
}

var consoleList = make(map[string]baseConsoleInterface)

type baseConsoleInterface interface {
	GetSignature() string
	GetDescription() string
	Handler()
}

type baseConsole struct {
	baseConsoleInterface
	Signature   string
	Description string
}

func (obj baseConsole) GetSignature() string {
	return obj.Signature
}
func (obj baseConsole) GetDescription() string {
	return obj.Description
}
func (obj baseConsole) Handler() {

}

type cmdConsole struct {
	baseConsole
}
