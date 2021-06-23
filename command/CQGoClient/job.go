package CqGoClient

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func jobRun(client remoteService) {
	c := cron.New()
	var err error
	//_, err := c.AddFunc("* * * * *", upFunc(heartSend, client))
	_, err = c.AddFunc("*/10 * * * *", upFunc(sendGoodNight, client))
	if err != nil {
		fmt.Println(err)
	}
	c.Run()
}

func upFunc(f func(client remoteService), client remoteService) func() {
	return func() {
		//_ = client.verifySession()
		f(client)
	}
}


func sendGoodNight(client remoteService) {

	_,_ = client.sendGroupMsg(694675063, time.Now().String()+"sendGoodNigh finish")

}
