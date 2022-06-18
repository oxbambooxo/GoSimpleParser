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
	if token.Type == TokenIdentifier || token.Type == TokenIntLiteral {
		return fmt.Sprintf("Token{%v: %s}", token.Type, string(token.buf))
	}
	return fmt.Sprintf("Token{%v}", token.Type)
}

type TokenType int

const (
	TokenPlus       TokenType = iota // +
	TokenMinus                       // -
	TokenStar                        // *
	TokenSlash                       // /
	TokenSemiColon                   // ;
	TokenLeftParen                   // (
	TokenRightParen                  // )
	TokenAssign                      // =
	TokenInt                         // int
	TokenIdentifier                  // 标识符
	TokenIntLiteral                  // 整型字面量
)

func (t TokenType) String() string {
	switch t {
	case TokenPlus:
		return "Plus"
	case TokenMinus:
		return "Minus"
	case TokenStar:
		return "Star"
	case TokenSlash:
		return "Slash"
	case TokenSemiColon:
		return "SemiColon"
	case TokenLeftParen:
		return "LeftParen"
	case TokenRightParen:
		return "RightParen"
	case TokenAssign:
		return "Assign"
	case TokenInt:
		return "Int"
	case TokenIdentifier:
		return "Identifier"
	case TokenIntLiteral:
		return "IntLiteral"
	default:
		return "unknown token type"
	}
}
