package study

import "dog/command/console"

var commandList = make(map[string]console.Console)

func GetAllConsoles() map[string]console.Console {
	return commandList
}
