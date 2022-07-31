package main

import (
	"fmt"
	"main/parser"
	"os"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func (evaluator *Evaluator) EnterEveryRule(ctx antlr.ParserRuleContext) {
	// fmt.Printf("%#v\n", ctx)
}

func (evaluator *Evaluator) ExitIntDeclaration(ctx *parser.IntDeclarationContext) {
	variableName := ctx.Identifier().GetText()
	var variableValue interface{}
	if ctx.Additive() != nil {
		variableValue = evaluator.peek()
	} else {
		evaluator.push(variableValue)
	}
	evaluator.local[variableName] = variableValue
}

func (evaluator *Evaluator) ExitAssignment(ctx *parser.AssignmentContext) {
	variableName := ctx.Identifier().GetText()
	evaluator.getObject(variableName)
	variableValue := evaluator.peek()
	evaluator.local[variableName] = variableValue
}

func (evaluator *Evaluator) ExitAdditive(ctx *parser.AdditiveContext) {
	if ctx.GetOp() == nil {
		return
	}

	rightV, leftV := evaluator.pop(), evaluator.pop()
	right, ok := rightV.(int64)
	if !ok {
		panic(fmt.Errorf("unsupport additive %v %s %v",
			rightV, ctx.GetOp().GetText(), leftV))
	}
	left, ok := leftV.(int64)
	if !ok {
		panic(fmt.Errorf("unsupport additive %v %s %v",
			rightV, ctx.GetOp().GetText(), leftV))
	}
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

	rightV, leftV := evaluator.pop(), evaluator.pop()
	right, ok := rightV.(int64)
	if !ok {
		panic(fmt.Errorf("unsupport multiplicative %v %s %v",
			rightV, ctx.GetOp().GetText(), leftV))
	}
	left, ok := leftV.(int64)
	if !ok {
		panic(fmt.Errorf("unsupport multiplicative %v %s %v",
			rightV, ctx.GetOp().GetText(), leftV))
	}
	switch ctx.GetOp().GetTokenType() {
	case parser.SimpleScriptParserStar:
		evaluator.push(left * right)
	case parser.SimpleScriptParserSlash:
		evaluator.push(left / right)
	default:
		panic(fmt.Errorf("unexpected op: %s", ctx.GetOp().GetText()))
	}
}

func (evaluator *Evaluator) ExitPrimary(ctx *parser.PrimaryContext) {

	switch {
	case ctx.IntegerLiteral() != nil:
		s := ctx.IntegerLiteral().GetText()
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(fmt.Errorf("unexpected int64: %s", s))
		}
		evaluator.push(i)
	case ctx.Identifier() != nil:
		variableName := ctx.Identifier().GetText()
		evaluator.push(evaluator.getObject(variableName))
	}
}

func (evaluator *Evaluator) ExitFunctionCall(ctx *parser.FunctionCallContext) {
	functionName := ctx.Identifier().GetText()

	var args []interface{}
	if ctx.ExpressionList() != nil {
		argsV := evaluator.pop()
		args = argsV.([]interface{})
	}

	var resultV interface{}
	switch functionName {
	case "exit":
		resultV = evaluator.callExit(args)
	case "print":
		resultV = evaluator.callPrint(args)
	default:
		panic(fmt.Errorf("undefined function %s", functionName))
	}
	evaluator.push(resultV)
}

func (evaluator *Evaluator) callExit(args []interface{}) interface{} {
	if len(args) > 1 {
		panic(fmt.Errorf("exit unsupport more than one argument"))
	}
	var exitStatusCode int
	if len(args) == 1 {
		num, ok := args[0].(int64)
		if !ok {
			panic(fmt.Errorf("exit bad argument: %v", args[0]))
		}
		exitStatusCode = int(num)
	}
	os.Exit(exitStatusCode)
	return 0
}

func (evaluator *Evaluator) callPrint(args []interface{}) interface{} {
	fmt.Println(args)
	return 0
}

func (evaluator *Evaluator) ExitExpressionList(ctx *parser.ExpressionListContext) {
	var results []interface{}

	for range ctx.AllExpression() {
		resultV := evaluator.pop()
		results = append(results, resultV)
	}

	// reverse slice
	for i, j := 0, len(results)-1; i < j; i, j = i+1, j-1 {
		results[i], results[j] = results[j], results[i]
	}
	evaluator.push(results)
}
