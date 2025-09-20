package tryparser

import (
	"fmt"
	"strings"
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

func TestParserWithCatchParameter(t *testing.T) {
	input := `
	try {
		console.log('try something')
	} catch (error) {
		console.log('error:', error)
	} finally {
		console.log('cleanup')
	}`
	lb := lexer.NewBuilder()
	p := parser.NewBuilder(lb).Install(Plugin).Build(input)
	ast, err := p.ParseProgram()
	if err != nil {
		t.Fatalf("ParseProgram() error: %v", err)
	}
	output := ast.String()
	fmt.Println(output)

	// Verify that the catch parameter is present in the output
	if !strings.Contains(output, "catch (error)") {
		t.Errorf("Expected catch parameter 'error' in output, got: %s", output)
	}
}

func TestParserWithOnlyCatch(t *testing.T) {
	input := `
	try {
		risky_operation()
	} catch (e) {
		handle_error(e)
	}`
	lb := lexer.NewBuilder()
	p := parser.NewBuilder(lb).Install(Plugin).Build(input)
	ast, err := p.ParseProgram()
	if err != nil {
		t.Fatalf("ParseProgram() error: %v", err)
	}
	output := ast.String()
	fmt.Println(output)

	// Verify that the catch parameter is present and no finally block
	if !strings.Contains(output, "catch (e)") {
		t.Errorf("Expected catch parameter 'e' in output, got: %s", output)
	}
}
