package main

import (
	// "encoding/xml"
	// "fmt"
	// "os"

	// "github.com/SuneelFreimuth/vinland/src/ast"
	// "github.com/SuneelFreimuth/vinland/src/parser"
	"os"

	"github.com/SuneelFreimuth/vinland/src/eval"
)

func main() {
	i := eval.NewInterpreter()
	i.Run(os.Stdout, os.Stdin)
}
