package study

import (
	"fmt"
	"os"
	"runtime"
	"dog/command/console"
)

func init() {
	c := console.Console{Signature: "cmdOs", Description: "this is a template", Handle: cmdOs}
	commandList[c.Signature] = c
}

func cmdOs(){
	hostname ,_ := os.Hostname()
	fmt.Println(hostname)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}

