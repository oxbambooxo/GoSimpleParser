// Code generated from SimpleScript.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SimpleScript

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SimpleScriptParser struct {
	*antlr.BaseParser
}

var simplescriptParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func simplescriptParserInit() {
	staticData := &simplescriptParserStaticData
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
		"programm", "statement", "intDeclaration", "assignment", "expression",
		"additive", "multiplicative", "primary", "functionCall", "expressionList",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 14, 97, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 0, 1,
		0, 1, 0, 1, 0, 4, 0, 25, 8, 0, 11, 0, 12, 0, 26, 3, 0, 29, 8, 0, 1, 1,
		1, 1, 1, 1, 3, 1, 34, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 40, 8, 2, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 3, 4, 48, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 5, 5, 56, 8, 5, 10, 5, 12, 5, 59, 9, 5, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 5, 6, 67, 8, 6, 10, 6, 12, 6, 70, 9, 6, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 3, 7, 78, 8, 7, 1, 8, 1, 8, 1, 8, 3, 8, 83, 8, 8, 1,
		8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 4, 9, 91, 8, 9, 11, 9, 12, 9, 92, 3, 9,
		95, 8, 9, 1, 9, 0, 2, 10, 12, 10, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 0,
		2, 1, 0, 4, 5, 1, 0, 6, 7, 99, 0, 28, 1, 0, 0, 0, 2, 33, 1, 0, 0, 0, 4,
		35, 1, 0, 0, 0, 6, 41, 1, 0, 0, 0, 8, 47, 1, 0, 0, 0, 10, 49, 1, 0, 0,
		0, 12, 60, 1, 0, 0, 0, 14, 77, 1, 0, 0, 0, 16, 79, 1, 0, 0, 0, 18, 94,
		1, 0, 0, 0, 20, 29, 3, 2, 1, 0, 21, 22, 3, 2, 1, 0, 22, 23, 5, 9, 0, 0,
		23, 25, 1, 0, 0, 0, 24, 21, 1, 0, 0, 0, 25, 26, 1, 0, 0, 0, 26, 24, 1,
		0, 0, 0, 26, 27, 1, 0, 0, 0, 27, 29, 1, 0, 0, 0, 28, 20, 1, 0, 0, 0, 28,
		24, 1, 0, 0, 0, 29, 1, 1, 0, 0, 0, 30, 34, 3, 4, 2, 0, 31, 34, 3, 6, 3,
		0, 32, 34, 3, 8, 4, 0, 33, 30, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 33, 32,
		1, 0, 0, 0, 34, 3, 1, 0, 0, 0, 35, 36, 5, 1, 0, 0, 36, 39, 5, 12, 0, 0,
		37, 38, 5, 3, 0, 0, 38, 40, 3, 10, 5, 0, 39, 37, 1, 0, 0, 0, 39, 40, 1,
		0, 0, 0, 40, 5, 1, 0, 0, 0, 41, 42, 5, 12, 0, 0, 42, 43, 5, 3, 0, 0, 43,
		44, 3, 10, 5, 0, 44, 7, 1, 0, 0, 0, 45, 48, 3, 10, 5, 0, 46, 48, 3, 16,
		8, 0, 47, 45, 1, 0, 0, 0, 47, 46, 1, 0, 0, 0, 48, 9, 1, 0, 0, 0, 49, 50,
		6, 5, -1, 0, 50, 51, 3, 12, 6, 0, 51, 57, 1, 0, 0, 0, 52, 53, 10, 1, 0,
		0, 53, 54, 7, 0, 0, 0, 54, 56, 3, 12, 6, 0, 55, 52, 1, 0, 0, 0, 56, 59,
		1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 11, 1, 0, 0, 0,
		59, 57, 1, 0, 0, 0, 60, 61, 6, 6, -1, 0, 61, 62, 3, 14, 7, 0, 62, 68, 1,
		0, 0, 0, 63, 64, 10, 1, 0, 0, 64, 65, 7, 1, 0, 0, 65, 67, 3, 14, 7, 0,
		66, 63, 1, 0, 0, 0, 67, 70, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 68, 69, 1,
		0, 0, 0, 69, 13, 1, 0, 0, 0, 70, 68, 1, 0, 0, 0, 71, 78, 5, 2, 0, 0, 72,
		78, 5, 12, 0, 0, 73, 74, 5, 10, 0, 0, 74, 75, 3, 8, 4, 0, 75, 76, 5, 11,
		0, 0, 76, 78, 1, 0, 0, 0, 77, 71, 1, 0, 0, 0, 77, 72, 1, 0, 0, 0, 77, 73,
		1, 0, 0, 0, 78, 15, 1, 0, 0, 0, 79, 80, 5, 12, 0, 0, 80, 82, 5, 10, 0,
		0, 81, 83, 3, 18, 9, 0, 82, 81, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 84,
		1, 0, 0, 0, 84, 85, 5, 11, 0, 0, 85, 17, 1, 0, 0, 0, 86, 95, 3, 8, 4, 0,
		87, 90, 3, 8, 4, 0, 88, 89, 5, 8, 0, 0, 89, 91, 3, 8, 4, 0, 90, 88, 1,
		0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93,
		95, 1, 0, 0, 0, 94, 86, 1, 0, 0, 0, 94, 87, 1, 0, 0, 0, 95, 19, 1, 0, 0,
		0, 11, 26, 28, 33, 39, 47, 57, 68, 77, 82, 92, 94,
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

// SimpleScriptParserInit initializes any static state used to implement SimpleScriptParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSimpleScriptParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SimpleScriptParserInit() {
	staticData := &simplescriptParserStaticData
	staticData.once.Do(simplescriptParserInit)
}

// NewSimpleScriptParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSimpleScriptParser(input antlr.TokenStream) *SimpleScriptParser {
	SimpleScriptParserInit()
	this := new(SimpleScriptParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &simplescriptParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "SimpleScript.g4"

	return this
}

// SimpleScriptParser tokens.
const (
	SimpleScriptParserEOF            = antlr.TokenEOF
	SimpleScriptParserInt            = 1
	SimpleScriptParserIntegerLiteral = 2
	SimpleScriptParserAssignmentOP   = 3
	SimpleScriptParserPlus           = 4
	SimpleScriptParserMinus          = 5
	SimpleScriptParserStar           = 6
	SimpleScriptParserSlash          = 7
	SimpleScriptParserComma          = 8
	SimpleScriptParserSemiColon      = 9
	SimpleScriptParserLeftParen      = 10
	SimpleScriptParserRightParen     = 11
	SimpleScriptParserIdentifier     = 12
	SimpleScriptParserWhitespace     = 13
	SimpleScriptParserNewline        = 14
)

// SimpleScriptParser rules.
const (
	SimpleScriptParserRULE_programm       = 0
	SimpleScriptParserRULE_statement      = 1
	SimpleScriptParserRULE_intDeclaration = 2
	SimpleScriptParserRULE_assignment     = 3
	SimpleScriptParserRULE_expression     = 4
	SimpleScriptParserRULE_additive       = 5
	SimpleScriptParserRULE_multiplicative = 6
	SimpleScriptParserRULE_primary        = 7
	SimpleScriptParserRULE_functionCall   = 8
	SimpleScriptParserRULE_expressionList = 9
)

// IProgrammContext is an interface to support dynamic dispatch.
type IProgrammContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgrammContext differentiates from other interfaces.
	IsProgrammContext()
}

type ProgrammContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgrammContext() *ProgrammContext {
	var p = new(ProgrammContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleScriptParserRULE_programm
	return p
}

func (*ProgrammContext) IsProgrammContext() {}

func NewProgrammContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgrammContext {
	var p = new(ProgrammContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleScriptParserRULE_programm

	return p
}

func (s *ProgrammContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgrammContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgrammContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgrammContext) AllSemiColon() []antlr.TerminalNode {
	return s.GetTokens(SimpleScriptParserSemiColon)
}

func (s *ProgrammContext) SemiColon(i int) antlr.TerminalNode {
	return s.GetToken(SimpleScriptParserSemiColon, i)
}

func (s *ProgrammContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgrammContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgrammContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleScriptListener); ok {
		listenerT.EnterProgramm(s)
	}
}

func (s *ProgrammContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleScriptListener); ok {
		listenerT.ExitProgramm(s)
	}
}

