package study

import (
	"dog/command/console"
)


func init() {
	c := console.Console{Signature: "cmdLeetCode", Description: "this is a template", Handle: cmdLeetCode}
	commandList[c.Signature] = c
}


func cmdLeetCode() {

}
