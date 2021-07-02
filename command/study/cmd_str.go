package study

import (
	"bytes"
	"dog/command/console"
	"fmt"
	"strings"
	"time"
)

func init() {
	c := console.Console{Signature: "cmdStr", Description: "this is a template", Handle: cmdStr}
	commandList[c.Signature] = c
}


func cmdStr() {
	fmt.Println("this is a cmdStr")
	k := 5
	d := [5]time.Duration{}
	for i := 0; i < k; i++ {
		d[i] = benchmarkStringFunction(30000, i)
	}

	for i := 0; i < k-1; i++ {
		fmt.Printf("way %d is %6.1f times of way %d\n", i, float32(d[i])/float32(d[k-1]), k-1)
	}
}



func benchmarkStringFunction(n int, index int) (d time.Duration) {
	v := "ni shuo wo shi bu shi tai wu liao le a?"
	var s string
	var buf bytes.Buffer

	t0 := time.Now()
	for i := 0; i < n; i++ {
		switch index {
		case 0: // fmt.Sprintf
			s = fmt.Sprintf("%s[%s]", s, v)
		case 1: // string +
			s = s + "[" + v + "]"
		case 2: // strings.Join
			s = strings.Join([]string{s, "[", v, "]"}, "")
		case 3: // temporary bytes.Buffer
			// 每次声明一个bytes.buffer 仍然比正常的字符串拼接快很多
			b := bytes.Buffer{}
			b.WriteString("[")
			b.WriteString(v)
			b.WriteString("]")
			s = b.String()
		case 4: // stable bytes.Buffer
			buf.WriteString("[")
			buf.WriteString(v)
			buf.WriteString("]")
		}

	}

	if index == 4 { // for stable bytes.Buffer
		s = buf.String()
	}
	fmt.Println(len(s)) // consume s to avoid compiler optimization
	t1 := time.Now()
	d = t1.Sub(t0)
	fmt.Printf("time of way(%d)=%v\n", index, d)
	return d
}