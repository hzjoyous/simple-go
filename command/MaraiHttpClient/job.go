package MaraiHttpClient

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func jobRun(client maraiHttpClient) {
	c := cron.New()
	var err error
	//_, err := c.AddFunc("* * * * *", upFunc(heartSend, client))
	_, err = c.AddFunc("3 6 * * *", upFunc(sendGoodMorning, client))
	_, err = c.AddFunc("3 23 * * *", upFunc(sendGoodNight, client))
	_, err = c.AddFunc("30 6 * * *", upFunc(breakfast, client))
	_, err = c.AddFunc("50 11 * * *", upFunc(lunch, client))
	_, err = c.AddFunc("54 17 * * *", upFunc(dinner, client))
	_, err = c.AddFunc("@every 20m", upFunc(continueSession, client))
	if err != nil {
		fmt.Println(err)
	}
	c.Run()
}

func upFunc(f func(client maraiHttpClient), client maraiHttpClient) func() {
	return func() {
		//_ = client.verifySession()
		f(client)
	}
}

func heartSend(client maraiHttpClient) {
	_, _ = client.sendGroupMessage("814935975", time.Now().String())
}

func sendGoodMorning(client maraiHttpClient) {
	for _, qq := range getQQFriendEntityList() {
		result, _ := client.sendFriendMessage(qq.getQQNumber(), getTextMessage("zaozaozao~"))
		fmt.Println(result)
	}
	_, _ = client.sendGroupMessage("814935975", time.Now().String()+"sendGoodMorning finish")
}
func sendGoodNight(client maraiHttpClient) {
	for _, qq := range getQQFriendEntityList() {
		result, _ := client.sendFriendMessage(qq.getQQNumber(), getFaceMessage("good night ~"))
		fmt.Println(result)
	}
	_, _ = client.sendGroupMessage("814935975", time.Now().String()+"sendGoodNigh finish")
}

func continueSession(client maraiHttpClient) {
	result, err := client.verify(client.adminQQNumber)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("session 更新成功", result.String())
	}
}

func breakfast(client maraiHttpClient){
	for _, qq := range getQQFriendEntityList() {
		result, _ := client.sendFriendMessage(qq.getQQNumber(), getTextMessage("吃早饭"),getFaceMessage("30"))
		fmt.Println(result)
	}
	_, _ = client.sendGroupMessage("814935975", time.Now().String()+"早饭 finish")
}

func lunch(client maraiHttpClient){
	for _, qq := range getQQFriendEntityList() {
		result, _ := client.sendFriendMessage(qq.getQQNumber(), getTextMessage("午饭午饭"),getFaceMessage("30"))
		fmt.Println(result)
	}
	_, _ = client.sendGroupMessage("814935975", time.Now().String()+"午饭 finish")
}

func dinner(client maraiHttpClient){
	for _, qq := range getQQFriendEntityList() {
		result, _ := client.sendFriendMessage(qq.getQQNumber(), getTextMessage("dinner~~~"),getFaceMessage("30"))
		fmt.Println(result)
	}
	_, _ = client.sendGroupMessage("814935975", time.Now().String()+"晚饭 finish")
}