package main

import (
	"fmt"
	"main/parser"
	"runtime/debug"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var globalEvaluator = &Evaluator{
	local: map[string]interface{}{},
}

func EvalString(str string) (value interface{}, err error) {
	defer func() {
		rerr := recover()
		if rerr != nil {
			switch terr := rerr.(type) {
			case string:
				fmt.Println(terr)
			default:
				fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
			}
			return
		}
	}()

	// Setup the input
	is := antlr.NewInputStream(str)

	// Create the Lexer
	lexer := parser.NewSimpleScriptLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewSimpleScriptParser(stream)
	p.SetErrorHandler(antlr.NewBailErrorStrategy())

	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(globalEvaluator, p.Programm())
	var result interface{}
	if len(globalEvaluator.stack) > 0 {
		result = globalEvaluator.pop()
	}
	return result, nil
}

type Evaluator struct {
	local map[string]interface{}
	*parser.BaseSimpleScriptListener
	stack []int64
}

func (evaluator *Evaluator) push(i int64) {
	evaluator.stack = append(evaluator.stack, i)
}

func (evaluator *Evaluator) pop() int64 {
	if len(evaluator.stack) == 0 {
		panic("stack is empty unable to pop")
	}
	result := evaluator.stack[len(evaluator.stack)-1]
	evaluator.stack = evaluator.stack[:len(evaluator.stack)-1]
	return result
}

func (evaluator *Evaluator) peek() int64 {
	if len(evaluator.stack) == 0 {
		panic("stack is empty unable to peek")
	}
	return evaluator.stack[len(evaluator.stack)-1]
}

func (evaluator *Evaluator) getVariable(variableName string) int64 {
	value, exist := evaluator.local[variableName]
	if !exist {
		panic("undefined variable " + variableName)
	}
	return value.(int64)
}
