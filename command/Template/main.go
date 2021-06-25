package Template

import (
	"fmt"
	"dog/command/console"
)

var commandList = make(map[string]console.Console)

func GetAllConsoles() map[string]console.Console {
	return commandList
}

func init() {
	c := console.Console{Signature: "template", Description: "this is a template", Handle: mainAction}
	commandList[c.Signature] = c
}

func mainAction() {
	fmt.Println("this is template main")
}
