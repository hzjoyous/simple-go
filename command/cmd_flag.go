package command

import (
	"flag"
	"fmt"
)

type cmdFlag struct {
	ConsoleInterface
}

func init() {
	var command ConsoleInterface
	command = new(cmdFlag)
	commandList[command.GetSignature()] = command
}

func (cmdFlag cmdFlag) GetSignature() string {
	return "cmdFlag"
}

func (cmdFlag cmdFlag) GetDescription() string {
	return "this is a Description"
}

func (cmdFlag cmdFlag) Handle() {

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
