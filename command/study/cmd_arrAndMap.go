package study

import (
	"fmt"
	"simple-go/command/console"
)

func init() {
	c := console.Console{Signature: "cmdArrAndMap", Description: "cmdArrAndMap", Handle: cmdArrAndMap}
	commandList[c.Signature] = c
}

func cmdArrAndMap() {
	var (
		intArr  [2]int
		intArr2 [2][2]int
		intArr3 []int
		mapArr map[int]string
	)


	fmt.Println(intArr)
	fmt.Println(intArr2)
	fmt.Println(mapArr)



	if mapArr == nil {
		fmt.Println("it is nil")
	}else{
		fmt.Println("it not not nil")
	}

	fmt.Println("mapArr len is", len(mapArr))

	mapArr = make(map[int]string, 0)  // 必要，否则panic
	fmt.Println(mapArr)
	if mapArr == nil {
		fmt.Println("it is nil")
	}else{
		fmt.Println("it not not nil")
	}

	fmt.Println("map len4 is", len(mapArr))

	mapArr[1] = "string"

	fmt.Println(mapArr)

	fmt.Println("综上表现，goMap的初始化的设计就是脱裤子放屁")


	cmdArrAndMapPrintSlice(intArr3)
	fmt.Println(len(intArr3),cap(intArr3))
	intArr3 = append(intArr3, 1,2,3,4,5,6,7)
	fmt.Println(len(intArr3),cap(intArr3))
	cmdArrAndMapPrintSlice(intArr3)
	intArr3[0] = 9
	cmdArrAndMapPrintSlice(intArr3)
}
func cmdArrAndMapPrintSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
