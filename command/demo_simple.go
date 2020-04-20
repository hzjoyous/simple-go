package command

import (
	"bytes"
	"fmt"
)

func init() {
	var command ConsoleInterface
	command = new(simple)
	commandList[command.GetSignature()] = command
}

type simple struct {
	ConsoleInterface
}

func (console simple) GetSignature() string {
	return "demoSimple"
}

func (console simple) GetDescription() string {
	return "this is description"
}
func (console simple) Handle() {
	n := 700000000
	startTime := GetMicroTime()
	//splicingStrV1(n)
	splicingStrV2(n / 10000)
	//time.Sleep(time.Duration(4)*time.Second)
	endTime := GetMicroTime()
	fmt.Println((endTime - startTime) / 1000)
}

func splicingStrV1(n int) string {
	var str bytes.Buffer
	str.WriteString("msg")
	for i := 1; i <= n; i++ {
		str.WriteString("msg")
	}
	str.Reset()
	str.WriteString("msg")
	return str.String()
}

func splicingStrV2(n int) string {
	str := "msg"
	for i := 1; i <= n; i++ {
		str += "msg"
	}
	return str
}
