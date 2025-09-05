# Try-Catch Parser Plugin for XJS

This is a plugin for the XJS transpiler that adds support for try-catch-finally statements, providing error handling capabilities similar to JavaScript.

## Usage

```go
package main

import (
    "fmt"

    "github.com/xjslang/xjs/lexer"
    "github.com/xjslang/xjs/parser"
    tryparser "github.com/xjslang/try-parser"
)

func main() {
    input := `
    try {
      doSomething();
    } catch (e) {
      handleError(e);
    } finally {
      console.log('done!')
    }`
    
    // Create lexer and parser
    l := lexer.New(input)
    p := parser.New(l)
    
    // Register the try-catch middleware and parse the program
    p.UseStatementHandler(tryparser.ParseStatement)
    ast := p.ParseProgram()
    fmt.Println(ast.String())
}
```
