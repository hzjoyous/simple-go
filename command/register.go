package command

import (
	CqGoClient "simple-go/command/CQGoClient"
	"simple-go/command/MiraiHttpClient"
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
	p(MiraiHttpClient.GetAllConsoles())
	p(tmp.GetAllConsoles())
	p(CqGoClient.GetAllConsoles())

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
