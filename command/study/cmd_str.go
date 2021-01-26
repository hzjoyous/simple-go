package study

import (
	"fmt"
	"simple-go/command/console"
)

func init() {
	c := console.Console{Signature: "cmdStr", Description: "this is a template", Handle: cmdStr}
	commandList[c.Signature] = c
}


func cmdStr() {
	fmt.Println("this is a cmdStr")
}
