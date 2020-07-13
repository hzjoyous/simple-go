package command

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

type demoIris struct {
	ConsoleInterface
}

func init() {
	var command ConsoleInterface
	command = new(demoIris)
	commandList[command.GetSignature()] = command
}

func (demoIris demoIris) GetSignature() string {
	return "demoIris"
}

func (demoIris demoIris) GetDescription() string {
	return "this is a Description"
}

func (demoIris demoIris) Handle() {
	fmt.Println("this is a demoIris")
	app := iris.New()
	//app.Logger().SetLevel("debug")
	app.Logger().SetLevel("prod")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	// Method:   GET
	// Resource: http://localhost:8080
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
	})

	// same as app.Handle("GET", "/ping", [...])
	// Method:   GET
	// Resource: http://localhost:8080/ping
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	// Method:   GET
	// Resource: http://localhost:8080/hello
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
