package command

import (
	"fmt"
	"net/http"
	"dog/command/console"
	"sort"
	"strings"
	"time"
)

var commandList = make(map[string]console.Console)

// Run 运行
func Run(args []string) {
	fmt.Println("run")
	var action string
	if len(args) > 1 {
		action = args[1]
	} else {
		action = "help"
	}
	fmt.Println(action, ":")
	if action == "help" {
		var keys []string
		for k := range commandList {
			keys = append(keys, k)
		}
		//按字典升序排列
		sort.Strings(keys)

		for _, k := range keys {
			fmt.Println("\t" + commandList[k].Signature)
		}

		if false {
			for _, consoleEntity := range commandList {
				fmt.Println("\t" + consoleEntity.Signature)

			}
		}

	} else {
		consoleEntity, ok := commandList[action]
		if ok {
			consoleEntity.Handle()
		} else {
			fmt.Println("command not found")
		}
	}

}


// GetMicroTime 获取当前时间戳
func GetMicroTime() int64 {
	return time.Now().UnixNano() / 1000000
}

func inArrayInt64(need int64, haystack []int64) bool {
	for _, v := range haystack {
		if need == v {
			return true
		}
	}
	return false
}

func inArrayString(need string, haystack []string) bool {
	for _, v := range haystack {
		if strings.Contains(need, v) {
			return true
		}
	}
	return false
}

func date() string {
	return time.Now().UTC().Format(http.TimeFormat)
}

//// 断言1
//console, ok := consoleValue.(console.Console)
//if ok {
//	fmt.Println("\t" + console.Signature)
//} else {
//	fmt.Println("none")
//}

// 断言2
//switch console := consoleValue.(type) {
//case console:
//	fmt.Println(console.GetDescription())
//default:
//	fmt.Println("none")
//}