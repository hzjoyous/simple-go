package command


type ConsoleInterface interface {
	GetSignature()string
	GetDescription()string
	Handle()
}

func init() {
	var command ConsoleInterface
	command = new(console)
	commandList[command.GetSignature()] = command
}

type console struct {
	ConsoleInterface
}

func (console console) GetSignature()string{
	return "console"
}

func (console console) GetDescription()string{
	return "this is description"
}

func (console console) Handle(){

}

