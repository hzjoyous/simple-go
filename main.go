package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"dog/command"
)

func main() {
	defer func(){
		if err:=recover() ;err!=nil{
			fmt.Println("recover ",err)
		}
	}()
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("name: ", os.Getenv("name"))
	fmt.Println("age: ", os.Getenv("age"))

	command.Run(os.Args)
}
