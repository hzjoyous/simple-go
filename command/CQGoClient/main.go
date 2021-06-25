package CqGoClient

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
	c := console.Console{Signature: "CQGoClient", Description: "CQGoClient", Handle: mainAction}
	commandList[c.Signature] = c
}

func mainAction() {
	fmt.Println("this is template main")
	rs := newRemoteService()
	resp, err := rs.getGroupList()
	if err != nil {
		return
	}
	var groupList GroupList
	err = json.Unmarshal([]byte(resp.Body()), &groupList)
	fmt.Println(resp)
	if err != nil {
		return
	}
	//for _,value:=range groupList.Data{
	//
	//	//value.GroupID
	//
	//}
	resp,err = rs.sendGroupMsg(694675063, time.Now().String()+"sendGoodNigh finish")

	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)

	jobRun(rs)
	//fmt.Println(resp.Request)

	for {
		time.Sleep(time.Minute)
	}
}
