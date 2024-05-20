package ast

type Listener interface {
	EnterEveryNode(Node)
	ExitEveryNode(Node)
	EnterBlock(*Block)
	ExitBlock(*Block)
	EnterStatementList(*StatementList)
	ExitStatementList(*StatementList)
	EnterBinding(*Binding)
	ExitBinding(*Binding)
	EnterFunctionDefinition(*FunctionDefinition)
	ExitFunctionDefinition(*FunctionDefinition)
	EnterCallExpr(*CallExpr)
	ExitCallExpr(*CallExpr)
	EnterOpExpr(*OpExpr)
	ExitOpExpr(*OpExpr)
	EnterIfExpression(*IfExpression)
	ExitIfExpression(*IfExpression)
	EnterNameAccess(*NameAccess)
	ExitNameAccess(*NameAccess)
	EnterLiteralInt(*LiteralInt)
	ExitLiteralInt(*LiteralInt)
	EnterLiteralFloat(*LiteralFloat)
	ExitLiteralFloat(*LiteralFloat)
	EnterLiteralString(*LiteralString)
	ExitLiteralString(*LiteralString)
	EnterLiteralBool(*LiteralBool)
	ExitLiteralBool(*LiteralBool)
}

type NullListener struct{}

func (l *NullListener) EnterEveryNode(Node)                         {}
func (l *NullListener) ExitEveryNode(Node)                          {}
func (l *NullListener) EnterBlock(*Block)                           {}
func (l *NullListener) ExitBlock(*Block)                            {}
func (l *NullListener) EnterStatementList(*StatementList)           {}
func (l *NullListener) ExitStatementList(*StatementList)            {}
func (l *NullListener) EnterBinding(*Binding)                       {}
func (l *NullListener) ExitBinding(*Binding)                        {}
func (l *NullListener) EnterFunctionDefinition(*FunctionDefinition) {}
func (l *NullListener) ExitFunctionDefinition(*FunctionDefinition)  {}
func (l *NullListener) EnterCallExpr(*CallExpr)                     {}
func (l *NullListener) ExitCallExpr(*CallExpr)                      {}
func (l *NullListener) EnterOpExpr(*OpExpr)                         {}
func (l *NullListener) ExitOpExpr(*OpExpr)                          {}
func (l *NullListener) EnterIfExpression(*IfExpression)             {}
func (l *NullListener) ExitIfExpression(*IfExpression)              {}
func (l *NullListener) EnterNameAccess(*NameAccess)                 {}
func (l *NullListener) ExitNameAccess(*NameAccess)                  {}
func (l *NullListener) EnterLiteralInt(*LiteralInt)                 {}
func (l *NullListener) ExitLiteralInt(*LiteralInt)                  {}
func (l *NullListener) EnterLiteralFloat(*LiteralFloat)             {}
func (l *NullListener) ExitLiteralFloat(*LiteralFloat)              {}
func (l *NullListener) EnterLiteralString(*LiteralString)           {}
func (l *NullListener) ExitLiteralString(*LiteralString)            {}
func (l *NullListener) EnterLiteralBool(*LiteralBool)               {}
func (l *NullListener) ExitLiteralBool(*LiteralBool)                {}
