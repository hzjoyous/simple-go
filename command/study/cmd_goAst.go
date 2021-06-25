package study

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"dog/command/console"
)


func init() {
	c := console.Console{Signature: "cmdGoAst", Description: "this is a template", Handle: cmdGoAst}
	commandList[c.Signature] = c
}

func cmdGoAst() {
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
