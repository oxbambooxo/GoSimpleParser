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
		"", "'int'", "", "'='", "'+'", "'-'", "'*'", "'/'", "';'", "'('", "')'",
	}
	staticData.symbolicNames = []string{
		"", "Int", "IntegerLiteral", "AssignmentOP", "Plus", "Minus", "Star",
		"Slash", "SemiColon", "LeftParen", "RightParen", "Identifier", "Whitespace",
		"Newline",
	}
	staticData.ruleNames = []string{
		"Int", "IntegerLiteral", "AssignmentOP", "Plus", "Minus", "Star", "Slash",
		"SemiColon", "LeftParen", "RightParen", "Identifier", "Whitespace",
		"Newline",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 13, 75, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 4,
		1, 33, 8, 1, 11, 1, 12, 1, 34, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 5,
		10, 55, 8, 10, 10, 10, 12, 10, 58, 9, 10, 1, 11, 4, 11, 61, 8, 11, 11,
		11, 12, 11, 62, 1, 11, 1, 11, 1, 12, 1, 12, 3, 12, 69, 8, 12, 1, 12, 3,
		12, 72, 8, 12, 1, 12, 1, 12, 0, 0, 13, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11,
		6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 1, 0, 4, 1, 0,
		48, 57, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97,
		122, 2, 0, 9, 9, 32, 32, 79, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5,
		1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 1, 27, 1, 0, 0, 0,
		3, 32, 1, 0, 0, 0, 5, 36, 1, 0, 0, 0, 7, 38, 1, 0, 0, 0, 9, 40, 1, 0, 0,
		0, 11, 42, 1, 0, 0, 0, 13, 44, 1, 0, 0, 0, 15, 46, 1, 0, 0, 0, 17, 48,
		1, 0, 0, 0, 19, 50, 1, 0, 0, 0, 21, 52, 1, 0, 0, 0, 23, 60, 1, 0, 0, 0,
		25, 71, 1, 0, 0, 0, 27, 28, 5, 105, 0, 0, 28, 29, 5, 110, 0, 0, 29, 30,
		5, 116, 0, 0, 30, 2, 1, 0, 0, 0, 31, 33, 7, 0, 0, 0, 32, 31, 1, 0, 0, 0,
		33, 34, 1, 0, 0, 0, 34, 32, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 4, 1, 0,
		0, 0, 36, 37, 5, 61, 0, 0, 37, 6, 1, 0, 0, 0, 38, 39, 5, 43, 0, 0, 39,
		8, 1, 0, 0, 0, 40, 41, 5, 45, 0, 0, 41, 10, 1, 0, 0, 0, 42, 43, 5, 42,
		0, 0, 43, 12, 1, 0, 0, 0, 44, 45, 5, 47, 0, 0, 45, 14, 1, 0, 0, 0, 46,
		47, 5, 59, 0, 0, 47, 16, 1, 0, 0, 0, 48, 49, 5, 40, 0, 0, 49, 18, 1, 0,
		0, 0, 50, 51, 5, 41, 0, 0, 51, 20, 1, 0, 0, 0, 52, 56, 7, 1, 0, 0, 53,
		55, 7, 2, 0, 0, 54, 53, 1, 0, 0, 0, 55, 58, 1, 0, 0, 0, 56, 54, 1, 0, 0,
		0, 56, 57, 1, 0, 0, 0, 57, 22, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 59, 61,
		7, 3, 0, 0, 60, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0,
		62, 63, 1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 65, 6, 11, 0, 0, 65, 24, 1,
		0, 0, 0, 66, 68, 5, 13, 0, 0, 67, 69, 5, 10, 0, 0, 68, 67, 1, 0, 0, 0,
		68, 69, 1, 0, 0, 0, 69, 72, 1, 0, 0, 0, 70, 72, 5, 10, 0, 0, 71, 66, 1,
		0, 0, 0, 71, 70, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 74, 6, 12, 0, 0, 74,
		26, 1, 0, 0, 0, 7, 0, 34, 54, 56, 62, 68, 71, 1, 6, 0, 0,
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
	SimpleScriptLexerSemiColon      = 8
	SimpleScriptLexerLeftParen      = 9
	SimpleScriptLexerRightParen     = 10
	SimpleScriptLexerIdentifier     = 11
	SimpleScriptLexerWhitespace     = 12
	SimpleScriptLexerNewline        = 13
)