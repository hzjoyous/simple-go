package bootstrap

import (
	"dog/util"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
)

func GetApp() application {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover ", err)
		}
	}()
	path := "conf"
	logPath := "./dog.log"
	about := "A faith tool"

	return application{
		about: about,
		log:   logInitialization(logPath),
		config: configManager{
			configInitialization(path),
		},
	}
}

type application struct {
	about  string
	log    *logrus.Logger
	config configManager
}
type configManager struct {
	Viper *viper.Viper
}

func (receiver *application) GetLogger() *logrus.Logger {
	return receiver.log
}
func (receiver configManager) get() {

}

func (receiver *configManager) Add(name string, configuration map[string]interface{}) {
	receiver.Viper.Set(name, configuration)
}

// Get 获取配置项，允许使用点式获取，如：app.name
func (receiver *configManager) Get(path string, defaultValue ...interface{}) interface{} {
	// 不存在的情况
	if !receiver.Viper.IsSet(path) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return receiver.Viper.Get(path)
}

func configInitialization(confDir string) *viper.Viper {
	configType := "env"
	configFile := ".env"
	configPath := confDir + "/" + configFile
	ViperEntity := viper.New()
	ViperEntity.SetConfigType(configType)
	// 判断目录是否存在，及创建
	if _, err := os.Stat(confDir); err != nil {
		// 判断是否存在
		if os.IsNotExist(err) {
			// file does not exist
			if err := os.MkdirAll(confDir, 0755); err != nil {
				log.Fatal(err)
			}
		}
	}
	if !util.IsExist(configPath) {
		err := ViperEntity.SafeWriteConfigAs(configPath)
		if err != nil {
			log.Fatal(err)
		}
	}

	ViperEntity.AddConfigPath(confDir)
	ViperEntity.SetConfigName(configFile)

	err := ViperEntity.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(ViperEntity.GetString("env"))
	return ViperEntity
}

func logInitialization(logPath string) *logrus.Logger {
	var logger = logrus.New()
	logger.Out = os.Stdout
	// You could set this to any `io.Writer` such as a file
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.Out = file
	} else {
		logger.Info("Failed to log to file, using default stderr")
	}
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logger.WithFields(logrus.Fields{
		"loggerInit": "success",
	}).Info("Logger success")
	return logger
}
