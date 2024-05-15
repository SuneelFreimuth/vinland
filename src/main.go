package main

import (
	"encoding/xml"
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
	var syntaxTree xml.Marshaler
	syntaxTree = ast.Build(parseTree)
	fmt.Println("========== BEGIN AST =========")
	encoded, _ := xml.MarshalIndent(syntaxTree, "", "  ")
	fmt.Println(string(encoded))
	fmt.Println("========== END AST =========")
}
