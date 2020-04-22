package command

import (
	"fmt"
	"math/big"
	"sort"
)

type demoMath struct {
	ConsoleInterface
}

func init() {
	var command ConsoleInterface
	command = new(demoMath)
	commandList[command.GetSignature()] = command
}

func (demoMath demoMath) GetSignature() string {
	return "demoMath"
}

func (demoMath demoMath) GetDescription() string {
	return "this is a Description"
}

func (demoMath demoMath) Handle() {
	fmt.Println("this is a demoMath")
	startTime := GetMicroTime()
	demoMathCounter = 0
	//fmt.Println(demoMathFibonacci(59424))
	//fmt.Println(demoMathOneFibonacci(59424))
	fmt.Println(demoMathOneFibonacci(5942))
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

	endTime := GetMicroTime()
	fmt.Println(float64(endTime-startTime) / 1000.0)
	fmt.Println("计数器统计方法执行次数",demoMathCounter)
}

var fibonacciList = make(map[int]*big.Int)

var (
	demoMathCounter int
)

func demoMathOneFibonacci(n int)(res*big.Int){
	a := big.NewInt(1)
	b := big.NewInt(1)
	c := big.NewInt(0)
	for i :=1 ;i<=n ;i++{
		c = demoMathAdd(a,b)
		a=b
		b=c
	}
	res = c
	return
}

func demoMathFibonacci(n int) (res *big.Int) {
	demoMathCounter+=1
	res, ok := fibonacciList[n]
	if ok {
		return
	}
	if n <= 1 {
		res = big.NewInt(1)
	} else {
		f1, ok := fibonacciList[n-1]
		if !ok {
			f1 = demoMathFibonacci(n - 1)
		}
		f2, ok := fibonacciList[n-2]
		if !ok {
			f2 = demoMathFibonacci(n - 2)
		}
		res = demoMathAdd(f1, f2)
	}
	fibonacciList[n] = res
	return
}

func demoMathAdd(x, y *big.Int) *big.Int {
	return big.NewInt(0).Add(x, y)
}
