package study

import (
	"fmt"
	"dog/command/console"
)


func init() {
	c := console.Console{Signature: "cmdStudy", Description: "this is a template", Handle: cmdStudy}
	commandList[c.Signature] = c
}


type ceshi struct {
	ceshiFunc func()
}

func cmdStudy() {
	fmt.Println("this is a cmdStudy")
	c := ceshi{ceshiFunc: ceshif}
	c.ceshiFunc()
}

func ceshif(){
	fmt.Println("niaho")
}
func uiaIf() {

}
