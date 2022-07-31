// Code generated from SimpleScript.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type SimpleScriptLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var simplescriptlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func simplescriptlexerLexerInit() {
	staticData := &simplescriptlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'int'", "", "'='", "'+'", "'-'", "'*'", "'/'", "','", "';'", "'('",
		"')'",
	}
	staticData.symbolicNames = []string{
		"", "Int", "IntegerLiteral", "AssignmentOP", "Plus", "Minus", "Star",
		"Slash", "Comma", "SemiColon", "LeftParen", "RightParen", "Identifier",
		"Whitespace", "Newline",
	}
	staticData.ruleNames = []string{
		"Int", "IntegerLiteral", "AssignmentOP", "Plus", "Minus", "Star", "Slash",
		"Comma", "SemiColon", "LeftParen", "RightParen", "Identifier", "Whitespace",
		"Newline",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 14, 79, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 1, 4, 1, 35, 8, 1, 11, 1, 12, 1, 36, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1,
		10, 1, 10, 1, 11, 1, 11, 5, 11, 59, 8, 11, 10, 11, 12, 11, 62, 9, 11, 1,
		12, 4, 12, 65, 8, 12, 11, 12, 12, 12, 66, 1, 12, 1, 12, 1, 13, 1, 13, 3,
		13, 73, 8, 13, 1, 13, 3, 13, 76, 8, 13, 1, 13, 1, 13, 0, 0, 14, 1, 1, 3,
		2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12,
		25, 13, 27, 14, 1, 0, 4, 1, 0, 48, 57, 3, 0, 65, 90, 95, 95, 97, 122, 4,
		0, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 9, 9, 32, 32, 83, 0, 1, 1, 0,
		0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0,
		0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1,
		0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25,
		1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 1, 29, 1, 0, 0, 0, 3, 34, 1, 0, 0, 0, 5,
		38, 1, 0, 0, 0, 7, 40, 1, 0, 0, 0, 9, 42, 1, 0, 0, 0, 11, 44, 1, 0, 0,
		0, 13, 46, 1, 0, 0, 0, 15, 48, 1, 0, 0, 0, 17, 50, 1, 0, 0, 0, 19, 52,
		1, 0, 0, 0, 21, 54, 1, 0, 0, 0, 23, 56, 1, 0, 0, 0, 25, 64, 1, 0, 0, 0,
		27, 75, 1, 0, 0, 0, 29, 30, 5, 105, 0, 0, 30, 31, 5, 110, 0, 0, 31, 32,
		5, 116, 0, 0, 32, 2, 1, 0, 0, 0, 33, 35, 7, 0, 0, 0, 34, 33, 1, 0, 0, 0,
		35, 36, 1, 0, 0, 0, 36, 34, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 4, 1, 0,
		0, 0, 38, 39, 5, 61, 0, 0, 39, 6, 1, 0, 0, 0, 40, 41, 5, 43, 0, 0, 41,
		8, 1, 0, 0, 0, 42, 43, 5, 45, 0, 0, 43, 10, 1, 0, 0, 0, 44, 45, 5, 42,
		0, 0, 45, 12, 1, 0, 0, 0, 46, 47, 5, 47, 0, 0, 47, 14, 1, 0, 0, 0, 48,
		49, 5, 44, 0, 0, 49, 16, 1, 0, 0, 0, 50, 51, 5, 59, 0, 0, 51, 18, 1, 0,
		0, 0, 52, 53, 5, 40, 0, 0, 53, 20, 1, 0, 0, 0, 54, 55, 5, 41, 0, 0, 55,
		22, 1, 0, 0, 0, 56, 60, 7, 1, 0, 0, 57, 59, 7, 2, 0, 0, 58, 57, 1, 0, 0,
		0, 59, 62, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 24,
		1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 63, 65, 7, 3, 0, 0, 64, 63, 1, 0, 0, 0,
		65, 66, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 68, 1,
		0, 0, 0, 68, 69, 6, 12, 0, 0, 69, 26, 1, 0, 0, 0, 70, 72, 5, 13, 0, 0,
		71, 73, 5, 10, 0, 0, 72, 71, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 76, 1,
		0, 0, 0, 74, 76, 5, 10, 0, 0, 75, 70, 1, 0, 0, 0, 75, 74, 1, 0, 0, 0, 76,
		77, 1, 0, 0, 0, 77, 78, 6, 13, 0, 0, 78, 28, 1, 0, 0, 0, 7, 0, 36, 58,
		60, 66, 72, 75, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// SimpleScriptLexerInit initializes any static state used to implement SimpleScriptLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSimpleScriptLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SimpleScriptLexerInit() {
	staticData := &simplescriptlexerLexerStaticData
	staticData.once.Do(simplescriptlexerLexerInit)
}

// NewSimpleScriptLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSimpleScriptLexer(input antlr.CharStream) *SimpleScriptLexer {
	SimpleScriptLexerInit()
	l := new(SimpleScriptLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &simplescriptlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "SimpleScript.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SimpleScriptLexer tokens.
const (
	SimpleScriptLexerInt            = 1
	SimpleScriptLexerIntegerLiteral = 2
	SimpleScriptLexerAssignmentOP   = 3
	SimpleScriptLexerPlus           = 4
	SimpleScriptLexerMinus          = 5
	SimpleScriptLexerStar           = 6
	SimpleScriptLexerSlash          = 7
	SimpleScriptLexerComma          = 8
	SimpleScriptLexerSemiColon      = 9
	SimpleScriptLexerLeftParen      = 10
	SimpleScriptLexerRightParen     = 11
	SimpleScriptLexerIdentifier     = 12
	SimpleScriptLexerWhitespace     = 13
	SimpleScriptLexerNewline        = 14
)
