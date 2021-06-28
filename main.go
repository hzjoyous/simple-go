package main

import (
	"dog/command"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover ", err)
		}
	}()
	viper.SetConfigType("env")
	confDir := "conf"
	if _, err := os.Stat(confDir); err != nil {
		if os.IsNotExist(err) {
			// file does not exist

			if err := os.MkdirAll(confDir, 0755); err != nil {
				return
			}
		} else {
			return
			// other error
		}
	}
	err := viper.SafeWriteConfigAs("conf/.env")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("name: ", os.Getenv("name"))
	fmt.Println("age: ", os.Getenv("age"))

	command.Run(os.Args)
}
