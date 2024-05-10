package ast

func Evaluate(root Node) float64 {
	return root.Accept(&Evaluator{}).(float64)
}

type Evaluator struct {}

func (e *Evaluator) VisitOperation(op *Operation) any {
	left := op.Children[0].Accept(e).(float64)
	right := op.Children[1].Accept(e).(float64)
	switch op.Op {
	case Add:
		return left + right
	case Sub:
		return left - right
	case Mul:
		return left * right
	default:
		return left / right
	}
}

func (e *Evaluator) VisitCall(c *Call) any {
	return nil
}

func (e *Evaluator) VisitNumber(n *Number) any {
	return n.Value
}