func (p *SimpleScriptParser) Programm() (localctx IProgrammContext) {
	this := p
	_ = this

	localctx = NewProgrammContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SimpleScriptParserRULE_programm)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(20)
			p.Statement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(24)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SimpleScriptParserInt)|(1<<SimpleScriptParserIntegerLiteral)|(1<<SimpleScriptParserLeftParen)|(1<<SimpleScriptParserIdentifier))) != 0) {
			{
				p.SetState(21)
				p.Statement()
			}
			{
				p.SetState(22)
				p.Match(SimpleScriptParserSemiColon)
			}

			p.SetState(26)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleScriptParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleScriptParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) IntDeclaration() IIntDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntDeclarationContext)
}

func (s *StatementContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleScriptListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleScriptListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *SimpleScriptParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SimpleScriptParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(30)
			p.IntDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(31)
			p.Assignment()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(32)
			p.Expression()
		}

	}

	return localctx
}

// IIntDeclarationContext is an interface to support dynamic dispatch.
type IIntDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntDeclarationContext differentiates from other interfaces.
	IsIntDeclarationContext()
}

type IntDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntDeclarationContext() *IntDeclarationContext {
	var p = new(IntDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleScriptParserRULE_intDeclaration
	return p
}

func (*IntDeclarationContext) IsIntDeclarationContext() {}

func NewIntDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntDeclarationContext {
	var p = new(IntDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleScriptParserRULE_intDeclaration

	return p
}

func (s *IntDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *IntDeclarationContext) Int() antlr.TerminalNode {
	return s.GetToken(SimpleScriptParserInt, 0)
}

func (s *IntDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(SimpleScriptParserIdentifier, 0)
}

func (s *IntDeclarationContext) AssignmentOP() antlr.TerminalNode {
	return s.GetToken(SimpleScriptParserAssignmentOP, 0)
}

func (s *IntDeclarationContext) Additive() IAdditiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditiveContext)
}

func (s *IntDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleScriptListener); ok {
		listenerT.EnterIntDeclaration(s)
	}
}

func (s *IntDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleScriptListener); ok {
		listenerT.ExitIntDeclaration(s)
	}
}

func (p *SimpleScriptParser) IntDeclaration() (localctx IIntDeclarationContext) {
	this := p
	_ = this

	localctx = NewIntDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SimpleScriptParserRULE_intDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(35)
		p.Match(SimpleScriptParserInt)
	}
	{
		p.SetState(36)
		p.Match(SimpleScriptParserIdentifier)
	}
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleScriptParserAssignmentOP {
		{
			p.SetState(37)
			p.Match(SimpleScriptParserAssignmentOP)
		}
		{
			p.SetState(38)
			p.additive(0)
		}

	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleScriptParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleScriptParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Identifier() antlr.TerminalNode {
	return s.GetToken(SimpleScriptParserIdentifier, 0)
}

func (s *AssignmentContext) AssignmentOP() antlr.TerminalNode {
	return s.GetToken(SimpleScriptParserAssignmentOP, 0)
}

func (s *AssignmentContext) Additive() IAdditiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditiveContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleScriptListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleScriptListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *SimpleScriptParser) Assignment() (localctx IAssignmentContext) {
	this := p
	_ = this

	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SimpleScriptParserRULE_assignment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(41)
		p.Match(SimpleScriptParserIdentifier)
	}
	{
		p.SetState(42)
		p.Match(SimpleScriptParserAssignmentOP)
	}
	{
		p.SetState(43)
		p.additive(0)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleScriptParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleScriptParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Additive() IAdditiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditiveContext)
}

