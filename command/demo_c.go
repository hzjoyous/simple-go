package command

import (
	"fmt"
	"time"
)

type demoC struct {
	ConsoleInterface
}

func init() {
	var command ConsoleInterface
	command = new(demoC)
	commandList[command.GetSignature()] = command
}

func (demoC demoC) GetSignature() string {
	return "demoC"
}

func (demoC demoC) GetDescription() string {
	return "this is a Description"
}

func (demoC demoC) Handle() {

	fmt.Println("this is a demoC")

	plate := make(chan int)
	apple := make(chan int)
	orange := make(chan int)
	go func() {
		for {
			time.Sleep(time.Second)
			dadPlate := <-plate
			fmt.Println("盘子可以放入一个水果 dad Plate", dadPlate)
			apple <- 1
			fmt.Println("dad放入一个苹果")
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			momPlate := <-plate
			fmt.Println("盘子可以放入一个水果mom Plate", momPlate)
			orange <- 1
			fmt.Println("mom放入一个橘子")
		}

	}()

	go func() {
		for {
			time.Sleep(time.Second)
			sonOrange := <-orange
			fmt.Println("儿子取走一个橘子 sonOrange Plate", sonOrange)
			plate <- 1
			fmt.Println("儿子清空盘子")
		}

	}()
	go func() {
		for {
			time.Sleep(time.Second)
			daughterApple := <-apple
			fmt.Println("女儿取走一个苹果 daughterApple Plate", daughterApple)
			plate <- 1
			fmt.Println("女儿清空盘子")
		}
	}()
	plate <- 1
	for {
		time.Sleep(time.Second)
	}
}
