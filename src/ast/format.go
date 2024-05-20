package ast

import (
	"fmt"
	"io"
)

func Format(out io.Writer, ast Node) {
	ast.Accept(&formatter{out: out})
}

type formatter struct {
	NullVisitor
	out io.Writer
}

func (f *formatter) VisitStatementList(stmtList *StatementList) any {
	if len(stmtList.Stmts) > 0 {
		stmtList.Stmts[0].Accept(f)
		for _, stmt := range stmtList.Stmts[1:] {
			fmt.Fprint(f.out, "\n")
			stmt.Accept(f)
		}
		fmt.Fprint(f.out, "\n")
	}
	return nil
}

func (f *formatter) VisitBinding(b *Binding) any {
	b.Name.Accept(f)
	fmt.Fprint(f.out, " = ")
	b.Value.Accept(f)
	return nil
}

func (f *formatter) VisitFunctionDefinition(defn *FunctionDefinition) any {
	fmt.Fprintf(f.out, "fn %s(", defn.Name.Name)
	if len(defn.Parameters) > 0 {
		fmt.Fprint(f.out, defn.Parameters[0].Name)
		for _, param := range defn.Parameters[1:] {
			fmt.Fprintf(f.out, ", %s", param.Name)
		}
	}
	fmt.Fprint(f.out, ") {\n")
	defn.Body.Accept(f)
	return nil
}

func (f *formatter) VisitCallExpr(call *CallExpr) any {
	fmt.Fprintf(f.out, "(%s ", call.Callee.Name)
	if len(call.Arguments) > 0 {
		call.Arguments[0].Accept(f)
		for _, arg := range call.Arguments[1:] {
			fmt.Fprint(f.out, " ")
			arg.Accept(f)
		}
	}
	fmt.Fprint(f.out, ")")
	return nil
}

func (f *formatter) VisitOpExpr(op *OpExpr) any {
	fmt.Fprint(f.out, "(")
	op.Left.Accept(f)
	fmt.Fprintf(f.out, " %s ", op.Operation.String())
	op.Right.Accept(f)
	fmt.Fprint(f.out, ")")
	return nil
}

func (f *formatter) VisitIfExpression(ifExpr *IfExpression) any {
	fmt.Fprint(f.out, "if ")
	ifExpr.Condition.Accept(f)
	fmt.Fprint(f.out, " then ")
	ifExpr.ThenExpr.Accept(f)
	fmt.Fprint(f.out, " else ")
	ifExpr.ElseExpr.Accept(f)
	return nil
}

func (f *formatter) VisitNameAccess(access *NameAccess) any {
	fmt.Fprint(f.out, access.Name.Name)
	return nil
}

func (f *formatter) VisitLiteralInt(n *LiteralInt) any {
	fmt.Fprintf(f.out, "%d", n.Value)
	return nil
}

func (f *formatter) VisitLiteralFloat(n *LiteralFloat) any {
	fmt.Fprintf(f.out, "%f", n.Value)
	return nil
}

func (f *formatter) VisitLiteralString(s *LiteralString) any {
	fmt.Fprintf(f.out, "\"%f\"", s.Value)
	return nil
}

func (f *formatter) VisitLiteralBool(b *LiteralBool) any {
	fmt.Fprintf(f.out, "%b", b.Value)
	return nil
}
