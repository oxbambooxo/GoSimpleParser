package main

import "fmt"

func NewToken(ttype TokenType, init byte) *Token {
	return &Token{Type: ttype, buf: []byte{init}}
}

type Token struct {
	Type TokenType
	buf  []byte
}

func (token *Token) Write(c byte) {
	token.buf = append(token.buf, c)
}

func (token *Token) Value() string {
	return string(token.buf)
}

func (token *Token) String() string {
	if token.Type == Identifier || token.Type == IntLiteral {
		return fmt.Sprintf("Token{%v: %s}", token.Type, string(token.buf))
	}
	return fmt.Sprintf("Token{%v}", token.Type)
}

type TokenType int

const (
	Plus       TokenType = iota // +
	Minus                       // -
	Star                        // *
	Slash                       // /
	SemiColon                   // ;
	LeftParen                   // (
	RightParen                  // )
	Assign                      // =
	Int                         // int
	Identifier                  // 标识符
	IntLiteral                  // 整型字面量
)

func (t TokenType) String() string {
	switch t {
	case Plus:
		return "Plus"
	case Minus:
		return "Minus"
	case Star:
		return "Star"
	case Slash:
		return "Slash"
	case SemiColon:
		return "SemiColon"
	case LeftParen:
		return "LeftParen"
	case RightParen:
		return "RightParen"
	case Assign:
		return "Assign"
	case Int:
		return "Int"
	case Identifier:
		return "Identifier"
	case IntLiteral:
		return "IntLiteral"
	default:
		return "unknown token type"
	}
}
