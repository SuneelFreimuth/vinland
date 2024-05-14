package ast

import (
	"github.com/SuneelFreimuth/vinland/src/ast/types"
)


type Node interface {
	Accept(Visitor) any
	Children() []Node
	EnterNode(l Listener)
	ExitNode(l Listener)
}


type StmtExpr interface {
	Node
}


type Expression interface {
	StmtExpr
}


type Statement interface {
	StmtExpr
}


type Declaration interface {
	Node
}


type BaseNode struct {
	Type types.Type
}

func (bn *BaseNode) Accept(v Visitor) any {
	return nil
}

func (bn *BaseNode) Children() []Node {
	return nil
}

func (bn *BaseNode) EnterNode(l Listener) {}
func (bn *BaseNode) ExitNode(l Listener) {}


type DeclarationList struct {
	BaseNode
	Decls []Declaration
}

func NewDeclarationList(decls []Declaration) *DeclarationList {
	return &DeclarationList{Decls: decls}
}

func (dl *DeclarationList) Accept(v Visitor) any {
	return v.VisitDeclarationList(dl)
}

func (dl *DeclarationList) Children() []Node {
	children := make([]Node, len(dl.Decls))
	for i, decl := range dl.Decls {
		children[i] = decl
	}
	return children
}

func (dl *DeclarationList) EnterNode(l Listener) {
	l.EnterDeclarationList(dl)
}

func (dl *DeclarationList) ExitNode(l Listener) {
	l.ExitDeclarationList(dl)
}


type StatementList struct {
	BaseNode
	Stmts     []StmtExpr
	FinalExpr Expression
}

func NewStatementList(stmts []StmtExpr, finalExpr Expression) *StatementList {
	return &StatementList{
		Stmts:     stmts,
		FinalExpr: finalExpr,
	}
}

func (stmts *StatementList) Accept(v Visitor) any {
	return v.VisitStatementList(stmts)
}

func (stmts *StatementList) Children() []Node {
	children := make([]Node, len(stmts.Stmts))
	for i, stmt := range stmts.Stmts {
		children[i] = stmt
	}
	if stmts.FinalExpr != nil {
		children = append(children, stmts.FinalExpr)
	}
	return children
}

func (sl *StatementList) EnterNode(l Listener) {
	l.EnterStatementList(sl)
}

func (sl *StatementList) ExitNode(l Listener) {
	l.ExitStatementList(sl)
}


type FunctionDefinition struct {
	BaseNode
	Function   Symbol
	Parameters []Symbol
	Body       *StatementList
}

func NewFunctionDefinition(
	function Symbol,
	parameters []Symbol,
	body *StatementList,
) *FunctionDefinition {
	return &FunctionDefinition{
		Function:   function,
		Parameters: parameters,
		Body:       body,
	}
}

func (defn *FunctionDefinition) Accept(v Visitor) any {
	return v.VisitFunctionDefinition(defn)
}

func (defn *FunctionDefinition) Children() []Node {
	return defn.Body.Children()
}

func (defn *FunctionDefinition) EnterNode(l Listener) {
	l.EnterFunctionDefinition(defn)
}

func (defn *FunctionDefinition) ExitNode(l Listener) {
	l.ExitFunctionDefinition(defn)
}


type Binding struct {
	BaseNode
	Name  Expression
	Value Expression
}

func NewBinding(location, value Expression) *Binding {
	return &Binding{
		Name:  location,
		Value: value,
	}
}

func (b *Binding) Accept(v Visitor) any {
	return v.VisitBinding(b)
}

func (b *Binding) Children() []Node {
	return []Node{b.Name, b.Value}
}

func (b *Binding) EnterNode(l Listener) {
	l.EnterBinding(b)
}

func (b *Binding) ExitNode(l Listener) {
	l.ExitBinding(b)
}


type LiteralInt struct {
	BaseNode
	Value int64
}

func NewLiteralInt(value int64) *LiteralInt {
	return &LiteralInt{Value: value}
}

func (li *LiteralInt) Accept(v Visitor) any {
	return v.VisitLiteralInt(li)
}

func (li *LiteralInt) EnterNode(l Listener) {
	l.EnterLiteralInt(li)
}

func (li *LiteralInt) ExitNode(l Listener) {
	l.ExitLiteralInt(li)
}


type LiteralFloat struct {
	BaseNode
	Value float64
}

func NewLiteralFloat(value float64) *LiteralFloat {
	return &LiteralFloat{Value: value}
}

func (lf *LiteralFloat) Accept(v Visitor) any {
	return v.VisitLiteralFloat(lf)
}

func (lf *LiteralFloat) EnterNode(l Listener) {
	l.EnterLiteralFloat(lf)
}

func (lf *LiteralFloat) ExitNode(l Listener) {
	l.ExitLiteralFloat(lf)
}


type LiteralString struct {
	BaseNode
	Value string
}

