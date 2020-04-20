package command

import "fmt"

type demoClass struct {
	ConsoleInterface
}

func init() {
	var command ConsoleInterface
	command = new(demoClass)
	commandList[command.GetSignature()] = command
}

func (demoClass demoClass) GetSignature() string {
	return "demoClass"
}

func (demoClass demoClass) GetDescription() string {
	return "this is a Description"
}

func (demoClass demoClass) Handle() {

	console := baseConsole{Signature: "signature", Description: "description"}
	consoleList[console.GetSignature()] = console
	consoleDemo := demoConsole{baseConsole{Signature: "asd", Description: "sada"}}
	consoleList[consoleDemo.GetSignature()] = consoleDemo
	fmt.Println(consoleDemo.GetSignature())
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

type demoConsole struct {
	baseConsole
}
