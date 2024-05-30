package eval

import (
	"io"
	"bufio"

	_ "github.com/SuneelFreimuth/vinland/src/ast"
	_ "github.com/SuneelFreimuth/vinland/src/parser"
)


type Interpreter struct {
	Runtime *Runtime
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		Runtime: NewRuntime(),
	}
}

func (i *Interpreter) Run(out io.Writer, in io.Reader) error {
	scan := bufio.NewScanner(in)
	for scan.Scan() {
		line := scan.Text()
		i.EvalStatement(line)
	}
	return scan.Err()
}

func (i *Interpreter) EvalStatement(stmt string) Value {
	parser_ := parser.NewParser(stmt)
	stmt := parser_.Stmt()
	syntaxTree := ast.Build(parseTree)
	return i.Runtime.Eval(syntaxTree)
	return nil
}
