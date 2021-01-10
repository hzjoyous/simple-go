package command

import "fmt"

type cmdStudy struct {
	ConsoleInterface
}

func init() {
	var command ConsoleInterface
	command = new(cmdStudy)
	commandList[command.GetSignature()] = command
}

func (cmdStudy cmdStudy) GetSignature() string {
	return "cmdStudy"
}

func (cmdStudy cmdStudy) GetDescription() string {
	return "this is a Description"
}

func (cmdStudy cmdStudy) Handle() {
	fmt.Println("this is a cmdStudy")
}

func uiaIf() {

}
