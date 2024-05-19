package parser

import (
	"github.com/antlr4-go/antlr/v4"
)

func NewParser(expr string) *VinlandParser {
	is := antlr.NewInputStream(expr)
	lexer := NewVinlandLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewVinlandParser(stream)
	return p
}


type SyntaxError interface{}