package main

import (
	"dog/bootstrap"
	"dog/command"
	"fmt"
	"os"
)

func main() {
	//log.SetFormatter(&log.JSONFormatter{})
	app := bootstrap.GetApp()
	fmt.Println(app)


	command.Run(os.Args)
}
