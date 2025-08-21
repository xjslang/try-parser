package main

import (
	"fmt"

	tparser "github.com/xjslang/try-parser"
	"github.com/xjslang/xjs/lexer"
	"github.com/xjslang/xjs/parser"
)

func main() {
	input := `
	try {
		console.log('try something')
	} catch {
		console.log('something unexpected happened')
	} finally {
		console.log('the end!')
	}`
	l := lexer.New(input)
	p := parser.New(l)
	p.UseParseStatement(tparser.ParseStatement)
	ast := p.ParseProgram()
	fmt.Println(ast.String())
}
