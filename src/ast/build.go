package ast

import (
	"strconv"

	"github.com/SuneelFreimuth/vinland/src/parser"
)

func Build(start parser.IProgramContext) Node {
	return start.Accept(&Builder{}).(Node)
}

type Builder struct {
	*parser.BaseVinlandVisitor
}

// func (b *Builder) VisitRoot(root *parser.ProgramContext) any {
// 	return root.Expr().Accept(b).(Node)
// }

// func (b *Builder) VisitParenthesized(ctx *parser.ParenthesizedContext) any {
// 	return ctx.Expr().Accept(b).(Node)
// }

// func (b *Builder) VisitMulDiv(ctx *parser.MulDivContext) any {
// 	left := ctx.Expr(0).Accept(b).(Node)
// 	op := OpFromString(ctx.GetOp().GetText())
// 	right := ctx.Expr(1).Accept(b).(Node)

// 	operation := NewOperation(op)
// 	operation.SetChild(0, left)
// 	operation.SetChild(1, right)

// 	return operation
// }

// func (b *Builder) VisitAddSub(ctx *parser.AddSubContext) any {
// 	left := ctx.Expr(0).Accept(b).(Node)
// 	op := OpFromString(ctx.GetOp().GetText())
// 	right := ctx.Expr(1).Accept(b).(Node)

// 	operation := NewOperation(op)
// 	operation.SetChild(0, left)
// 	operation.SetChild(1, right)

// 	return operation
// }

// func (b *Builder) VisitNumber(ctx *parser.NumberContext) any {
// 	n, err := strconv.ParseFloat(ctx.GetText(), 64)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return NewNumber(float64(n))
// }