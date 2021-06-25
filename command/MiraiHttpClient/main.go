package MiraiHttpClient

import (
	"encoding/json"
	"fmt"
	"dog/command/console"
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

func managerRun() MiraiHttpClient {

	var client MiraiHttpClient

	authKey := "INITKEYYoPjKLS3"
	adminQQNumber := "536607429"

	client = newMiraiClient(authKey, adminQQNumber, "http://127.0.0.1:8080")



	result, _ := client.getAbout()

	fmt.Println(result.String())

	_ = client.verifySession()

	return client

}

func mainAction() {

	client := managerRun()

	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)

	fmt.Println(tm.Format("2006-01-02 03:04:05 PM"))

	hiClient := newHiTokotoClient("")
	resp, _ := hiClient.getOneTokoto()
	var hitResp HiTokotoResponse
	_ = json.Unmarshal(resp.Body(), &hitResp)

	msg := hitResp.Hitokoto + "----" + hitResp.From
	_, _ = client.sendGroupMessage("694675063", msg)

	go jobRun(client)

	for {
		time.Sleep(time.Minute)
	}
}

func airGroupList(client MiraiHttpClient, message string) {
	if len(message) == 0 {
		message = "from go Miraihttpclient send javaMirai-api-http"
	}
	resp, _ := client.getGroupList()
	type GroupList []struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		Permission string `json:"permission"`
	}
	var groupList GroupList
	err := json.Unmarshal([]byte(resp.Body()), &groupList)
	fmt.Println(resp)
	if err != nil {
		return
	}
	for _, groupItem := range groupList {
		message, err := client.sendGroupMessage(ToString(groupItem.ID), message)
		if err != nil {
			return
		}
		fmt.Println(message)
	}
}

func airFriendList(client MiraiHttpClient) {
	result, _ := client.friendList()
	var qqFriendList []qqFriendEntity
	_ = json.Unmarshal([]byte(result.String()), &qqFriendList)
	for _, qqEntity := range qqFriendList {
		if qqEntity.getQQNumber() == "2677659138" || qqEntity.getQQNumber() == "3425517617" {
			continue
		}
		fmt.Println(qqEntity.ID, qqEntity.Remark)
		fmt.Println(qqEntity.getQQNumber())
		result, _ := client.sendFriendMessage(qqEntity.getQQNumber(), getTextMessage("头像不错"), getFaceMessage("201"))
		fmt.Println(result)
	}
}
