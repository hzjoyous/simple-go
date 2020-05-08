package command

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type demoSpider struct {
	ConsoleInterface
}

func init() {
	var command ConsoleInterface
	command = new(demoSpider)
	commandList[command.GetSignature()] = command
}

func (demoSpider demoSpider) GetSignature() string {
	return "demoSpider"
}

func (demoSpider demoSpider) GetDescription() string {
	return "this is a Description"
}

func (demoSpider demoSpider) Handle() {
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
