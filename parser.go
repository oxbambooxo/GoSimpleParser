package main

import (
	"errors"
	"fmt"
)

/*

 * programm -> intDeclare | assignmentStatement | expressionStatement
 * intDeclare -> 'int' Id ( = additive) (';')
 * assignmentStatement -> Id '=' addtive (';')
 * expressionStatement -> addtive (';')
 * additive -> multiplicative ( (+ | -) multiplicative)*
 * multiplicative -> primary ( (* | /) primary)*
 * primary -> IntLiteral | Id | (additive)

 */

func ParserString(str string) (*ASTNode, error) {
	lex := &Lexer{str: str}
	tokens, err := lex.tokenize()
	if err != nil {
		return nil, err
	}
	parser := &Parser{tokens: tokens}
	return parser.parse()
}

func PrintAST(node *ASTNode) {
	fmt.Printf("===================\n")
	printAST("", node)
	fmt.Printf("===================\n")
}

func printAST(prefix string, node *ASTNode) {
	if node == nil {
		fmt.Printf("unknown ast node\n")
		return
	}
	if node.value == "" {
		fmt.Printf("%sType: %v\n", prefix, node.NodeType)
	} else {
		fmt.Printf("%sType: %v Value: %s\n", prefix, node.NodeType, node.value)
	}
	for _, child := range node.childrens {
		printAST(prefix+"\t", child)
	}
}

var (
	ErrASTEOF          = errors.New("ast end of parse")
	ErrUnknownStatment = errors.New("unknown statement")
)

type Parser struct {
	tokens []*Token
	pos    int
}

func (parser *Parser) tokenPeek() (*Token, error) {
	if parser.pos >= len(parser.tokens) {
		return nil, ErrASTEOF
	}
	return parser.tokens[parser.pos], nil
}

func (parser *Parser) tokenNext() {
	parser.pos++
}

func (parser *Parser) tokenPosition() int {
	return parser.pos
}

func (parser *Parser) tokenSeek(pos int) {
	parser.pos = pos
}

func (parser *Parser) parse() (*ASTNode, error) {
	root := NewASTNode(NodeProgram, "")
	for {
		token, err := parser.tokenPeek()
		if err != nil {
			if err == ErrASTEOF {
				break
			}
			return nil, err
		}
		child := parser.parseIntDeclare()
		if child == nil {
			child = parser.assignmentStatement()
		}
		if child == nil {
			child = parser.expressionStatement()
		}
		if child == nil {
			if token.Type == TokenSemiColon {
				parser.tokenNext()
				continue
			}
			fmt.Printf("unknown statement: %v\n", token)
			return nil, ErrUnknownStatment
		}
		root.AddChild(child)
	}
	return root, nil
}

func (parser *Parser) parseIntDeclare() (node *ASTNode) {
	token, err := parser.tokenPeek()
	if err != nil {
		return
	}
	if token.Type != TokenInt {
		return
	}
	parser.tokenNext()
	token, err = parser.tokenPeek()
	if err != nil || token.Type != TokenIdentifier {
		panic("invalid variable name")
	}
	parser.tokenNext()
	node = NewASTNode(NodeIntDeclaration, token.Value())

	token, err = parser.tokenPeek()
	if err != nil {
		return
	}
	if token.Type == TokenAssign {
		parser.tokenNext()
		child := parser.parseAdditive()
		if child == nil {
			panic("invalid variable assignment")
		}
		node.AddChild(child)
	}

	parser.parseOptionalSemiColon()
	return
}

func (parser *Parser) parseAdditive() (node *ASTNode) {
	child1 := parser.parseMultiplicative() //应用add规则
	if child1 == nil {
		return
	}
	node = child1
	for { //循环应用add'规则
		token, err := parser.tokenPeek()
		if err != nil || (token.Type != TokenPlus && token.Type != TokenMinus) {
			break
		}
		parser.tokenNext()
		child2 := parser.parseMultiplicative() //计算下级节点
		if child2 == nil {
			panic("invalid additive statment")
		}
		if token.Type == TokenPlus {
			node = NewASTNode(NodeAdditive, "")
		} else {
			node = NewASTNode(NodeSubtract, "")
		}
		node.AddChild(child1) //注意，新节点在顶层，保证正确的结合性
		node.AddChild(child2)
		child1 = node
	}
	return node
}

func (parser *Parser) parseMultiplicative() (node *ASTNode) {
	child1 := parser.parsePrimary()
	node = child1
	for {
		token, err := parser.tokenPeek()
		if err != nil || (token.Type != TokenStar && token.Type != TokenSlash) {
			break
		}
		parser.tokenNext()
		child2 := parser.parsePrimary()
		if child2 == nil {
			panic("invalid multiplicative ststment")
		}
		if token.Type == TokenStar {
			node = NewASTNode(NodeMultiply, "")
		} else {
			node = NewASTNode(NodeDivision, "")
		}
		node.AddChild(child1)
		node.AddChild(child2)
		child1 = node
	}
	return node
}

func (parser *Parser) parsePrimary() (node *ASTNode) {
	token, err := parser.tokenPeek()
	if err != nil {
		return
	}
	switch token.Type {
	case TokenIntLiteral:
		parser.tokenNext()
		node = NewASTNode(NodeIntLiteral, token.Value())
	case TokenIdentifier:
		parser.tokenNext()
		node = NewASTNode(NodeIdentifier, token.Value())
	case TokenLeftParen:
		parser.tokenNext()
		node = parser.parseAdditive()
		if node == nil {
			panic("invalid statment inside parenthesis")
		}
		token, err := parser.tokenPeek()
		if err != nil || token.Type != TokenRightParen {
			panic("miss right parenthesis")
		}
		parser.tokenNext()
	}
	return
}

func (parser *Parser) expressionStatement() (node *ASTNode) {
	// 语句没有强制的 ; 分隔符, 不需要回退, 子语句解析失败不消耗 token
	// pos := parser.tokenPosition()
	exp := parser.parseAdditive()
	if exp == nil {
		return
	}
	parser.parseOptionalSemiColon()

	node = NewASTNode(NodeExpressionStmt, "")
	node.AddChild(exp)
	return
}

func (parser *Parser) assignmentStatement() (node *ASTNode) {
	pos := parser.tokenPosition()
	token, err := parser.tokenPeek() //预读，看看下面是不是标识符
	if err != nil || token.Type != TokenIdentifier {
		return
	}
	parser.tokenNext()

	node = NewASTNode(NodeAssignmentStmt, token.Value())

	token, err = parser.tokenPeek() //预读，看看下面是不是标识符
	if err != nil || token.Type != TokenAssign {
		node = nil
		parser.tokenSeek(pos)
		return
	}
	parser.tokenNext()

	child := parser.parseAdditive()
	if child == nil { //出错，等号右面没有一个合法的表达式
		panic("invalide assignment statement")
	} else {
		node.AddChild(child)
	}

	parser.parseOptionalSemiColon()
	return
}

func (parser *Parser) parseOptionalSemiColon() {
	token, err := parser.tokenPeek()
	if err != nil {
		return
	}
	if token.Type == TokenSemiColon {
		parser.tokenNext()
	} else {
		panic("unexpect token " + token.Value())
	}
}
