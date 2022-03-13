package main

import (
	"errors"
)

func TokenizeString(str string) ([]*Token, error) {
	lex := &Lexer{str: str}
	return lex.tokenize()
}

var (
	ErrLexEOF = errors.New("lex end of source")
)

type Lexer struct {
	str string
	pos int
}

func (lex *Lexer) peek() (byte, error) {
	if lex.pos >= len(lex.str) {
		return 0, ErrLexEOF
	}
	return lex.str[lex.pos], nil
}

func (lex *Lexer) next() {
	lex.pos++
}

func (lex *Lexer) tokenize() ([]*Token, error) {
	var tokens []*Token
	for {
		lex.lexWhitespace()

		token, err := lex.lexToken()
		if err != nil {
			if err == ErrLexEOF {
				break
			}
			return nil, err
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}

func (lex *Lexer) lexWhitespace() {
	for {
		c, err := lex.peek()
		if err != nil {
			return
		}
		if isWhitespace(c) {
			lex.next()
			continue
		}
		break
	}
}

func (lex *Lexer) lexToken() (*Token, error) {
	c, err := lex.peek()
	if err != nil {
		return nil, err
	}
	lex.next()
	switch {
	case isAlpha(c):
		return lex.lexIdentifier(c), nil
	case c == '=':
		return NewToken(TokenAssign, c), nil
	case c == '+':
		return NewToken(TokenPlus, c), nil
	case c == '-':
		return NewToken(TokenMinus, c), nil
	case c == '*':
		return NewToken(TokenStar, c), nil
	case c == '/':
		return NewToken(TokenSlash, c), nil
	case c == ';':
		return NewToken(TokenSemiColon, c), nil
	case c == '(':
		return NewToken(TokenLeftParen, c), nil
	case c == ')':
		return NewToken(TokenRightParen, c), nil
	case isDigit(c):
		return lex.lexDigital(c), nil
	default:
		return nil, errors.New("unsupported: " + string(c))
	}
}

var keywords = map[string]TokenType{
	"int": TokenInt,
}

func (lex *Lexer) lexIdentifier(c byte) *Token {
	token := NewToken(TokenIdentifier, c)
	for {
		c, err := lex.peek()
		if err != nil {
			return token
		}
		if isAlpha(c) || isDigit(c) {
			lex.next()
			token.Write(c)
			continue
		}
		if tokenType, ok := keywords[token.Value()]; ok {
			token.Type = tokenType
		}
		return token
	}
}

func (lex *Lexer) lexDigital(c byte) *Token {
	token := NewToken(TokenIntLiteral, c)
	for {
		c, err := lex.peek()
		if err != nil {
			return token
		}
		if isDigit(c) {
			lex.next()
			token.Write(c)
			continue
		}
		return token
	}
}
