package ast

import (
	"fmt"
	"io"
)


func Print(out io.Writer, ast Node) {
	ast.Walk(&printer{out: out, depth: -1})
}


type printer struct {
	NullListener
	out io.Writer
	depth int
}

func (p *printer) print(a ...any) {
	fmt.Fprint(p.out, a...)
}

func (p *printer) println(a ...any) {
	fmt.Println(a...)
}

func (p *printer) printf(format string, a ...any) {
	fmt.Fprintf(p.out, format, a...)
}

func (p *printer) EnterEveryNode(declList *DeclarationList) {
	p.depth += 1
}

func (p *printer) ExitEveryNode(declList *DeclarationList) {
	p.depth -= 1
}

func (p *printer) EnterDeclarationList(declList *DeclarationList) {
	p.print("DeclarationList {")
}

func (p *printer) ExitDeclarationList(declList *DeclarationList) {
	p.print("}")
}

func (p *printer) EnterStatementList(stmtList *StatementList) {
	if len(stmtList.Stmts) > 0 {
		stmtList.Stmts[0].Accept(p)
		for _, stmt := range stmtList.Stmts[1:] {
			p.print("\n")
			stmt.Accept(p)
		}
		if stmtList.FinalExpr != nil {
			stmtList.FinalExpr.Accept(p)
		}
		p.print("\n")
	}
}

func (p *printer) EnterBinding(b *Binding) {
	p.print("Binding(")
	b.Name.Accept(p)
	p.print(", ")
	b.Value.Accept(p)
	p.print(")")
}

func (p *printer) EnterFunctionDefinition(defn *FunctionDefinition) any {
	fmt.Fprintf(p.out, "FunctionDefinition(%s, (", defn.Function.Name)
	if len(defn.Parameters) > 0 {
		p.print(defn.Parameters[0].Name)
		for _, param := range defn.Parameters[1:] {
			fmt.Fprintf(p.out, ", %s", param.Name)
		}
	}
	p.print(")) {\n")
	defn.Body.Accept(p)
	return nil
}

func (p *printer) EnterCallExpr(call *CallExpr) any {
	fmt.Fprintf(p.out, "Call(%s ", call.Callee.Name)
	if len(call.Arguments) > 0 {
		call.Arguments[0].Accept(p)
		for _, arg := range call.Arguments[1:] {
			p.print(" ")
			arg.Accept(p)
		}
	}
	p.print(")")
	return nil
}

func (p *printer) EnterOpExpr(op *OpExpr) any {
	p.print("(")
	op.Left.Accept(p)
	fmt.Fprintf(p.out, " %s ", op.Operation.String())
	op.Right.Accept(p)
	p.print(")")
	return nil
}

func (p *printer) EnterIfExpression(ifExpr *IfExpression) any {
	p.print("if ")
	ifExpr.Condition.Accept(p)
	p.print(" then ")
	ifExpr.ThenExpr.Accept(p)
	p.print(" else ")
	ifExpr.ElseExpr.Accept(p)
	return nil
}

func (p *printer) EnterNameAccess(access *NameAccess) any {
	p.print(access.Name.Name)
	return nil
}

func (p *printer) EnterLiteralInt(n *LiteralInt) any {
	fmt.Fprintf(p.out, "%d", n.Value)
	return nil
}

func (p *printer) EnterLiteralFloat(n *LiteralFloat) any {
	fmt.Fprintf(p.out, "%p", n.Value)
	return nil
}

func (p *printer) EnterLiteralString(s *LiteralString) any {
	fmt.Fprintf(p.out, "\"%p\"", s.Value)
	return nil
}

func (p *printer) EnterLiteralBool(b *LiteralBool) any {
	fmt.Fprintf(p.out, "%b", b.Value)
	return nil
}
