package command

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type cmdGoAst struct {
	ConsoleInterface
}

func init() {
	var command ConsoleInterface
	command = new(cmdGoAst)
	commandList[command.GetSignature()] = command
}

func (cmdGoAst cmdGoAst) GetSignature() string {
	return "cmdGoAst"
}

func (cmdGoAst cmdGoAst) GetDescription() string {
	return "this is a Description"
}

func (cmdGoAst cmdGoAst) Handle() {
	fmt.Println("this is a cmdGoAst")
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
