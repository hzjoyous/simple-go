package study

import (
	"fmt"
	"simple-go/command/console"
)


func init() {
	c := console.Console{Signature: "cmdClass", Description: "this is a template", Handle: cmdClass}
	commandList[c.Signature] = c
}


func cmdClass() {

	consoleEntity := baseConsole{Signature: "signature", Description: "description"}
	consoleList[consoleEntity.GetSignature()] = consoleEntity
	consoleCmd := cmdConsole{baseConsole{Signature: "asd", Description: "sada"}}
	consoleList[consoleCmd.GetSignature()] = consoleCmd
	fmt.Println(consoleCmd.GetSignature())
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
