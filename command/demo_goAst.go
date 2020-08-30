package command

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type demoGoAst struct {
	ConsoleInterface
}

func init() {
	var command ConsoleInterface
	command = new(demoGoAst)
	commandList[command.GetSignature()] = command
}

func (demoGoAst demoGoAst) GetSignature() string {
	return "demoGoAst"
}

func (demoGoAst demoGoAst) GetDescription() string {
	return "this is a Description"
}

func (demoGoAst demoGoAst) Handle() {
	fmt.Println("this is a demoGoAst")
	src := `
package main
func main() {
    println("Hello, World!")
}
`

	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	// Print the AST.
	ast.Print(fset, f)
}
