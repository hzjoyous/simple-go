package command

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type cmdChan struct {
	ConsoleInterface
}

func init() {
	var command ConsoleInterface
	command = new(cmdChan)
	commandList[command.GetSignature()] = command

}

func (cmdChan cmdChan) GetSignature() string {
	thisCommand := flag.NewFlagSet("cmdChan", flag.ExitOnError)
	fooEnable := thisCommand.Bool("enable", false, "enable")
	fooName := thisCommand.String("name", "", "name")

	if len(os.Args) > 2 {

		thisCommand.Parse(os.Args[2:])
		fmt.Println("subcommand 'cmdCommand'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", thisCommand.Args())
	}

	return "cmdChan"
}

func (cmdChan cmdChan) GetDescription() string {
	return "this is a Description"
}

func (cmdChan cmdChan) Handle() {

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
