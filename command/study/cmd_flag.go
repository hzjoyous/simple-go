package study

import (
	"flag"
	"fmt"
	"dog/command/console"
)

func init() {
	c := console.Console{Signature: "cmdFlag", Description: "this is a template", Handle: cmdFlag}
	commandList[c.Signature] = c
}

func cmdFlag() {

	var (
		h    bool
		name string
	)

	flag.BoolVar(&h, "h", false, "this help")
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.Parse()

	fmt.Println("h is", h)
	fmt.Println(name)
	if h {
		flag.Usage()
	}
}
