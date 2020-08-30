package command

import "fmt"

type demoScanf struct {
	ConsoleInterface
}

func init() {
	var command ConsoleInterface
	command = new(demoScanf)
	commandList[command.GetSignature()] = command
}

func (demoScanf demoScanf) GetSignature() string {
	return "demoScanf"
}

func (demoScanf demoScanf) GetDescription() string {
	return "this is a Description"
}

func (demoScanf demoScanf) Handle() {

	fmt.Println("this is a demoScanf")
	var input string
	fmt.Scanf("%s", &input)
	fmt.Printf("Your input string is %s", input)
}
