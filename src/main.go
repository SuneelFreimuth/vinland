package main

import (
	"fmt"
	"os"

	"github.com/SuneelFreimuth/vinland/src/ast"
	"github.com/SuneelFreimuth/vinland/src/parser"
)

func main() {
	// input, err := os.ReadFile(".\\test\\parse\\test02.vin")
	input, err := os.ReadFile("./test/parse/test02.vin")
	if err != nil {
		fmt.Println(err)
		return
	}
	parseTree := parser.Parse(string(input))
	parser.Print(os.Stdout, parseTree)
	syntaxTree := ast.Build(parseTree)
	fmt.Println("========== BEGIN AST =========")
	ast.Print(os.Stdout, syntaxTree)
	fmt.Println("========== END AST =========")
}
