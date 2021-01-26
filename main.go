package main

import (
	"fmt"
	"log"
	"os"
	"simple-go/command"
	"github.com/joho/godotenv"
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
