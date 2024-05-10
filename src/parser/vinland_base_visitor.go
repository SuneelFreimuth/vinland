// Code generated from Vinland.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Vinland
import "github.com/antlr4-go/antlr/v4"

type BaseVinlandVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseVinlandVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVinlandVisitor) VisitDeclList(ctx *DeclListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVinlandVisitor) VisitDecl(ctx *DeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVinlandVisitor) VisitStmtBlock(ctx *StmtBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVinlandVisitor) VisitStmtList(ctx *StmtListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVinlandVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVinlandVisitor) VisitBinding(ctx *BindingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVinlandVisitor) VisitFunctionDefinition(ctx *FunctionDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVinlandVisitor) VisitParamList(ctx *ParamListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVinlandVisitor) VisitParam(ctx *ParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVinlandVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVinlandVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVinlandVisitor) VisitMulDiv(ctx *MulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVinlandVisitor) VisitParenthesized(ctx *ParenthesizedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVinlandVisitor) VisitCallExpr(ctx *CallExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVinlandVisitor) VisitCall(ctx *CallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVinlandVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}
