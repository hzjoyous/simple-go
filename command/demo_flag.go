package command

import (
	"flag"
	"fmt"
)

type demoFlag struct {
	ConsoleInterface
}

func init() {
	var command ConsoleInterface
	command = new(demoFlag)
	commandList[command.GetSignature()] = command
}

func (demoFlag demoFlag) GetSignature() string {
	return "demoFlag"
}

func (demoFlag demoFlag) GetDescription() string {
	return "this is a Description"
}

func (demoFlag demoFlag) Handle() {

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
