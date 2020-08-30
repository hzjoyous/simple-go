package command

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type demoChan struct {
	ConsoleInterface
}

func init() {
	var command ConsoleInterface
	command = new(demoChan)
	commandList[command.GetSignature()] = command

}

func (demoChan demoChan) GetSignature() string {
	thisCommand := flag.NewFlagSet("demoChan", flag.ExitOnError)
	fooEnable := thisCommand.Bool("enable", false, "enable")
	fooName := thisCommand.String("name", "", "name")

	if len(os.Args) > 2 {

		thisCommand.Parse(os.Args[2:])
		fmt.Println("subcommand 'demoCommand'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", thisCommand.Args())
	}

	return "demoChan"
}

func (demoChan demoChan) GetDescription() string {
	return "this is a Description"
}

func (demoChan demoChan) Handle() {

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
