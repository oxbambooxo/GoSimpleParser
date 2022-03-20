package main

import (
	"fmt"
	"strconv"
)

var globalEnv = Env{local: map[string]interface{}{}}

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
	return globalEnv.Eval(root), nil
}

type Env struct {
	local map[string]interface{}
}

func (env *Env) Eval(node *ASTNode) (result interface{}) {
	switch node.NodeType {
	case NodeProgram:
		for _, child := range node.Childrens {
			result = env.Eval(child)
		}
	case NodeIntDeclaration:
		result = env.evalIntDeclaration(node)
	case NodeExpression:
		result = env.evalExpression(node)
	case NodeAssignment:
		result = env.evalAssignment(node)
	case NodeMultiply:
		result = env.evalMultiply(node)
	case NodeDivision:
		result = env.evalDivision(node)
	case NodeAdditive:
		result = env.evalAdditive(node)
	case NodeSubtract:
		result = env.evalSubtract(node)
	case NodeIdentifier:
		result = env.evalIdentifier(node)
	case NodeIntLiteral:
		result = env.evalIntLiteral(node)
	}
	return
}

func (env *Env) evalIntDeclaration(node *ASTNode) (result interface{}) {
	env.local[node.Value] = 0 // 初始化变量
	if len(node.Childrens) > 0 {
		val := env.Eval(node.Childrens[0])
		env.local[node.Value] = val
	}
	return env.local[node.Value]
}

func (env *Env) evalExpression(node *ASTNode) (result interface{}) {
	for _, child := range node.Childrens {
		result = env.Eval(child)
	}
	return
}

func (env *Env) evalAssignment(node *ASTNode) (result interface{}) {
	env.checkVariable(node.Value)
	val := env.Eval(node.Childrens[0])
	env.local[node.Value] = val
	return val
}

func (env *Env) evalMultiply(node *ASTNode) (result interface{}) {
	val1 := env.Eval(node.Childrens[0]).(int64)
	val2 := env.Eval(node.Childrens[1]).(int64)
	return val1 * val2
}

func (env *Env) evalDivision(node *ASTNode) (result interface{}) {
	val1 := env.Eval(node.Childrens[0]).(int64)
	val2 := env.Eval(node.Childrens[1]).(int64)
	return val1 / val2
}

func (env *Env) evalAdditive(node *ASTNode) (result interface{}) {
	val1 := env.Eval(node.Childrens[0]).(int64)
	val2 := env.Eval(node.Childrens[1]).(int64)
	return val1 + val2
}

func (env *Env) evalSubtract(node *ASTNode) (result interface{}) {
	val1 := env.Eval(node.Childrens[0]).(int64)
	val2 := env.Eval(node.Childrens[1]).(int64)
	return val1 - val2
}

func (env *Env) evalIdentifier(node *ASTNode) (result interface{}) {
	env.checkVariable(node.Value)
	return env.local[node.Value]
}

func (env *Env) evalIntLiteral(node *ASTNode) (result interface{}) {
	i, _ := strconv.ParseInt(node.Value, 10, 64)
	return i
}

func (env *Env) checkVariable(variableName string) {
	if _, exist := env.local[variableName]; !exist {
		panic("unknown variable " + variableName)
	}
}
