package main

import (
	"dog/command"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	defer func(){
		if err:=recover() ;err!=nil{
			fmt.Println("recover ",err)
		}
	}()
	viper.SetConfigType("env")

	err := viper.SafeWriteConfigAs("conf/.env")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("name: ", os.Getenv("name"))
	fmt.Println("age: ", os.Getenv("age"))

	command.Run(os.Args)
}
