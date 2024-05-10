package ast

import (
	"fmt"
)

func Print(ast Node) {
	ast.Accept(&Printer{})
}

type Printer struct {}

func (p *Printer) VisitOperation(op *Operation) any {
	fmt.Print("(")

	op.Children[0].Accept(p)
	op_ := OpToString(op.Op)
	fmt.Printf(" %s ", op_)
	op.Children[1].Accept(p)

	fmt.Print(")")
	
	return nil
}

func (p *Printer) VisitNumber(n *Number) any {
	fmt.Print(n.Value)
	return nil
}

func (p *Printer) VisitCall(c *Call) any {
	fmt.Print(FunctionToString(c.Function))
	return nil
}