func (s *ExpressionContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleScriptListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleScriptListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SimpleScriptParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SimpleScriptParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(45)
			p.additive(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(46)
			p.FunctionCall()
		}

	}

	return localctx
}

// IAdditiveContext is an interface to support dynamic dispatch.
type IAdditiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsAdditiveContext differentiates from other interfaces.
	IsAdditiveContext()
}

type AdditiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyAdditiveContext() *AdditiveContext {
	var p = new(AdditiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleScriptParserRULE_additive
	return p
}

func (*AdditiveContext) IsAdditiveContext() {}

func NewAdditiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveContext {
	var p = new(AdditiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleScriptParserRULE_additive

	return p
}

func (s *AdditiveContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveContext) GetOp() antlr.Token { return s.op }

func (s *AdditiveContext) SetOp(v antlr.Token) { s.op = v }

func (s *AdditiveContext) Multiplicative() IMultiplicativeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicativeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeContext)
}

func (s *AdditiveContext) Additive() IAdditiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditiveContext)
}

func (s *AdditiveContext) Plus() antlr.TerminalNode {
	return s.GetToken(SimpleScriptParserPlus, 0)
}

func (s *AdditiveContext) Minus() antlr.TerminalNode {
	return s.GetToken(SimpleScriptParserMinus, 0)
}

func (s *AdditiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdditiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleScriptListener); ok {
		listenerT.EnterAdditive(s)
	}
}

func (s *AdditiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleScriptListener); ok {
		listenerT.ExitAdditive(s)
	}
}

func (p *SimpleScriptParser) Additive() (localctx IAdditiveContext) {
	return p.additive(0)
}

func (p *SimpleScriptParser) additive(_p int) (localctx IAdditiveContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewAdditiveContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAdditiveContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, SimpleScriptParserRULE_additive, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.multiplicative(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAdditiveContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, SimpleScriptParserRULE_additive)
			p.SetState(52)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(53)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*AdditiveContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == SimpleScriptParserPlus || _la == SimpleScriptParserMinus) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*AdditiveContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(54)
				p.multiplicative(0)
			}

		}
		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IMultiplicativeContext is an interface to support dynamic dispatch.
type IMultiplicativeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsMultiplicativeContext differentiates from other interfaces.
	IsMultiplicativeContext()
}

type MultiplicativeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyMultiplicativeContext() *MultiplicativeContext {
	var p = new(MultiplicativeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleScriptParserRULE_multiplicative
	return p
}

func (*MultiplicativeContext) IsMultiplicativeContext() {}

func NewMultiplicativeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeContext {
	var p = new(MultiplicativeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleScriptParserRULE_multiplicative

	return p
}

func (s *MultiplicativeContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeContext) GetOp() antlr.Token { return s.op }

func (s *MultiplicativeContext) SetOp(v antlr.Token) { s.op = v }

func (s *MultiplicativeContext) Primary() IPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *MultiplicativeContext) Multiplicative() IMultiplicativeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicativeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeContext)
}

func (s *MultiplicativeContext) Star() antlr.TerminalNode {
	return s.GetToken(SimpleScriptParserStar, 0)
}

func (s *MultiplicativeContext) Slash() antlr.TerminalNode {
	return s.GetToken(SimpleScriptParserSlash, 0)
}

func (s *MultiplicativeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplicativeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleScriptListener); ok {
		listenerT.EnterMultiplicative(s)
	}
}

func (s *MultiplicativeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleScriptListener); ok {
		listenerT.ExitMultiplicative(s)
	}
}

func (p *SimpleScriptParser) Multiplicative() (localctx IMultiplicativeContext) {
	return p.multiplicative(0)
}

