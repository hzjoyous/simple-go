package command

import "fmt"

type demoArrAndMap struct {
	ConsoleInterface
}

func init() {
	var command ConsoleInterface
	command = new(demoArrAndMap)
	commandList[command.GetSignature()] = command
}

func (demoArrAndMap demoArrAndMap) GetSignature() string {
	return "demoArrAndMap"
}

func (demoArrAndMap demoArrAndMap) GetDescription() string {
	return "this is a Description"
}

func (demoArrAndMap demoArrAndMap) Handle() {
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


	demoArrAndMapPrintSlice(intArr3)
	fmt.Println(len(intArr3),cap(intArr3))
	intArr3 = append(intArr3, 1,2,3,4,5,6,7)
	fmt.Println(len(intArr3),cap(intArr3))
	demoArrAndMapPrintSlice(intArr3)
	intArr3[0] = 9
	demoArrAndMapPrintSlice(intArr3)
}
func demoArrAndMapPrintSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
