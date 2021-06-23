package MiraiHttpClient

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func jobRun(client MiraiHttpClient) {
	c := cron.New()
	var err error
	//_, err := c.AddFunc("* * * * *", upFunc(heartSend, client))
	//_, err = c.AddFunc("3 6 * * *", upFunc(sendGoodMorning, client))
	//_, err = c.AddFunc("3 23 * * *", upFunc(sendGoodNight, client))
	//_, err = c.AddFunc("30 6 * * *", upFunc(breakfast, client))
	//_, err = c.AddFunc("50 11 * * *", upFunc(lunch, client))
	//_, err = c.AddFunc("54 17 * * *", upFunc(dinner, client))
	_, err = c.AddFunc("0 * * * *", upFunc(sendHello, client))
	_, err = c.AddFunc("@every 20m", upFunc(continueSession, client))
	if err != nil {
		fmt.Println(err)
	}
	c.Run()
}

func sendHello(client MiraiHttpClient) {
	_, _ = client.sendGroupMessage("694675063", time.Now().String()+"了")
}

func upFunc(f func(client MiraiHttpClient), client MiraiHttpClient) func() {
	return func() {
		//_ = client.verifySession()
		f(client)
	}
}

func heartSend(client MiraiHttpClient) {
	_, _ = client.sendGroupMessage("814935975", time.Now().String())
}

func sendGoodMorning(client MiraiHttpClient) {
	for _, qq := range getQQFriendEntityList() {
		result, _ := client.sendFriendMessage(qq.getQQNumber(), getTextMessage("zaozaozao~"))
		fmt.Println(result)
	}
	_, _ = client.sendGroupMessage("814935975", time.Now().String()+"sendGoodMorning finish")
}
func sendGoodNight(client MiraiHttpClient) {
	for _, qq := range getQQFriendEntityList() {
		result, _ := client.sendFriendMessage(qq.getQQNumber(), getFaceMessage("good night ~"))
		fmt.Println(result)
	}
	_, _ = client.sendGroupMessage("814935975", time.Now().String()+"sendGoodNigh finish")
}

func continueSession(client MiraiHttpClient) {
	result, err := client.verify(client.adminQQNumber)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("session 更新成功", result.String())
	}
}

func breakfast(client MiraiHttpClient){
	for _, qq := range getQQFriendEntityList() {
		result, _ := client.sendFriendMessage(qq.getQQNumber(), getTextMessage("吃早饭"),getFaceMessage("30"))
		fmt.Println(result)
	}
	_, _ = client.sendGroupMessage("814935975", time.Now().String()+"早饭 finish")
}

func lunch(client MiraiHttpClient){
	for _, qq := range getQQFriendEntityList() {
		result, _ := client.sendFriendMessage(qq.getQQNumber(), getTextMessage("午饭午饭"),getFaceMessage("30"))
		fmt.Println(result)
	}
	_, _ = client.sendGroupMessage("814935975", time.Now().String()+"午饭 finish")
}

func dinner(client MiraiHttpClient){
	for _, qq := range getQQFriendEntityList() {
		result, _ := client.sendFriendMessage(qq.getQQNumber(), getTextMessage("dinner~~~"),getFaceMessage("30"))
		fmt.Println(result)
	}
	_, _ = client.sendGroupMessage("814935975", time.Now().String()+"晚饭 finish")
}