package main

import (
	// "fmt"
	// "github.com/antlr4-go/antlr/v4"

	"fmt"
	"os"

	"github.com/SuneelFreimuth/vinland/src/ast"
	"github.com/SuneelFreimuth/vinland/src/parser"
)

// "bufio"
// "os"

func main() {
	input, err := os.ReadFile(".\\test\\parse\\test02.vin")
	if err != nil {
		fmt.Println(err)
		return
	}
	parseTree := parser.Parse(string(input))
	parser.Print(os.Stdout, parseTree)
	syntaxTree := ast.Build(parseTree)
	fmt.Println("AST:")
	ast.Format(os.Stdout, syntaxTree)
}
