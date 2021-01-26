package study

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"simple-go/command/console"
)


func init() {
	c := console.Console{Signature: "cmdSpider", Description: "this is a template", Handle: cmdSpider}
	commandList[c.Signature] = c
}

func cmdSpider() {
	resp, err := http.Get("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	handleError(err, "http.Get url")
	defer resp.Body.Close()
	// 2.读取页面内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	handleError(err, "ioutil.ReadAll")
	// 字节转字符串
	pageStr := string(pageBytes)
	fmt.Println(pageStr)
}

func handleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}
