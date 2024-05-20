package ast

import (
	"fmt"
	"io"
	"strings"
)

func Print(out io.Writer, ast Node) {
	p := printer{
		out:       out,
		depth:     0,
		yieldNext: false,
	}
	ast.Accept(&p)
}

type printer struct {
	NullVisitor
	out       io.Writer
	depth     int
	yieldNext bool
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
	if p.yieldNext {
		p.print("yield ")
		p.yieldNext = false
	}
}

func (p *printer) VisitStatementList(stmtList *StatementList) any {
	for _, stmt := range stmtList.Stmts {
		stmt.Accept(p)
	}
	return nil
}

func (p *printer) VisitBlock(b *Block) any {
	p.indent()
	p.println("Block {")
	p.depth += 1

	p.visitBlockInternal(b)

	p.depth -= 1
	p.indent()
	p.println("}")

	return nil
}

func (p *printer) visitBlockInternal(b *Block) {
	for _, stmt := range b.StatementList.Stmts {
		stmt.Accept(p)
	}
	if b.Yields() {
		p.yieldNext = true
		b.FinalExpr.Accept(p)
	}
}

func (p *printer) VisitBinding(b *Binding) any {
	p.indent()
	p.println("Binding(")
	p.depth += 1

	b.Name.Accept(p)
	b.Value.Accept(p)

	p.depth -= 1
	p.indent()
	p.println(")")
	return nil
}

func (p *printer) VisitFunctionDefinition(defn *FunctionDefinition) any {
	p.indent()
	p.printf("FunctionDefinition(%s, (", defn.Name.Name)
	if len(defn.Parameters) > 0 {
		p.print(defn.Parameters[0].Name)
		for _, param := range defn.Parameters[1:] {
			fmt.Fprintf(p.out, ", %s", param.Name)
		}
	}
	p.println(")) {")
	p.depth += 1

	p.visitBlockInternal(defn.Body)

	p.depth -= 1
	p.indent()
	p.println("}")

	return nil
}

func (p *printer) VisitCallExpr(call *CallExpr) any {
	p.indent()
	p.printf("Call(%s\n", call.Callee.Name)
	p.depth += 1

	for _, arg := range call.Arguments {
		arg.Accept(p)
	}

	p.depth -= 1
	p.indent()
	p.println(")")

	return nil
}

func (p *printer) VisitOpExpr(op *OpExpr) any {
	p.indent()
	p.printf("OpExpr(%s\n", op.Operation.String())
	p.depth += 1

	op.Left.Accept(p)
	op.Right.Accept(p)

	p.depth -= 1
	p.indent()
	p.println(")")

	return nil
}

func (p *printer) VisitIfExpression(ifExpr *IfExpression) any {
	p.indent()
	p.println("IfExpression(")
	p.depth += 1

	ifExpr.Condition.Accept(p)
	ifExpr.ThenExpr.Accept(p)
	ifExpr.ElseExpr.Accept(p)

	p.depth -= 1
	p.indent()
	p.println(")")

	return nil
}

func (p *printer) VisitNameAccess(access *NameAccess) any {
	p.indent()
	p.printf("NameAccess(%s)\n", access.Name.Name)
	return nil
}

func (p *printer) VisitLiteralInt(n *LiteralInt) any {
	p.indent()
	p.printf("LiteralInt(%d)\n", n.Value)
	return nil
}

func (p *printer) VisitLiteralFloat(n *LiteralFloat) any {
	p.indent()
	p.printf("LiteralFloat(%f)\n", n.Value)
	return nil
}

func (p *printer) VisitLiteralString(s *LiteralString) any {
	p.indent()
	p.printf("LiteralString(%s)\n", s.Value)
	return nil
}

func (p *printer) VisitLiteralBool(b *LiteralBool) any {
	p.indent()
	p.printf("LiteralBool(%v)\n", b.Value)
	return nil
}
