package main

import (
	"fmt"
	"strconv"
)

var globalEvaluator = Evaluator{local: map[string]interface{}{}}

func EvalString(str string) (value interface{}, err error) {
	lex := &Lexer{str: str}
	tokens, err := lex.tokenize()
	if err != nil {
		return nil, err
	}
	parser := &Parser{tokens: tokens}
	root, err := parser.parse()
	if err != nil {
		return nil, err
	}
	defer func() {
		rerr := recover()
		if rerr != nil {
			err = fmt.Errorf("%v", rerr)
			return
		}
	}()
	return globalEvaluator.Eval(root), nil
}

type Evaluator struct {
	local map[string]interface{}
}

func (evaluator *Evaluator) Eval(node *ASTNode) (result interface{}) {
	switch node.NodeType {
	case NodeProgram:
		for _, child := range node.Childrens {
			result = evaluator.Eval(child)
		}
	case NodeIntDeclaration:
		result = evaluator.evalIntDeclaration(node)
	case NodeExpression:
		result = evaluator.evalExpression(node)
	case NodeAssignment:
		result = evaluator.evalAssignment(node)
	case NodeMultiply:
		result = evaluator.evalMultiply(node)
	case NodeDivision:
		result = evaluator.evalDivision(node)
	case NodeAdditive:
		result = evaluator.evalAdditive(node)
	case NodeSubtract:
		result = evaluator.evalSubtract(node)
	case NodeIdentifier:
		result = evaluator.evalIdentifier(node)
	case NodeIntLiteral:
		result = evaluator.evalIntLiteral(node)
	}
	return
}

func (evaluator *Evaluator) evalIntDeclaration(node *ASTNode) (result interface{}) {
	var variable int64
	evaluator.local[node.Value] = variable // 初始化变量
	if len(node.Childrens) > 0 {
		val := evaluator.Eval(node.Childrens[0])
		evaluator.local[node.Value] = val
	}
	return evaluator.local[node.Value]
}

func (evaluator *Evaluator) evalExpression(node *ASTNode) (result interface{}) {
	for _, child := range node.Childrens {
		result = evaluator.Eval(child)
	}
	return
}

func (evaluator *Evaluator) evalAssignment(node *ASTNode) (result interface{}) {
	evaluator.checkVariable(node.Value)
	val := evaluator.Eval(node.Childrens[0])
	evaluator.local[node.Value] = val
	return val
}

func (evaluator *Evaluator) evalMultiply(node *ASTNode) (result interface{}) {
	val1 := evaluator.Eval(node.Childrens[0]).(int64)
	val2 := evaluator.Eval(node.Childrens[1]).(int64)
	return val1 * val2
}

func (evaluator *Evaluator) evalDivision(node *ASTNode) (result interface{}) {
	val1 := evaluator.Eval(node.Childrens[0]).(int64)
	val2 := evaluator.Eval(node.Childrens[1]).(int64)
	return val1 / val2
}

func (evaluator *Evaluator) evalAdditive(node *ASTNode) (result interface{}) {
	val1 := evaluator.Eval(node.Childrens[0]).(int64)
	val2 := evaluator.Eval(node.Childrens[1]).(int64)
	return val1 + val2
}

func (evaluator *Evaluator) evalSubtract(node *ASTNode) (result interface{}) {
	val1 := evaluator.Eval(node.Childrens[0]).(int64)
	val2 := evaluator.Eval(node.Childrens[1]).(int64)
	return val1 - val2
}

func (evaluator *Evaluator) evalIdentifier(node *ASTNode) (result interface{}) {
	evaluator.checkVariable(node.Value)
	return evaluator.local[node.Value]
}

func (evaluator *Evaluator) evalIntLiteral(node *ASTNode) (result interface{}) {
	i, _ := strconv.ParseInt(node.Value, 10, 64)
	return i
}

func (evaluator *Evaluator) checkVariable(variableName string) {
	if _, exist := evaluator.local[variableName]; !exist {
		panic("unknown variable " + variableName)
	}
}
