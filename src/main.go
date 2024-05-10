package main

import (
	"fmt"
	"bufio"
	"os"

	"github.com/antlr4-go/antlr/v4"

	"github.com/SuneelFreimuth/vinland/src/ast"
	"github.com/SuneelFreimuth/vinland/src/parser"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		parseTree := parse(line)
		tree := ast.Build(parseTree)
		ast.Print(tree)
		fmt.Println()
		fmt.Println(ast.Evaluate(tree))
	}
}

func parse(expr string) parser.IProgramContext {
	is := antlr.NewInputStream(expr)

	// Create the Lexer
	lexer := parser.NewVinlandLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewVinlandParser(stream)

	return p.Program()
}