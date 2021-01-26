package study

import (
	"fmt"
	"simple-go/command/console"
	"sync"
	"time"
)


func init() {
	c := console.Console{Signature: "cmdMonkey", Description: "this is a template", Handle: cmdMonkey}
	commandList[c.Signature] = c
}


func cmdMonkey() {
	fmt.Println("this is a cmdMonkey")
	monkeyInit()
	go pM1Run()
	go pM2Run()

	for {
		time.Sleep(time.Second)
	}
}

func pM1Run() {
	i := 0
	for {
		i += 1
		go m1Run(i)
		time.Sleep(time.Second)
	}
}
func pM2Run() {
	i := 0
	for {
		i += 1
		go m2Run(i)
		time.Sleep(time.Second)
	}
}

var m1List int
var m2List int
var m1L map[int]int
var m2L map[int]int

type waitMonkey struct {
	tag          int
	monkeyNumber int
}

var maybeWaitList map[waitMonkey]int
var canAdd int
var canUse = sync.Mutex{}

func monkeyInit() {
	m1List = 0
	m2List = 0
	canAdd = 0
	m1L = make(map[int]int)
	m2L = make(map[int]int)

	maybeWaitList = make(map[waitMonkey]int)
}

func simpleP(i int) {
	for i <= 0 {
		time.Sleep(time.Second)
	}

}

func m1Run(monkeyNumber int) {

	tag := "1:北遍第"
	wm := waitMonkey{1, monkeyNumber}
	if m1List == 0 && m2List == 0 {
		fmt.Println(tag, monkeyNumber, "队伍没人，猴子上树")
		maybeWaitList[wm] = 1
		fmt.Println("当前被阻塞的monkey", maybeWaitList)
		canUse.Lock()
		delete(maybeWaitList, wm)
		m1List += 1
		m1L[monkeyNumber] = 1

	} else {
		if canAdd == 1 && m1List != 0 {
			fmt.Println(tag, monkeyNumber, "本队有人，直接上树")
			m1List += 1
			m1L[monkeyNumber] = 1
		} else {
			fmt.Println(tag, monkeyNumber, "已经发车")
			maybeWaitList[wm] = 1
			fmt.Println("当前被阻塞的monkey", maybeWaitList)
			canUse.Lock()
			delete(maybeWaitList, wm)
			fmt.Println(tag, monkeyNumber, "本队无人申请上树")
			m1List += 1
			m1L[monkeyNumber] = 1
		}
	}

	fmt.Println(tag, monkeyNumber, "准备发车")
	time.Sleep(time.Second * 3)

	fmt.Println(tag, monkeyNumber, "#################################")
	fmt.Println(tag, monkeyNumber, m1L)
	fmt.Println(tag, monkeyNumber, "开始下车")
	fmt.Println(tag, monkeyNumber, "#################################")
	canAdd = 0
	m1List -= 1
	delete(m1L, monkeyNumber)
	if m1List == 0 {
		canAdd = 1
		canUse.Unlock()
	} else {
		fmt.Println(tag, monkeyNumber, m1List, "车上尚有人")
	}

}

func m2Run(monkeyNumber int) {
	tag := "2:南遍第"
	wm := waitMonkey{2, monkeyNumber}
	if m1List == 0 && m2List == 0 {
		fmt.Println(tag, monkeyNumber, "队伍没人，猴子上树")
		maybeWaitList[wm] = 1
		fmt.Println("当前被阻塞的monkey", maybeWaitList)

		canUse.Lock()
		delete(maybeWaitList, wm)
		m2List += 1
		m2L[monkeyNumber] = 1
	} else {
		if canAdd == 1 && m2List != 0 {
			fmt.Println(tag, monkeyNumber, "本队有人，直接上树")
			m2List += 1
			m2L[monkeyNumber] = 1
		} else {
			fmt.Println(tag, monkeyNumber, "已经发车")
			maybeWaitList[wm] = 1
			fmt.Println("当前被阻塞的monkey", maybeWaitList)
			canUse.Lock()
			delete(maybeWaitList, wm)
			fmt.Println(tag, monkeyNumber, "开始上树")
			m2List += 1
			m2L[monkeyNumber] = 1
		}
	}
	fmt.Println(tag, monkeyNumber, "准备发车")
	time.Sleep(time.Second * 3)

	canAdd = 0

	fmt.Println(tag, monkeyNumber, "#################################")
	fmt.Println(tag, monkeyNumber, m2L)
	fmt.Println(tag, monkeyNumber, "开始下车")
	fmt.Println(tag, monkeyNumber, "#################################")

	m2List -= 1
	delete(m2L, monkeyNumber)
	if m2List == 0 {
		canAdd = 1
		canUse.Unlock()
	} else {
		fmt.Println(tag, monkeyNumber, m2List, "车上尚有人")
	}

}
