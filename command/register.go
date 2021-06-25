package command

import (
	CqGoClient "dog/command/CQGoClient"
	"dog/command/MiraiHttpClient"
	"dog/command/Template"
	"dog/command/console"
	"dog/command/cq"
	"dog/command/other"
	"dog/command/study"
	"dog/command/tmp"
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
