package study

import (
	"fmt"
	"simple-go/command/console"
	"time"
)


func init() {
	c := console.Console{Signature: "cmdChan", Description: "cmdChan", Handle: cmdChan}
	commandList[c.Signature] = c
}

func cmdChan() {

	chanVar := make(chan int)
	go func() {
		for {
			varDad := <-chanVar
			fmt.Println("dad", varDad)
		}
	}()
	go func() {
		for {
			varMom := <-chanVar
			fmt.Println("mom", varMom)
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			time.Sleep(time.Second)
			chanVar <- i
		}
	}()

	for {
		time.Sleep(time.Second)
	}
}
