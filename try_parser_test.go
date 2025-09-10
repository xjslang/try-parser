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
	ast, err := p.ParseProgram()
	if err != nil {
		t.Fatalf("ParseProgram() error: %v", err)
	}
	fmt.Println(ast.String())
}

func BenchmarkParse(b *testing.B) {
	//
}
