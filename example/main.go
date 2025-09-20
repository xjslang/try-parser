package main

import (
	"fmt"

	tryparser "github.com/xjslang/try-parser"
	"github.com/xjslang/xjs/lexer"
	"github.com/xjslang/xjs/parser"
)

func main() {
	// Example 1: try-catch-finally without catch parameter
	input1 := `
	try {
		console.log('attempting operation')
	} catch {
		console.log('operation failed')
	} finally {
		console.log('cleanup completed')
	}`

	fmt.Println("=== Example 1: try-catch-finally without parameter ===")
	lb1 := lexer.NewBuilder()
	p1 := parser.NewBuilder(lb1).Install(tryparser.Plugin).Build(input1)
	ast1, err1 := p1.ParseProgram()
	if err1 != nil {
		fmt.Printf("Error: %v\n", err1)
	} else {
		fmt.Println(ast1.String())
	}

	// Example 2: try-catch-finally with catch parameter
	input2 := `
	try {
		risky_operation()
	} catch (error) {
		console.log('Error occurred:', error.message)
	} finally {
		cleanup_resources()
	}`

	fmt.Println("\n=== Example 2: try-catch-finally with parameter ===")
	lb2 := lexer.NewBuilder()
	p2 := parser.NewBuilder(lb2).Install(tryparser.Plugin).Build(input2)
	ast2, err2 := p2.ParseProgram()
	if err2 != nil {
		fmt.Printf("Error: %v\n", err2)
	} else {
		fmt.Println(ast2.String())
	}

	// Example 3: try-catch only with parameter
	input3 := `
	try {
		dangerous_code()
	} catch (e) {
		handle_error(e)
	}`

	fmt.Println("\n=== Example 3: try-catch only with parameter ===")
	lb3 := lexer.NewBuilder()
	p3 := parser.NewBuilder(lb3).Install(tryparser.Plugin).Build(input3)
	ast3, err3 := p3.ParseProgram()
	if err3 != nil {
		fmt.Printf("Error: %v\n", err3)
	} else {
		fmt.Println(ast3.String())
	}
}
