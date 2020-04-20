package command

import (
	"fmt"
	"sort"
	"time"
)

var commandList = make(map[string]ConsoleInterface)

func GetMicroTime() int64 {
	return time.Now().UnixNano() / 1000000
}

func Run(args []string) {
	fmt.Println("run")
	var action string
	if len(args) > 1 {
		action = args[1]
	} else {
		action = "help"
	}
	fmt.Println(action)
	if action == "help" {
		var keys []string
		for k := range commandList {
			keys = append(keys, k)
		}
		//按字典升序排列
		sort.Strings(keys)

		for _, k := range keys {
			fmt.Println("\t" + commandList[k].GetSignature())
		}

		if false {
			for _, consoleValue := range commandList {
				// 断言1
				console, ok := consoleValue.(ConsoleInterface)
				if ok {
					fmt.Println("\t" + console.GetSignature())
				} else {
					fmt.Println("none")
				}

				// 断言2
				//switch console := consoleValue.(type) {
				//case ConsoleInterface:
				//	fmt.Println(console.GetDescription())
				//default:
				//	fmt.Println("none")
				//}
			}
		}

	} else {

		console, ok := commandList[action]
		if ok {
			console.Handle()
		} else {
			fmt.Println("command not found")
		}
	}

}
