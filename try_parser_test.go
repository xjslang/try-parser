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
	lb := lexer.NewBuilder()
	p := parser.NewBuilder(lb).Install(Plugin).Build(input)
	ast, err := p.ParseProgram()
	if err != nil {
		t.Fatalf("ParseProgram() error: %v", err)
	}
	fmt.Println(ast.String())
}
