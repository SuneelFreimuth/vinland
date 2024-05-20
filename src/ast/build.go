package ast

import (
	"strconv"

	"github.com/SuneelFreimuth/vinland/src/parser"
)

func Build(start parser.IProgramContext) Node {
	return start.Accept(NewBuilder()).(Node)
}


type Builder struct {
	*parser.BaseVinlandVisitor

	SymbolTable *SymbolTable
}

func NewBuilder() *Builder {
	return &Builder{
		SymbolTable: NewSymbolTable(),
	}
}


func (b *Builder) VisitProgram(ctx *parser.ProgramContext) any {
	return ctx.StmtList().Accept(b)
}

func (b *Builder) VisitBlock(ctx *parser.BlockContext) any {
	b.SymbolTable.Enter()
	defer b.SymbolTable.Exit()
	stmtList := ctx.StmtList().Accept(b).(*StatementList)

	var finalExpr Expression
	if expr0 := ctx.Expr0(); expr0 != nil {
		finalExpr = expr0.Accept(b).(Expression)
	}

	return NewBlock(stmtList, finalExpr)
}

func (b *Builder) VisitStmtList(ctx *parser.StmtListContext) any {
	stmts := make([]Statement, len(ctx.AllStmt()))
	for i, stmt := range ctx.AllStmt() {
		stmts[i] = stmt.Accept(b).(Statement)
	}
	return NewStatementList(stmts)
}

func (b *Builder) VisitStmt(ctx *parser.StmtContext) any {
	if binding := ctx.Binding(); binding != nil {
		return binding.Accept(b)
	}
	if expr := ctx.Expr0(); expr != nil {
		return expr.Accept(b)
	}
	return ctx.FunctionDefinition().Accept(b)
}

func (b *Builder) VisitBinding(ctx *parser.BindingContext) any {
	location := ctx.Identifier().GetText()
	symbol, _ := b.SymbolTable.Put(location)
	name := NewNameAccess(symbol)
	value := ctx.Expr0().Accept(b).(Expression);
	return NewBinding(name, value)
}

func (b *Builder) VisitFunctionDefinition(ctx *parser.FunctionDefinitionContext) any {
	name, _ := b.SymbolTable.Put(ctx.Identifier().GetText())

	b.SymbolTable.Enter()

	params := ctx.ParamList().AllParam()
	paramSymbols := make([]Symbol, len(params))
	for i, param := range params {
		paramSymbols[i], _ = b.SymbolTable.Put(param.GetText())
	}

	// Avoid creating another new scope since the function's body is a block.
	stmtList := ctx.Block().StmtList().Accept(b).(*StatementList)
	var finalExpr Expression
	if expr0 := ctx.Block().Expr0(); expr0 != nil {
		finalExpr = expr0.Accept(b).(Expression)
	}
	block := NewBlock(stmtList, finalExpr)

	b.SymbolTable.Exit()

	return NewFunctionDefinition(name, paramSymbols, block)
}

func (b *Builder) VisitExpr0(ctx *parser.Expr0Context) any {
	if expr1 := ctx.Expr1(); expr1 != nil {
		return expr1.Accept(b)
	}
	if ifExpr := ctx.IfExpr(); ifExpr != nil {
		return ifExpr.Accept(b)
	}
	return ctx.Block().Accept(b)
}

func (b *Builder) VisitIfExpr(ctx *parser.IfExprContext) any {
	return NewIfExpression(
		ctx.Expr0(0).Accept(b).(Expression),
		ctx.Expr0(1).Accept(b).(Expression),
		ctx.Expr0(2).Accept(b).(Expression),
	)
}

func (b *Builder) VisitExpr1(ctx *parser.Expr1Context) any {
	if op0 := ctx.Op0(); op0 != nil {
		return NewOpExpr(
			OpFromString(op0.GetText()),
			ctx.Expr2(0).Accept(b).(Expression),
			ctx.Expr2(1).Accept(b).(Expression),
		)
	}
	return ctx.Expr2(0).Accept(b)
}

func (b *Builder) VisitExpr2(ctx *parser.Expr2Context) any {
	if op1 := ctx.Op1(); op1 != nil {
		return NewOpExpr(
			OpFromString(op1.GetText()),
			ctx.Expr3(0).Accept(b).(Expression),
			ctx.Expr3(1).Accept(b).(Expression),
		)
	}
	return ctx.Expr3(0).Accept(b)
}

func (b *Builder) VisitExpr3(ctx *parser.Expr3Context) any {
	if op2 := ctx.Op2(); op2 != nil {
		return NewOpExpr(
			OpFromString(op2.GetText()),
			ctx.Expr3().Accept(b).(Expression),
			ctx.Expr4().Accept(b).(Expression),
		)
	}
	return ctx.Expr4().Accept(b)
}

func (b *Builder) VisitExpr4(ctx *parser.Expr4Context) any {
	if expr4 := ctx.Expr4(); expr4 != nil {
		return NewOpExpr(Not, expr4.Accept(b).(Expression), nil)
	}
	if expr0 := ctx.Expr0(); expr0 != nil {
		return expr0.Accept(b)
	}
	if name := ctx.Name(); name != nil {
		n, exists := b.SymbolTable.Lookup(name.GetText())
		if !exists {
			// Error: name "n" is not defined
		}
		return NewNameAccess(n)
	}
	if callExpr := ctx.CallExpr(); callExpr != nil {
		return callExpr.Accept(b)
	}
	return ctx.Literal().Accept(b)
}

func (b *Builder) VisitLiteral(ctx *parser.LiteralContext) any {
	if l := ctx.IntLiteral(); l != nil {
		return l.Accept(b)
	}
	if l := ctx.FloatLiteral(); l != nil {
		n, _ := strconv.ParseFloat(l.GetText(), 64)
		return NewLiteralFloat(n)
	}
	if l := ctx.BooleanLiteral(); l != nil {
		return NewLiteralBool(l.GetText() == "true")
	}
	s := ctx.StringLiteral().GetText()
	return NewLiteralString(s[1:len(s)-1])
}

func (b *Builder) VisitIntLiteral(ctx *parser.IntLiteralContext) any {
	if dec := ctx.DecimalNumber(); dec != nil {
		n, _ := strconv.ParseInt(dec.GetText(), 10, 64)
		return NewLiteralInt(n)
	}
	if hex := ctx.HexadecimalNumber(); hex != nil {
		n, _ := strconv.ParseInt(hex.GetText()[2:], 16, 64)
		return NewLiteralInt(n)
	}
	// Binary literal
	n, _ := strconv.ParseInt(ctx.GetText()[2:], 2, 64)
	return NewLiteralInt(n)
}

func (b *Builder) VisitCallExpr(ctx *parser.CallExprContext) any {
	callee, exists := b.SymbolTable.Lookup(ctx.Identifier().GetText())
	if !exists {
		// Error: function is not defined
		exists = exists
	}

	expr0s := ctx.AllExpr0()
	args := make([]Expression, len(expr0s))
	for i, expr0 := range expr0s {
		args[i] = expr0.Accept(b).(Expression)
	}

	return NewCallExpr(callee, args)
}
