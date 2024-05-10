// Code generated from Vinland.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Vinland
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by VinlandParser.
type VinlandVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by VinlandParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by VinlandParser#declList.
	VisitDeclList(ctx *DeclListContext) interface{}

	// Visit a parse tree produced by VinlandParser#decl.
	VisitDecl(ctx *DeclContext) interface{}

	// Visit a parse tree produced by VinlandParser#stmtBlock.
	VisitStmtBlock(ctx *StmtBlockContext) interface{}

	// Visit a parse tree produced by VinlandParser#stmtList.
	VisitStmtList(ctx *StmtListContext) interface{}

	// Visit a parse tree produced by VinlandParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by VinlandParser#binding.
	VisitBinding(ctx *BindingContext) interface{}

	// Visit a parse tree produced by VinlandParser#functionDefinition.
	VisitFunctionDefinition(ctx *FunctionDefinitionContext) interface{}

	// Visit a parse tree produced by VinlandParser#paramList.
	VisitParamList(ctx *ParamListContext) interface{}

	// Visit a parse tree produced by VinlandParser#param.
	VisitParam(ctx *ParamContext) interface{}

	// Visit a parse tree produced by VinlandParser#Number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by VinlandParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by VinlandParser#MulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}

	// Visit a parse tree produced by VinlandParser#Parenthesized.
	VisitParenthesized(ctx *ParenthesizedContext) interface{}

	// Visit a parse tree produced by VinlandParser#CallExpr.
	VisitCallExpr(ctx *CallExprContext) interface{}

	// Visit a parse tree produced by VinlandParser#call.
	VisitCall(ctx *CallContext) interface{}

	// Visit a parse tree produced by VinlandParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}
}
