package command

import (
	"math/rand"
	"time"
)

var cqRandMessageList map[string]string
var cqRandMessageKey []string
var maxCQRandMessage int

func getRandMessage() string {
	return cqRandMessageList[cqRandMessageKey[rand.Intn(maxCQRandMessage)]]
}

func init() {
	rand.Seed(time.Now().UnixNano())
	cqRandMessageList = map[string]string{
		"1":  ".....",
		"2":  "卧槽无情",
		"3":  getRandFace(),
		"4":  getRandFace(),
		"5":  getRandFaceRepeat(3),
		"6":  "呵呵",
		"7":  "hehe",
		"8":  "好无聊",
		"9":  getRandFaceRepeat(6),
		"10": "好无聊" + getRandFace(),
		"11": "......",
		"12": "......",
		"13": "......",
		"14": "......",
		"15": "......",
		"16": "......",
		"17": "......",
		"18": "......",
		"19": "......",
		"20": "......",
	}
	for k := range cqRandMessageList {
		cqRandMessageKey = append(cqRandMessageKey, k)
	}
	maxCQRandMessage = len(cqRandMessageList) - 1
}
