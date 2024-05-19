package ast

import (
	"fmt"
	"io"
	"strings"
)

func Print(out io.Writer, ast Node) {
	Walk(ast, &printer{out: out, depth: -1})
}

type printer struct {
	NullListener
	out   io.Writer
	depth int
}

func (p *printer) print(a ...any) {
	fmt.Fprint(p.out, a...)
}

func (p *printer) println(a ...any) {
	fmt.Fprint(p.out, a...)
	fmt.Fprint(p.out, "\n")
}

func (p *printer) printf(format string, a ...any) {
	fmt.Fprintf(p.out, format, a...)
}

func (p *printer) indent() {
	p.print(strings.Repeat("  ", p.depth))
}

func (p *printer) EnterEveryNode(_ Node) {
	p.depth += 1
	p.indent()
}

func (p *printer) ExitEveryNode(_ Node) {
	p.depth -= 1
}

func (p *printer) EnterDeclarationList(declList *DeclarationList) {
	p.println("DeclarationList {")
}

func (p *printer) ExitDeclarationList(declList *DeclarationList) {
	p.indent()
	p.println("}")
}

func (p *printer) EnterBinding(b *Binding) {
	p.println("Binding(")
}

func (p *printer) ExitBinding(b *Binding) {
	p.indent()
	p.println(")")
}

func (p *printer) EnterFunctionDefinition(defn *FunctionDefinition) {
	p.printf("FunctionDefinition(%s, (", defn.Name.Name)
	if len(defn.Parameters) > 0 {
		p.print(defn.Parameters[0].Name)
		for _, param := range defn.Parameters[1:] {
			fmt.Fprintf(p.out, ", %s", param.Name)
		}
	}
	p.println(")) {")
}

func (p *printer) ExitFunctionDefinition(defn *FunctionDefinition) {
	p.indent()
	p.println("}")
}

func (p *printer) EnterCallExpr(call *CallExpr) {
	p.printf("Call(%s\n", call.Callee.Name)
}

func (p *printer) ExitCallExpr(call *CallExpr) {
	p.indent()
	p.println(")")
}

func (p *printer) EnterOpExpr(op *OpExpr) {
	p.printf("OpExpr(%s\n", op.Operation.String())
}

func (p *printer) ExitOpExpr(op *OpExpr) {
	p.indent()
	p.println(")")
}

func (p *printer) EnterIfExpression(ifExpr *IfExpression) {
	p.println("IfExpression(")
}

func (p *printer) ExitIfExpression(ifExpr *IfExpression) {
	p.indent()
	p.println(")")
}

func (p *printer) EnterNameAccess(access *NameAccess) {
	p.printf("NameAccess(%s)\n", access.Name.Name)
}

func (p *printer) EnterLiteralInt(n *LiteralInt) {
	p.printf("LiteralInt(%d)\n", n.Value)
}

func (p *printer) EnterLiteralFloat(n *LiteralFloat) {
	p.printf("LiteralFloat(%f)\n", n.Value)
}

func (p *printer) EnterLiteralString(s *LiteralString) {
	p.printf("LiteralString(%s)\n", s.Value)
}

func (p *printer) EnterLiteralBool(b *LiteralBool) {
	p.printf("LiteralBool(%v)\n", b.Value)
}
