package main

import (
	"dog/command"
	"dog/util"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

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
	if !util.IsExist("conf/.env") {
		err := viper.SafeWriteConfigAs("conf/.env")
		if err != nil {
			log.Fatal(err)
		}
	}
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	command.Run(os.Args)
}
