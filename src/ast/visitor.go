package ast

type Visitor interface {
	VisitStatementList(*StatementList) any
	VisitBlock(*Block) any
	VisitBinding(*Binding) any
	VisitFunctionDefinition(*FunctionDefinition) any
	VisitCallExpr(*CallExpr) any
	VisitOpExpr(*OpExpr) any
	VisitIfExpression(*IfExpression) any
	VisitNameAccess(*NameAccess) any
	VisitLiteralInt(*LiteralInt) any
	VisitLiteralFloat(*LiteralFloat) any
	VisitLiteralString(*LiteralString) any
	VisitLiteralBool(*LiteralBool) any
}

type NullVisitor struct{}

func (v *NullVisitor) VisitStatementList(*StatementList) any {
	return nil
}

func (v *NullVisitor) VisitBlock(*Block) any {
	return nil
}

func (v *NullVisitor) VisitBinding(*Binding) any {
	return nil
}

func (v *NullVisitor) VisitFunctionDefinition(*FunctionDefinition) any {
	return nil
}

func (v *NullVisitor) VisitCallExpr(*CallExpr) any {
	return nil
}

func (v *NullVisitor) VisitOpExpr(*OpExpr) any {
	return nil
}

func (v *NullVisitor) VisitIfExpression(*IfExpression) any {
	return nil
}

func (v *NullVisitor) VisitNameAccess(*NameAccess) any {
	return nil
}

func (v *NullVisitor) VisitLiteralInt(*LiteralInt) any {
	return nil
}

func (v *NullVisitor) VisitLiteralFloat(*LiteralFloat) any {
	return nil
}

func (v *NullVisitor) VisitLiteralString(*LiteralString) any {
	return nil
}

func (v *NullVisitor) VisitLiteralBool(*LiteralBool) any {
	return nil
}
