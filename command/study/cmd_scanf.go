package study

import (
	"fmt"
	"dog/command/console"
)


func init() {
	c := console.Console{Signature: "cmdScanf", Description: "this is a template", Handle: cmdScanf}
	commandList[c.Signature] = c
}


func cmdScanf() {

	fmt.Println("this is a cmdScanf")
	var input string
	fmt.Scanf("%s", &input)
	fmt.Printf("Your input string is %s", input)
}
