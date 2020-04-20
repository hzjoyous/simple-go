package command

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type demoGin struct {
	ConsoleInterface
}

func init() {
	var command ConsoleInterface
	command = new(demoGin)
	commandList[command.GetSignature()] = command
}

func (demoGin demoGin) GetSignature() string {
	return "demoGin"
}

func (demoGin demoGin) GetDescription() string {
	return "this is a Description"
}

func (demoGin demoGin) Handle() {
	r := setupRouter()
	_ = r.Run(":8080")
	fmt.Println(" Listen and Server in 0.0.0.0:8080")
}

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	//r := gin.Default()
	r := gin.New()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}
