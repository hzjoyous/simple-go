package MiraiHttpClient

import (
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func jobRun(client MiraiHttpClient) {
	c := cron.New()
	var err error
	//_, err := c.AddFunc("* * * * *", upFunc(heartSend, client))
	_, err = c.AddFunc("*/10 * * * *", upFunc(sendHello, client))
	_, err = c.AddFunc("@every 20m", upFunc(continueSession, client))
	if err != nil {
		fmt.Println(err)
	}
	c.Run()
}

func sendHello(client MiraiHttpClient) {
	hiClient := newHiTokotoClient("")
	resp, _ := hiClient.getOneTokoto()
	var hitResp HiTokotoResponse
	_ = json.Unmarshal(resp.Body(), &hitResp)

	msg := hitResp.Hitokoto + "----" + hitResp.From
	_, _ = client.sendGroupMessage("694675063", msg)
}

func upFunc(f func(client MiraiHttpClient), client MiraiHttpClient) func() {
	return func() {
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

func continueSession(client MiraiHttpClient) {
	result, err := client.verify(client.adminQQNumber)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("session 更新成功", result.String())
	}
}
