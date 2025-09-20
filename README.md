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
    
    // Create lexer and parser with the plugin
    lb := lexer.NewBuilder()
    p := parser.NewBuilder(lb).Install(tryparser.Plugin).Build(input)
    
    // Parse the program
    ast, err := p.ParseProgram()
    if err != nil {
        panic(err)
    }
    fmt.Println(ast.String())
}
```
