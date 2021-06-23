package MiraiHttpClient

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type messageListenServer struct {
}

func (receiver messageListenServer) run() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/listenEvent", listenEvent)
	err := r.Run(":8000")
	if err != nil {
		fmt.Println(err)
	}
}

func listenEvent(c *gin.Context) {
	fmt.Println("this is listenEvent")
}
