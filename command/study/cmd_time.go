package study

import (
	"fmt"
	"dog/command/console"
	"time"
)


func init() {
	c := console.Console{Signature: "cmdStudy", Description: "this is a template", Handle: cmdStudy}
	commandList[c.Signature] = c
}



func cmdTime() {
	fmt.Println(time.Now().Unix())
}