func (p *SimpleScriptParser) multiplicative(_p int) (localctx IMultiplicativeContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewMultiplicativeContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMultiplicativeContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, SimpleScriptParserRULE_multiplicative, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		p.Primary()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewMultiplicativeContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, SimpleScriptParserRULE_multiplicative)
			p.SetState(63)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(64)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*MultiplicativeContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == SimpleScriptParserStar || _la == SimpleScriptParserSlash) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*MultiplicativeContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(65)
				p.Primary()
			}

		}
		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleScriptParserRULE_primary
	return p
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleScriptParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(SimpleScriptParserIntegerLiteral, 0)
}

func (s *PrimaryContext) Identifier() antlr.TerminalNode {
	return s.GetToken(SimpleScriptParserIdentifier, 0)
}

func (s *PrimaryContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(SimpleScriptParserLeftParen, 0)
}

func (s *PrimaryContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrimaryContext) RightParen() antlr.TerminalNode {
	return s.GetToken(SimpleScriptParserRightParen, 0)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleScriptListener); ok {
		listenerT.EnterPrimary(s)
	}
}

func (s *PrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleScriptListener); ok {
		listenerT.ExitPrimary(s)
	}
}

func (p *SimpleScriptParser) Primary() (localctx IPrimaryContext) {
	this := p
	_ = this

	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SimpleScriptParserRULE_primary)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(77)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SimpleScriptParserIntegerLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(71)
			p.Match(SimpleScriptParserIntegerLiteral)
		}

	case SimpleScriptParserIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(72)
			p.Match(SimpleScriptParserIdentifier)
		}

	case SimpleScriptParserLeftParen:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(73)
			p.Match(SimpleScriptParserLeftParen)
		}
		{
			p.SetState(74)
			p.Expression()
		}
		{
			p.SetState(75)
			p.Match(SimpleScriptParserRightParen)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleScriptParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleScriptParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) Identifier() antlr.TerminalNode {
	return s.GetToken(SimpleScriptParserIdentifier, 0)
}

func (s *FunctionCallContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(SimpleScriptParserLeftParen, 0)
}

func (s *FunctionCallContext) RightParen() antlr.TerminalNode {
	return s.GetToken(SimpleScriptParserRightParen, 0)
}

func (s *FunctionCallContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleScriptListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleScriptListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (p *SimpleScriptParser) FunctionCall() (localctx IFunctionCallContext) {
	this := p
	_ = this

	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SimpleScriptParserRULE_functionCall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Match(SimpleScriptParserIdentifier)
	}
	{
		p.SetState(80)
		p.Match(SimpleScriptParserLeftParen)
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SimpleScriptParserIntegerLiteral)|(1<<SimpleScriptParserLeftParen)|(1<<SimpleScriptParserIdentifier))) != 0 {
		{
			p.SetState(81)
			p.ExpressionList()
		}

	}
	{
		p.SetState(84)
		p.Match(SimpleScriptParserRightParen)
	}

	return localctx
}

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleScriptParserRULE_expressionList
	return p
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleScriptParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionListContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(SimpleScriptParserComma)
}

func (s *ExpressionListContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(SimpleScriptParserComma, i)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleScriptListener); ok {
		listenerT.EnterExpressionList(s)
	}
}

func (s *ExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleScriptListener); ok {
		listenerT.ExitExpressionList(s)
	}
}

func (p *SimpleScriptParser) ExpressionList() (localctx IExpressionListContext) {
	this := p
	_ = this

	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SimpleScriptParserRULE_expressionList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(86)
			p.Expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(87)
			p.Expression()
		}
		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == SimpleScriptParserComma {
			{
				p.SetState(88)
				p.Match(SimpleScriptParserComma)
			}
			{
				p.SetState(89)
				p.Expression()
			}

			p.SetState(92)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

func (p *SimpleScriptParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *AdditiveContext = nil
		if localctx != nil {
			t = localctx.(*AdditiveContext)
		}
		return p.Additive_Sempred(t, predIndex)

	case 6:
		var t *MultiplicativeContext = nil
		if localctx != nil {
			t = localctx.(*MultiplicativeContext)
		}
		return p.Multiplicative_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SimpleScriptParser) Additive_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SimpleScriptParser) Multiplicative_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
