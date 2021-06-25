package study

import (
	"bytes"
	"fmt"
	"dog/command/console"
	"dog/command/util"
)

func init() {
	c := console.Console{Signature: "template", Description: "this is a template", Handle: simple}
	commandList[c.Signature] = c
}

func simple() {
	n := 1000000000
	startTime := util.GetMicroTime()
	splicingStrV1(n)
	//splicingStrV2(n )
	//time.Sleep(time.Duration(4)*time.Second)
	endTime := util.GetMicroTime()
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
