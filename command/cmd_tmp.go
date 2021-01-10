package command

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

type cmdTmp struct {
	ConsoleInterface
}

func init() {
	var command ConsoleInterface
	command = new(cmdTmp)
	commandList[command.GetSignature()] = command
}

func (cmdTmp cmdTmp) GetSignature() string {
		return "cmdTmp"
}

func (cmdTmp cmdTmp) GetDescription() string {
	return "this is a Description"
}


func (cmdTmp cmdTmp) Handle() {
	fmt.Println(time.Now().UTC().Format(http.TimeFormat))
	fmt.Println(time.Now().Hour()<5)
	fmt.Println(strings.Contains("嗯,接着说", "接着说"))
}