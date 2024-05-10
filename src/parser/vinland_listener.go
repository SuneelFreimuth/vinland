// Code generated from Vinland.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Vinland
import "github.com/antlr4-go/antlr/v4"

// VinlandListener is a complete listener for a parse tree produced by VinlandParser.
type VinlandListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterDeclList is called when entering the declList production.
	EnterDeclList(c *DeclListContext)

	// EnterDecl is called when entering the decl production.
	EnterDecl(c *DeclContext)

	// EnterStmtBlock is called when entering the stmtBlock production.
	EnterStmtBlock(c *StmtBlockContext)

	// EnterStmtList is called when entering the stmtList production.
	EnterStmtList(c *StmtListContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterBinding is called when entering the binding production.
	EnterBinding(c *BindingContext)

	// EnterFunctionDefinition is called when entering the functionDefinition production.
	EnterFunctionDefinition(c *FunctionDefinitionContext)

	// EnterParamList is called when entering the paramList production.
	EnterParamList(c *ParamListContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterNumber is called when entering the Number production.
	EnterNumber(c *NumberContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterParenthesized is called when entering the Parenthesized production.
	EnterParenthesized(c *ParenthesizedContext)

	// EnterCallExpr is called when entering the CallExpr production.
	EnterCallExpr(c *CallExprContext)

	// EnterCall is called when entering the call production.
	EnterCall(c *CallContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitDeclList is called when exiting the declList production.
	ExitDeclList(c *DeclListContext)

	// ExitDecl is called when exiting the decl production.
	ExitDecl(c *DeclContext)

	// ExitStmtBlock is called when exiting the stmtBlock production.
	ExitStmtBlock(c *StmtBlockContext)

	// ExitStmtList is called when exiting the stmtList production.
	ExitStmtList(c *StmtListContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitBinding is called when exiting the binding production.
	ExitBinding(c *BindingContext)

	// ExitFunctionDefinition is called when exiting the functionDefinition production.
	ExitFunctionDefinition(c *FunctionDefinitionContext)

	// ExitParamList is called when exiting the paramList production.
	ExitParamList(c *ParamListContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitNumber is called when exiting the Number production.
	ExitNumber(c *NumberContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitParenthesized is called when exiting the Parenthesized production.
	ExitParenthesized(c *ParenthesizedContext)

	// ExitCallExpr is called when exiting the CallExpr production.
	ExitCallExpr(c *CallExprContext)

	// ExitCall is called when exiting the call production.
	ExitCall(c *CallContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)
}
