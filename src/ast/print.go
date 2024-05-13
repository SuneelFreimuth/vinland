package ast

import (
	"fmt"
	"io"
)

func Print(out io.Writer, ast Node) {
	ast.Accept(&printer{out: out})
}

type printer struct {
	NullVisitor
	out io.Writer
}

func (p *printer) VisitDeclarationList(declList *DeclarationList) any {
	if len(declList.Decls) > 0 {
		declList.Decls[0].Accept(p)
		for _, decl := range declList.Decls[1:] {
			fmt.Fprint(p.out, "\n")
			decl.Accept(p)
		}
	}
	return nil
}

func (p *printer) VisitStatementList(stmtList *StatementList) any {
	if len(stmtList.Stmts) > 0 {
		stmtList.Stmts[0].Accept(p)
		for _, stmt := range stmtList.Stmts[1:] {
			fmt.Fprint(p.out, "\n")
			stmt.Accept(p)
		}
		if stmtList.FinalExpr != nil {
			stmtList.FinalExpr.Accept(p)
		}
		fmt.Fprint(p.out, "\n")
	}
	return nil
}

func (p *printer) VisitBinding(b *Binding) any {
	b.Name.Accept(p)
	fmt.Fprint(p.out, " = ")
	b.Value.Accept(p)
	return nil
}

func (p *printer) VisitFunctionDefinition(defn *FunctionDefinition) any {
	fmt.Fprintf(p.out, "fn %s(", defn.Function.Name)
	if len(defn.Parameters) > 0 {
		fmt.Fprint(p.out, defn.Parameters[0].Name)
		for _, param := range defn.Parameters[1:] {
			fmt.Fprintf(p.out, ", %s", param.Name)
		}
	}
	fmt.Fprint(p.out, ") {\n")
	defn.Body.Accept(p)
	return nil
}

func (p *printer) VisitCallExpr(call *CallExpr) any {
	fmt.Fprintf(p.out, "%s(", call.Callee.Name)
	if len(call.Arguments) > 0 {
		fmt.Fprint(p.out, call.Arguments[0])
		for _, arg := range call.Arguments[1:] {
			fmt.Fprint(p.out, ", ")
			arg.Accept(p)
		}
	}
	fmt.Fprint(p.out, ")")
	return nil
}

func (p *printer) VisitOpExpr(op *OpExpr) any {
	fmt.Fprint(p.out, "(")
	op.Left.Accept(p)
	fmt.Fprintf(p.out, " %s ", op.Operation.String())
	op.Right.Accept(p)
	fmt.Fprint(p.out, ")")
	return nil
}

func (p *printer) VisitIfExpression(ifExpr *IfExpression) any {
	fmt.Fprint(p.out, "if ")
	ifExpr.Condition.Accept(p)
	fmt.Fprint(p.out, " then ")
	ifExpr.ThenExpr.Accept(p)
	fmt.Fprint(p.out, " else ")
	ifExpr.ElseExpr.Accept(p)
	return nil
}

func (p *printer) VisitNameAccess(access *NameAccess) any {
	fmt.Fprint(p.out, access.Name)
	return nil
}

func (p *printer) VisitLiteralInt(n *LiteralInt) any {
	fmt.Fprintf(p.out, "%d", n.Value)
	return nil
}

func (p *printer) VisitLiteralFloat(n *LiteralFloat) any {
	fmt.Fprintf(p.out, "%f", n.Value)
	return nil
}

func (p *printer) VisitLiteralString(s *LiteralString) any {
	fmt.Fprintf(p.out, "\"%f\"", s.Value)
	return nil
}

func (p *printer) VisitLiteralBool(b *LiteralBool) any {
	fmt.Fprintf(p.out, "%b", b.Value)
	return nil
}
