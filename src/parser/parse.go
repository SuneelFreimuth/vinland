package parser

import (
	"fmt"
	"io"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

func Parse(expr string) (IProgramContext, SyntaxError) {
	is := antlr.NewInputStream(expr)

	// Create the Lexer
	lexer := NewVinlandLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewVinlandParser(stream)

	return p.Program(), nil
}


type SyntaxError interface{}


func Print(out io.Writer, program IProgramContext) {
	p := newPrinter(out)
	antlr.ParseTreeWalkerDefault.Walk(p, program)
}

type printer struct {
	*BaseVinlandListener

	out io.Writer
	depth int
}

func newPrinter(out io.Writer) *printer {
	return &printer{out: out, depth: -1}
}

func (p *printer) print(rule string) {
	indent := strings.Repeat("  ", p.depth)
	fmt.Fprintf(p.out, "%s%s\n", indent, rule)
}

func (p *printer) EnterEveryRule(ctx antlr.ParserRuleContext) {
	p.depth += 1
}

func (p *printer) ExitEveryRule(ctx antlr.ParserRuleContext) {
	p.depth -= 1
}

func (p *printer) VisitErrorNode(node antlr.ErrorNode) {
	p.print(node.GetSymbol().GetText())
}

func (p *printer) EnterProgram(ctx *ProgramContext) {
	p.print("program")
}

func (p *printer) EnterDeclList(ctx *DeclListContext) {
	p.print("declList")
}

func (p *printer) EnterDecl(ctx *DeclContext) {
	p.print("decl")
}

func (p *printer) EnterStmtBlock(ctx *StmtBlockContext) {
	p.print("stmtBlock")
}

func (p *printer) EnterStmtList(ctx *StmtListContext) {
	p.print("stmtList")
}

func (p *printer) EnterStmt(ctx *StmtContext) {
	p.print("stmt")
}

func (p *printer) EnterBinding(ctx *BindingContext) {
	p.print("binding")
}

func (p *printer) EnterFunctionDefinition(ctx *FunctionDefinitionContext) {
	p.print("functionDefinition")
}

func (p *printer) EnterParamList(ctx *ParamListContext) {
	p.print("paramList")
}

func (p *printer) EnterParam(ctx *ParamContext) {
	p.print("param")
}

func (p *printer) EnterOp0(ctx *Op0Context) {
	p.print("op0")
}

func (p *printer) EnterOp1(ctx *Op1Context) {
	p.print("op1")
}

func (p *printer) EnterOp2(ctx *Op2Context) {
	p.print("op2")
}

func (p *printer) EnterExpr0(ctx *Expr0Context) {
	p.print("expr0")
}

func (p *printer) EnterExpr1(ctx *Expr1Context) {
	p.print("expr1")
}

func (p *printer) EnterExpr2(ctx *Expr2Context) {
	p.print("expr2")
}

func (p *printer) EnterExpr3(ctx *Expr3Context) {
	p.print("expr3")
}

func (p *printer) EnterExpr4(ctx *Expr4Context) {
	p.print("expr4")
}

func (p *printer) EnterName(ctx *NameContext) {
	p.print("name")
}

func (p *printer) EnterCallExpr(ctx *CallExprContext) {
	p.print("callExpr")
}

func (p *printer) EnterIfExpr(ctx *IfExprContext) {
	p.print("ifExpr")
}

func (p *printer) EnterLiteral(ctx *LiteralContext) {
	p.print("literal")
}
