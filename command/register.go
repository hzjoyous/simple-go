package command

import (
	"simple-go/command/MaraiHttpClient"
	"simple-go/command/Template"
	"simple-go/command/console"
	"simple-go/command/cq"
	"simple-go/command/other"
	"simple-go/command/study"
	"simple-go/command/tmp"
)

func init() {
	p(Template.GetAllConsoles())
	p(study.GetAllConsoles())
	p(cq.GetAllConsoles())
	p(other.GetAllConsoles())
	p(MaraiHttpClient.GetAllConsoles())
	p(tmp.GetAllConsoles())

}

func p(newMap map[string]console.Console){
	pushCommandList(newMap)
}

func pushCommandList(newMap map[string]console.Console){
	commandList = consoleMapMerge(commandList,newMap)
}

func consoleMapMerge(mapA map[string]console.Console,mapB map[string]console.Console) map[string]console.Console{
	for consoleName,consoleEntity := range mapB {
		mapA[consoleName] = consoleEntity
	}
	return mapA
}
