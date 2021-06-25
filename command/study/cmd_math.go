package study

import (
	"fmt"
	"math/big"
	"dog/command/console"
	"dog/util"
	"sort"
)

func init() {
	c := console.Console{Signature: "cmdMath", Description: "this is a template", Handle: cmdMath}
	commandList[c.Signature] = c
}


func cmdMath() {
	fmt.Println("this is a cmdMath")
	startTime := util.GetMicroTime()
	cmdMathCounter = 0
	//fmt.Println(cmdMathFibonacci(59424))
	//fmt.Println(cmdMathOneFibonacci(59424))
	fmt.Println(cmdMathOneFibonacci(5942))
	show := false

	if show {
		var keys []int
		for k := range fibonacciList {
			keys = append(keys, k)
		}
		//按字典升序排列
		sort.Ints(keys)
		for _, k := range keys {
			fmt.Println("\t", k, ":", fibonacciList[k])
		}
		//for  k,v := range fibonacciList{
		//	fmt.Println(k,v)
		//}
	}

	endTime := util.GetMicroTime()
	fmt.Println(float64(endTime-startTime) / 1000.0)
	fmt.Println("计数器统计方法执行次数", cmdMathCounter)
}

var fibonacciList = make(map[int]*big.Int)

var (
	cmdMathCounter int
)

func cmdMathOneFibonacci(n int) (res *big.Int) {
	a := big.NewInt(1)
	b := big.NewInt(1)
	c := big.NewInt(0)
	for i := 1; i <= n; i++ {
		c = cmdMathAdd(a, b)
		a = b
		b = c
	}
	res = c
	return
}

func cmdMathFibonacci(n int) (res *big.Int) {
	cmdMathCounter++
	res, ok := fibonacciList[n]
	if ok {
		return
	}
	if n <= 1 {
		res = big.NewInt(1)
	} else {
		f1, ok := fibonacciList[n-1]
		if !ok {
			f1 = cmdMathFibonacci(n - 1)
		}
		f2, ok := fibonacciList[n-2]
		if !ok {
			f2 = cmdMathFibonacci(n - 2)
		}
		res = cmdMathAdd(f1, f2)
	}
	fibonacciList[n] = res
	return
}

func cmdMathAdd(x, y *big.Int) *big.Int {
	return big.NewInt(0).Add(x, y)
}
