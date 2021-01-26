package study

import (
	"fmt"
	"simple-go/command/console"
	"time"
)

func init() {
	c := console.Console{Signature: "cmdDefer", Description: "this is a template", Handle: cmdDefer}
	commandList[c.Signature] = c
}

func cmdDefer(){
	defer func() {fmt.Println("a")}()
	defer func() {fmt.Println("b")}()
	defer func() {fmt.Println("c")}()
	cmdDeferFunc1()

	for {
		time.Sleep(time.Second)
	}
}


func cmdDeferFunc1(){
	defer func(){
		fmt.Println("this is cmdDeferFunc1 defer")
	}()
	fmt.Println("this is cmdDeferFunc1")
	cmdDeferFunc2()
}

func cmdDeferFunc2(){
	defer func(){
		fmt.Println("this is cmdDeferFunc2 defer")
	}()
	fmt.Println("this is cmdDeferFunc2")
}