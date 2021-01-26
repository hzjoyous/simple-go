package MaraiHttpClient

import (
	"fmt"
	"simple-go/command/console"
	"time"
)

var commandList = make(map[string]console.Console)

func GetAllConsoles() map[string]console.Console {
	return commandList
}

func init() {
	c := console.Console{Signature: "mclient", Description: "", Handle: mainDefendAction}
	commandList[c.Signature] = c
}

func mainDefendAction() {
	var panicNumber int
	panicNumber = 0
	for {
		if panicNumber > 3 {
			fmt.Println("出错三次，退出")
			break
		}
		mainAction()
		panicNumber += 1
	}
}

func mainAction() {

	client := managerRun()
	//result, _ := client.friendList()

	//var qqFriendList []qqFriendEntity
	//_ = json.Unmarshal([]byte(result.String()), &qqFriendList)
	//
	//for _, qqEntity := range qqFriendList {
	//	if qqEntity.getQQNumber() == "2677659138" || qqEntity.getQQNumber() == "3425517617"{
	//		continue
	//	}
	//	fmt.Println(qqEntity.ID, qqEntity.Remark)
	//	fmt.Println(qqEntity.getQQNumber())
	//	result, _ := client.sendFriendMessage(qqEntity.getQQNumber(), getTextMessage("头像不错"),getFaceMessage("201"))
	//	fmt.Println(result)
	//}
	go jobRun(client)

	//for _, qq := range qqFriendEntityList() {
	//	result, _ := client.sendFriendMessage(qq.QQNumber, "chi wan le ~")
	//	fmt.Println(result)
	//}

	for {
		time.Sleep(time.Minute)
	}
}
