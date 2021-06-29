package study

import "dog/command/console"

var commandList = make(map[string]console.Console)

func GetAllConsoles() map[string]console.Console {
	return commandList
}

func addCommand(name string, desc string, handle func()) {
	c := console.Console{Signature: name, Description: desc, Handle: handle}
	commandList[c.Signature] = c
}
