package main

import (
	"testing"
	"go/token"
	"go/parser"
	"go/ast"
	"go/scanner"
	"fmt"
	"strings"
)

type Visitor int

func (v Visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}
	fmt.Printf("%s%T\n", strings.Repeat("\t", int(v)), n)
	return v + 1
}

func TestASTWalk(t *testing.T) {
	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", "package main; var a = 3", parser.ParseComments)
	if err != nil {
		panic(err)
	}
	var v Visitor
	ast.Walk(v, f)
}

func TestInspectAST(t *testing.T) {
	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", []byte(`package main
import "fmt"
import "strings"

func test1() {
    hello := "Hello"
    world := "World"
    words := []string{hello, world}
    SayHello(words)
}

// SayHello says Hello
func SayHello(words []string) bool {
    fmt.Println(joinStrings(words))
    return true
}

// joinStrings joins strings
func joinStrings(words []string) string {
    return strings.Join(words, ", ")
}`), parser.ParseComments)
	if err != nil {
		panic(err)
	}
	//fmt.Println(f.Package)
	//fmt.Println(f.Name)
	//for _,i := range f.Imports{
	//	fmt.Println(i.Path)
	//}
	for _, decl := range f.Decls {
		fmt.Println(decl)
		fn, ok := decl.(*ast.FuncDecl)
		if ok {
			fmt.Println("func: ", fn.Name)
		}

		gn, ok := decl.(*ast.GenDecl)
		if ok {
			fmt.Println("gen: ", gn.Lparen)
		}

		bd, ok := decl.(*ast.BadDecl)
		if ok {
			fmt.Println("bd: ", bd.From)
		}
	}

	ast.Inspect(f, func(n ast.Node) bool {
		// Find Return Statements
		ret, ok := n.(*ast.ReturnStmt)
		if ok {
			fmt.Printf("return statement found on line %v %v\n", fset.Position(ret.Pos()), fset.Position(ret.End()))
			return true
		}
		return true
	})

}

func TestParserAST(t *testing.T) {
	src := []byte(`/*comment0*/
package main
import "fmt"
//comment1
/*comment2*/
func main() {
  fmt.Println("Hello, world!")
}
`)

	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	// Print the AST.
	ast.Print(fset, f)
}

func TestParserAST1(t *testing.T) {
	src := []byte(`package main
import "fmt"
//comment
func main() {
 fmt.Println("Hello, world!")
}
`)

	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	s.Init(file, src, nil, scanner.ScanComments)

	for {
		pos, tok, lit := s.Scan()
		fmt.Printf("%-6s%-8s%q\n", fset.Position(pos), tok, lit)

		if tok == token.EOF {
			break
		}
	}
}
