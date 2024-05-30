package eval

import (
	_ "fmt"
	_ "io"

	"github.com/SuneelFreimuth/vinland/src/ast"
)


type Runtime struct {
	Names []Scope
}

func NewRuntime() *Runtime {
	return &Runtime{Names: []Scope{}}
}

func (r *Runtime) Eval(n ast.Node) Value {
	e := evaluator{}
	return n.Accept(&e)
}


type evaluator struct {
	ast.NullVisitor
	Runtime *Runtime
}

func NewEvaluator(r *Runtime) *evaluator {
	return &evaluator{
		Runtime: r,
	}
}

// func (e *Evaluator) VisitStatementList(*StatementList) any {

// }

// func (e *Evaluator) VisitBinding(*Binding) any {

// }

// func (e *Evaluator) VisitFunctionDefinition(*FunctionDefinition) any {

// }

// func (e *Evaluator) VisitCallExpr(*CallExpr) any {

// }

// func (e *Evaluator) VisitOpExpr(*OpExpr) any {

// }

// func (e *Evaluator) VisitIfExpression(*IfExpression) any {

// }

// func (e *Evaluator) VisitNameAccess(*NameAccess) any {

// }

func (e *evaluator) VisitLiteralInt(li *ast.LiteralInt) any {
	return NewInt(li.Value)
}

// func (e *Evaluator) VisitLiteralFloat(*LiteralFloat) any {

// }

// func (e *Evaluator) VisitLiteralString(*LiteralString) any {

// }

// func (e *Evaluator) VisitLiteralBool(*LiteralBool) any {

// }


type Scope struct {
	Names map[Name]Value
}


type Name struct {
	Name string
}

func NewName(name string) Name {
	return Name{name}
}
