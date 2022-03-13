package main

import (
	"errors"
)

func TokenizeString(str string) ([]*Token, error) {
	lex := &Lexer{str: str}
	return lex.tokenize()
}

var (
	ErrEOF = errors.New("lex end of source")
)

type Lexer struct {
	str string
	pos int
}

func (lex *Lexer) peek() (byte, error) {
	if lex.pos >= len(lex.str) {
		return 0, ErrEOF
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
			if err == ErrEOF {
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
		return NewToken(Assign, c), nil
	case c == '+':
		return NewToken(Plus, c), nil
	case c == '-':
		return NewToken(Minus, c), nil
	case c == '*':
		return NewToken(Star, c), nil
	case c == '/':
		return NewToken(Slash, c), nil
	case c == ';':
		return NewToken(SemiColon, c), nil
	case c == '(':
		return NewToken(LeftParen, c), nil
	case c == ')':
		return NewToken(RightParen, c), nil
	case isDigit(c):
		return lex.lexDigital(c), nil
	default:
		return nil, errors.New("unsupported: " + string(c))
	}
}

var keywords = map[string]TokenType{
	"int": Int,
}

func (lex *Lexer) lexIdentifier(c byte) *Token {
	token := NewToken(Identifier, c)
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
	token := NewToken(IntLiteral, c)
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
