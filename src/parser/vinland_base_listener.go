// Code generated from Vinland.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Vinland
import "github.com/antlr4-go/antlr/v4"

// BaseVinlandListener is a complete listener for a parse tree produced by VinlandParser.
type BaseVinlandListener struct{}

var _ VinlandListener = &BaseVinlandListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVinlandListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVinlandListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVinlandListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVinlandListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseVinlandListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseVinlandListener) ExitProgram(ctx *ProgramContext) {}

// EnterDeclList is called when production declList is entered.
func (s *BaseVinlandListener) EnterDeclList(ctx *DeclListContext) {}

// ExitDeclList is called when production declList is exited.
func (s *BaseVinlandListener) ExitDeclList(ctx *DeclListContext) {}

// EnterDecl is called when production decl is entered.
func (s *BaseVinlandListener) EnterDecl(ctx *DeclContext) {}

// ExitDecl is called when production decl is exited.
func (s *BaseVinlandListener) ExitDecl(ctx *DeclContext) {}

// EnterStmtBlock is called when production stmtBlock is entered.
func (s *BaseVinlandListener) EnterStmtBlock(ctx *StmtBlockContext) {}

// ExitStmtBlock is called when production stmtBlock is exited.
func (s *BaseVinlandListener) ExitStmtBlock(ctx *StmtBlockContext) {}

// EnterStmtList is called when production stmtList is entered.
func (s *BaseVinlandListener) EnterStmtList(ctx *StmtListContext) {}

// ExitStmtList is called when production stmtList is exited.
func (s *BaseVinlandListener) ExitStmtList(ctx *StmtListContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseVinlandListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseVinlandListener) ExitStmt(ctx *StmtContext) {}

// EnterBinding is called when production binding is entered.
func (s *BaseVinlandListener) EnterBinding(ctx *BindingContext) {}

// ExitBinding is called when production binding is exited.
func (s *BaseVinlandListener) ExitBinding(ctx *BindingContext) {}

// EnterFunctionDefinition is called when production functionDefinition is entered.
func (s *BaseVinlandListener) EnterFunctionDefinition(ctx *FunctionDefinitionContext) {}

// ExitFunctionDefinition is called when production functionDefinition is exited.
func (s *BaseVinlandListener) ExitFunctionDefinition(ctx *FunctionDefinitionContext) {}

// EnterParamList is called when production paramList is entered.
func (s *BaseVinlandListener) EnterParamList(ctx *ParamListContext) {}

// ExitParamList is called when production paramList is exited.
func (s *BaseVinlandListener) ExitParamList(ctx *ParamListContext) {}

// EnterParam is called when production param is entered.
func (s *BaseVinlandListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseVinlandListener) ExitParam(ctx *ParamContext) {}

// EnterNumber is called when production Number is entered.
func (s *BaseVinlandListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production Number is exited.
func (s *BaseVinlandListener) ExitNumber(ctx *NumberContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseVinlandListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseVinlandListener) ExitAddSub(ctx *AddSubContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseVinlandListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseVinlandListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterParenthesized is called when production Parenthesized is entered.
func (s *BaseVinlandListener) EnterParenthesized(ctx *ParenthesizedContext) {}

// ExitParenthesized is called when production Parenthesized is exited.
func (s *BaseVinlandListener) ExitParenthesized(ctx *ParenthesizedContext) {}

// EnterCallExpr is called when production CallExpr is entered.
func (s *BaseVinlandListener) EnterCallExpr(ctx *CallExprContext) {}

// ExitCallExpr is called when production CallExpr is exited.
func (s *BaseVinlandListener) ExitCallExpr(ctx *CallExprContext) {}

// EnterCall is called when production call is entered.
func (s *BaseVinlandListener) EnterCall(ctx *CallContext) {}

// ExitCall is called when production call is exited.
func (s *BaseVinlandListener) ExitCall(ctx *CallContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseVinlandListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseVinlandListener) ExitLiteral(ctx *LiteralContext) {}
