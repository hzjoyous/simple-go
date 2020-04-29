package command

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

type demoTmp struct {
	ConsoleInterface
}

func init() {
	var command ConsoleInterface
	command = new(demoTmp)
	commandList[command.GetSignature()] = command
}

func (demoTmp demoTmp) GetSignature() string {
		return "demoTmp"
}

func (demoTmp demoTmp) GetDescription() string {
	return "this is a Description"
}


func (demoTmp demoTmp) Handle() {
	fmt.Println(time.Now().UTC().Format(http.TimeFormat))
	fmt.Println(time.Now().Hour()<5)
	fmt.Println(strings.Contains("嗯,接着说", "接着说"))
}