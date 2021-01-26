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
	_, err = c.AddFunc("@every 20m", func() { fmt.Println("Every hour thirty, starting an hour thirty from now") })
	if err != nil {
		fmt.Println(err)
	}
	c.Run()
}

func upFunc(f func(client maraiHttpClient), client maraiHttpClient) func() {
	return func() {
		_ = client.verifySession()
		f(client)
	}
}
func heartSend(client maraiHttpClient) {
	_, _ = client.sendGroupMessage("814935975", time.Now().String())
}
func sendGoodMorning(client maraiHttpClient) {
	for _, qq := range getQQPeopleList() {
		result, _ := client.sendFriendMessage(qq.QQNumber, "zaozaozao~")
		fmt.Println(result)
	}
	_, _ = client.sendGroupMessage("814935975", time.Now().String()+"sendGoodMorning finish")
}
func sendGoodNight(client maraiHttpClient) {
	for _, qq := range getQQPeopleList() {
		result, _ := client.sendFriendMessage(qq.QQNumber, "biubiubiu~")
		fmt.Println(result)
	}
	_, _ = client.sendGroupMessage("814935975", time.Now().String()+"sendGoodNigh finish")
}
