package ast


type Node interface {
	Accept(Visitor) any
}

type Visitor interface {
	VisitOperation(op *Operation) any
	VisitNumber(n *Number) any
	VisitCall(c *Call) any
}

type BaseNode struct {
	Children []Node
}

func (bn *BaseNode) SetChild(i int, n Node) {
	if i >= len(bn.Children) {
		newChildren := make([]Node, i + 1)
		copy(newChildren, bn.Children)
		bn.Children = newChildren
	}
	bn.Children[i] = n
}


type Number struct {
	BaseNode
	Value float64
}

func NewNumber(value float64) *Number {
	return &Number{Value: value}
}

func (n *Number) Accept(v Visitor) any {
	return v.VisitNumber(n)
}


type Operation struct {
	BaseNode
	Op Op
}

func NewOperation(op Op) *Operation {
	return &Operation{Op: op}
}

func (op *Operation) Accept(v Visitor) any {
	return v.VisitOperation(op)
}

type Op int

const (
	Add Op = iota
	Sub
	Mul
	Div
)

func OpFromString(s string) Op {
	switch s {
	case "+": return Add
	case "-": return Sub
	case "*": return Mul
	default: return Div
	}
}

func OpToString(op Op) string {
	switch op {
	case Add:
		return "+"
	case Sub:
		return "-"
	case Mul:
		return "*"
	case Div:
		return "/"
	default:
		return "%"
	}
}


type Function int

const (
	Sine Function = iota
	Cosine
	Floor
	Ceil
)

func FunctionToString(f Function) string {
	switch f {
	case Sine:
		return "sin"
	case Cosine:
		return "cos"
	case Floor:
		return "floor"
	default:
		return "ceil"
	}
}

type Call struct {
	BaseNode
	Function Function
}

func NewCall(f Function) *Call {
	return &Call{ Function: f }
}

func (c *Call) Accept(v Visitor) any {
	return v.VisitCall(c)
}