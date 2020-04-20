package command

import (
	"fmt"
	"strconv"
	"time"
)

type demoGo struct {
	ConsoleInterface
}

func init() {
	var command ConsoleInterface
	command = new(demoGo)
	commandList[command.GetSignature()] = command
}

func (demoGo demoGo) GetSignature() string {
	return "demoGo"
}

func (demoGo demoGo) GetDescription() string {
	return "this is a Description"
}

func (demoGo demoGo) Handle() {
	go func() {
		fmt.Println("A:我是一个来自闭包的携程，end")
	}()
	numData1 := make(chan int)
	numData2 := make(chan int)

	go goB()
	go goC(numData1)
	go goD(numData2)
	go goE()

	for i := 1; i <= 10; i++ {
		resultNumData1 := <-numData1
		numData2 <-resultNumData1
		fmt.Println("Main:我是主程,我接受到了来自C程序的信息"+strconv.Itoa(resultNumData1))
	}
	go goF()
}

func goB() {
	fmt.Println("B:我是一个来自函数的携程,end")
}

func goC(numData1 chan int) {
	n := 0
	for {
		n += 1
		numData1 <- n
		fmt.Println("C:say2:我是一个来自函数的携程,我向通道中发出了一条数据："+strconv.Itoa(n))
		time.Sleep(1 * time.Second)
	}
}

func goD(numData2 chan int){
	for{
		resultNumData2:=<-numData2
		fmt.Println("D:say2:我是一个来自函数的携程，接收到通道内数据："+strconv.Itoa(resultNumData2))
	}
}

func goE(){
	for{
		fmt.Println("E:我是一个来自函数的携程，我只会这一句")
		time.Sleep(1*time.Millisecond*500)
	}
}

func goF(){
	fmt.Println("F:我是一个来自函数的携程,我预感我将不会执行完毕end")
}