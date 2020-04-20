package command

func init() {
	var command ConsoleInterface
	command = new(demoOs)
	commandList[command.GetSignature()] = command
}

type demoOs struct {
	ConsoleInterface
}

func (demoOs demoOs) GetSignature()string{
	return "demoOs"
}

func (demoOs demoOs) GetDescription()string{
	return "this is demoOs"
}

func (demoOs demoOs) Handle(){
}

