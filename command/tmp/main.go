package tmp

import (
	"dog/command/console"
	"fmt"
	"github.com/robfig/cron/v3"
	"math/rand"
	"time"
)

var commandList = make(map[string]console.Console)

func GetAllConsoles() map[string]console.Console {
	return commandList
}

func init() {
	c := console.Console{Signature: "tmp", Description: "this is a template", Handle: mainAction}
	commandList[c.Signature] = c
}

func mainAction() {
	fmt.Println("this is template main")
	format := "2006-01-02 15:04:05"
	a, _ := time.Parse(format, "2021-01-25 06:00:00")
	now := time.Now()
	fmt.Println(a.Unix())
	fmt.Println(now.Unix())
	//t.Sleep(t.Second * 3)
	fmt.Println(now.Unix())
	fmt.Println(now.After(a))
	fmt.Println(a.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	//t.Sleep(t.Second * 3)
	fmt.Println(now.Hour())
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println("时间未到")
	c := cron.New()
	_, _ = c.AddFunc("30 * * * *", func() { fmt.Println("Every hour on the half hour") })
	c.Start()

	rand.Seed(time.Now().Unix())
	num := 0
	counter := 0
	t := 0
	for {
		counter = 0
		t += 1
		num = 0
		for {
			if rand.Intn(2) == 0 {
				num += 1
				//fmt.Println("+1", num)
			} else {
				num -= 1
				//fmt.Println("-1", num)
			}
			counter += 1
			if num == 8 {
				fmt.Println("第", t, "次数:", counter)
				break
			}
		}
	}

}
