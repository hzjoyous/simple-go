package command

import (
	"fmt"
	"dog/command/console"
)

func init() {
	c := console.Console{Signature: "cmd", Description: "this is a cmd", Handle: handle}
	commandList[c.Signature] = c
}

func handle() {
	fmt.Println("this is a demo")
}
