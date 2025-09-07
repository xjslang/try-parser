package tryparser

import (
	"fmt"
	"testing"

	"github.com/xjslang/xjs/lexer"
	"github.com/xjslang/xjs/parser"
)

func TestParser(t *testing.T) {
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
	p.UseStatementParser(ParseTryStatement)
	ast := p.ParseProgram()
	fmt.Println(ast.String())
}

func BenchmarkParse(b *testing.B) {
	//
}
