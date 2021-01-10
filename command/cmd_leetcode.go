package command

type cmdLeetCode struct {
	ConsoleInterface
}

func init() {
	var command ConsoleInterface
	command = new(cmdLeetCode)
	commandList[command.GetSignature()] = command
}

func (cmdLeetCode cmdLeetCode) GetSignature() string {
	return "cmdLeetCode"
}

func (cmdLeetCode cmdLeetCode) GetDescription() string {
	return "this is a Description"
}

func (cmdLeetCode cmdLeetCode) Handle() {

}