func NewLiteralString(value string) *LiteralString {
	return &LiteralString{Value: value}
}

func (ls *LiteralString) Accept(v Visitor) any {
	return v.VisitLiteralString(ls)
}

func (ls *LiteralString) EnterNode(l Listener) {
	l.EnterLiteralString(ls)
}

func (ls *LiteralString) ExitNode(l Listener) {
	l.ExitLiteralString(ls)
}


type LiteralBool struct {
	BaseNode
	Value bool
}

func NewLiteralBool(value bool) *LiteralBool {
	return &LiteralBool{Value: value}
}

func (lb *LiteralBool) Accept(v Visitor) any {
	return v.VisitLiteralBool(lb)
}

func (lb *LiteralBool) EnterNode(l Listener) {
	l.EnterLiteralBool(lb)
}

func (lb *LiteralBool) ExitNode(l Listener) {
	l.ExitLiteralBool(lb)
}


type IfExpression struct {
	BaseNode
	Condition Expression
	ThenExpr  Expression
	ElseExpr  Expression
}

func NewIfExpression(condition, thenExpr, elseExpr Expression) *IfExpression {
	return &IfExpression{
		Condition: condition,
		ThenExpr:  thenExpr,
		ElseExpr:  elseExpr,
	}
}

func (ifExpr *IfExpression) Accept(v Visitor) any {
	return v.VisitIfExpression(ifExpr)
}

func (ifExpr *IfExpression) Children() []Node {
	return []Node{
		ifExpr.Condition,
		ifExpr.ThenExpr,
		ifExpr.ElseExpr,
	}
}

func (ifExpr *IfExpression) EnterNode(l Listener) {
	l.EnterIfExpression(ifExpr)
}

func (ifExpr *IfExpression) ExitNode(l Listener) {
	l.ExitIfExpression(ifExpr)
}


type OpExpr struct {
	BaseNode
	Operation Operation
	Left      Expression
	Right     Expression
}

type Operation int

const (
	Add Operation = iota
	Sub
	Mul
	Div
	Mod
	Not
	Or
	And
	GE
	GT
	LE
	LT
	EQ
	NE
)

func OpFromString(s string) Operation {
	switch s {
	case "+":
		return Add
	case "-":
		return Sub
	case "*":
		return Mul
	case "/":
		return Div
	case "%":
		return Mod
	case "!":
		return Not
	case "||":
		return Or
	case "&&":
		return And
	case ">=":
		return GE
	case ">":
		return GT
	case "<=":
		return LE
	case "<":
		return LT
	case "==":
		return EQ
	default:
		return NE
	}
}

func (op Operation) String() string {
	switch op {
	case Add:
		return "+"
	case Sub:
		return "-"
	case Mul:
		return "*"
	case Div:
		return "/"
	case Mod:
		return "%"
	case Or:
		return "||"
	case And:
		return "&&"
	case GE:
		return ">="
	case GT:
		return ">"
	case LE:
		return "<="
	case LT:
		return "<"
	case EQ:
		return "=="
	default:
		return "!="
	}
}

func NewOpExpr(op Operation, left Expression, right Expression) *OpExpr {
	return &OpExpr{
		Operation: op,
		Left:      left,
		Right:     right,
	}
}

func (expr *OpExpr) Accept(v Visitor) any {
	return v.VisitOpExpr(expr)
}

func (expr *OpExpr) Children() []Node {
	return []Node{
		expr.Left,
		expr.Right,
	}
}

func (expr *OpExpr) EnterNode(l Listener) {
	l.EnterOpExpr(expr)
}

func (expr *OpExpr) ExitNode(l Listener) {
	l.ExitOpExpr(expr)
}


type CallExpr struct {
	BaseNode
	Callee    Symbol
	Arguments []Expression
}

func NewCallExpr(callee Symbol, arguments []Expression) *CallExpr {
	return &CallExpr{
		Callee:    callee,
		Arguments: arguments,
	}
}

func (call *CallExpr) Accept(v Visitor) any {
	return v.VisitCallExpr(call)
}

func (call *CallExpr) Children() []Node {
	children := make([]Node, len(call.Arguments))
	for i, arg := range call.Arguments {
		children[i] = arg
	}
	return children
}

func (call *CallExpr) EnterNode(l Listener) {
	l.EnterCallExpr(call)
}

func (call *CallExpr) ExitNode(l Listener) {
	l.ExitCallExpr(call)
}


type NameAccess struct {
	BaseNode
	Name Symbol
}

func NewNameAccess(name Symbol) *NameAccess {
	return &NameAccess{Name: name}
}

func (na *NameAccess) Accept(v Visitor) any {
	return v.VisitNameAccess(na)
}

func (na *NameAccess) EnterNode(l Listener) {
	l.EnterNameAccess(na)
}

func (na *NameAccess) ExitNode(l Listener) {
	l.ExitNameAccess(na)
}
