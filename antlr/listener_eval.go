package main

import (
	"fmt"
	"main/parser"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func (evaluator *Evaluator) EnterEveryRule(ctx antlr.ParserRuleContext) {
	// fmt.Printf("%#v\n", ctx)
}

func (evaluator *Evaluator) ExitIntDeclaration(ctx *parser.IntDeclarationContext) {
	variableName := ctx.Identifier().GetText()
	var variableValue int64
	if ctx.Additive() != nil {
		variableValue = evaluator.peek()
	} else {
		evaluator.push(variableValue)
	}
	evaluator.local[variableName] = variableValue
}

func (evaluator *Evaluator) ExitAssignment(ctx *parser.AssignmentContext) {
	variableName := ctx.Identifier().GetText()
	evaluator.getVariable(variableName)
	variableValue := evaluator.peek()
	evaluator.local[variableName] = variableValue
}

func (evaluator *Evaluator) ExitAdditive(ctx *parser.AdditiveContext) {
	if ctx.GetOp() == nil {
		return
	}

	right, left := evaluator.pop(), evaluator.pop()
	switch ctx.GetOp().GetTokenType() {
	case parser.SimpleScriptParserPlus:
		evaluator.push(left + right)
	case parser.SimpleScriptParserMinus:
		evaluator.push(left - right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", ctx.GetOp().GetText()))
	}
}

func (evaluator *Evaluator) ExitMultiplicative(ctx *parser.MultiplicativeContext) {
	if ctx.GetOp() == nil {
		return
	}

	right, left := evaluator.pop(), evaluator.pop()
	switch ctx.GetOp().GetTokenType() {
	case parser.SimpleScriptParserStar:
		evaluator.push(left * right)
	case parser.SimpleScriptParserSlash:
		evaluator.push(left / right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", ctx.GetOp().GetText()))
	}
}

func (evaluator *Evaluator) ExitPrimary(ctx *parser.PrimaryContext) {

	switch {
	case ctx.IntegerLiteral() != nil:
		s := ctx.IntegerLiteral().GetText()
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(fmt.Sprintf("unexpected int64: %s", s))
		}
		evaluator.push(i)
	case ctx.Identifier() != nil:
		variableName := ctx.Identifier().GetText()
		evaluator.push(evaluator.getVariable(variableName))
	}
}
