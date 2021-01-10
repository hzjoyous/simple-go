package command

import "fmt"

type cmdScanf struct {
	ConsoleInterface
}

func init() {
	var command ConsoleInterface
	command = new(cmdScanf)
	commandList[command.GetSignature()] = command
}

func (cmdScanf cmdScanf) GetSignature() string {
	return "cmdScanf"
}

func (cmdScanf cmdScanf) GetDescription() string {
	return "this is a Description"
}

func (cmdScanf cmdScanf) Handle() {

	fmt.Println("this is a cmdScanf")
	var input string
	fmt.Scanf("%s", &input)
	fmt.Printf("Your input string is %s", input)
}
