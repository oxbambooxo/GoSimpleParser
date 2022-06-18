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
		"", "'int'", "", "'='", "'+'", "'-'", "'*'", "'/'", "';'", "'('", "')'",
	}
	staticData.symbolicNames = []string{
		"", "Int", "IntegerLiteral", "AssignmentOP", "Plus", "Minus", "Star",
		"Slash", "SemiColon", "LeftParen", "RightParen", "Identifier", "Whitespace",
		"Newline",
	}
	staticData.ruleNames = []string{
		"programm", "intDeclaration", "assignment", "expression", "additive",
		"multiplicative", "primary",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 13, 66, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 1, 0, 5, 0, 18, 8, 0, 10, 0, 12,
		0, 21, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 27, 8, 1, 1, 1, 3, 1, 30, 8,
		1, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 36, 8, 2, 1, 3, 1, 3, 3, 3, 40, 8, 3,
		1, 4, 1, 4, 1, 4, 5, 4, 45, 8, 4, 10, 4, 12, 4, 48, 9, 4, 1, 5, 1, 5, 1,
		5, 5, 5, 53, 8, 5, 10, 5, 12, 5, 56, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 3, 6, 64, 8, 6, 1, 6, 0, 0, 7, 0, 2, 4, 6, 8, 10, 12, 0, 2, 1, 0,
		4, 5, 1, 0, 6, 7, 69, 0, 19, 1, 0, 0, 0, 2, 22, 1, 0, 0, 0, 4, 31, 1, 0,
		0, 0, 6, 37, 1, 0, 0, 0, 8, 41, 1, 0, 0, 0, 10, 49, 1, 0, 0, 0, 12, 63,
		1, 0, 0, 0, 14, 18, 3, 2, 1, 0, 15, 18, 3, 4, 2, 0, 16, 18, 3, 6, 3, 0,
		17, 14, 1, 0, 0, 0, 17, 15, 1, 0, 0, 0, 17, 16, 1, 0, 0, 0, 18, 21, 1,
		0, 0, 0, 19, 17, 1, 0, 0, 0, 19, 20, 1, 0, 0, 0, 20, 1, 1, 0, 0, 0, 21,
		19, 1, 0, 0, 0, 22, 23, 5, 1, 0, 0, 23, 26, 5, 11, 0, 0, 24, 25, 5, 3,
		0, 0, 25, 27, 3, 8, 4, 0, 26, 24, 1, 0, 0, 0, 26, 27, 1, 0, 0, 0, 27, 29,
		1, 0, 0, 0, 28, 30, 5, 8, 0, 0, 29, 28, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0,
		30, 3, 1, 0, 0, 0, 31, 32, 5, 11, 0, 0, 32, 33, 5, 3, 0, 0, 33, 35, 3,
		8, 4, 0, 34, 36, 5, 8, 0, 0, 35, 34, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36,
		5, 1, 0, 0, 0, 37, 39, 3, 8, 4, 0, 38, 40, 5, 8, 0, 0, 39, 38, 1, 0, 0,
		0, 39, 40, 1, 0, 0, 0, 40, 7, 1, 0, 0, 0, 41, 46, 3, 10, 5, 0, 42, 43,
		7, 0, 0, 0, 43, 45, 3, 10, 5, 0, 44, 42, 1, 0, 0, 0, 45, 48, 1, 0, 0, 0,
		46, 44, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 9, 1, 0, 0, 0, 48, 46, 1, 0,
		0, 0, 49, 54, 3, 12, 6, 0, 50, 51, 7, 1, 0, 0, 51, 53, 3, 12, 6, 0, 52,
		50, 1, 0, 0, 0, 53, 56, 1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 54, 55, 1, 0, 0,
		0, 55, 11, 1, 0, 0, 0, 56, 54, 1, 0, 0, 0, 57, 64, 5, 2, 0, 0, 58, 64,
		5, 11, 0, 0, 59, 60, 5, 9, 0, 0, 60, 61, 3, 6, 3, 0, 61, 62, 5, 10, 0,
		0, 62, 64, 1, 0, 0, 0, 63, 57, 1, 0, 0, 0, 63, 58, 1, 0, 0, 0, 63, 59,
		1, 0, 0, 0, 64, 13, 1, 0, 0, 0, 9, 17, 19, 26, 29, 35, 39, 46, 54, 63,
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
	SimpleScriptParserSemiColon      = 8
	SimpleScriptParserLeftParen      = 9
	SimpleScriptParserRightParen     = 10
	SimpleScriptParserIdentifier     = 11
	SimpleScriptParserWhitespace     = 12
	SimpleScriptParserNewline        = 13
)

// SimpleScriptParser rules.
const (
	SimpleScriptParserRULE_programm       = 0
	SimpleScriptParserRULE_intDeclaration = 1
	SimpleScriptParserRULE_assignment     = 2
	SimpleScriptParserRULE_expression     = 3
	SimpleScriptParserRULE_additive       = 4
	SimpleScriptParserRULE_multiplicative = 5
	SimpleScriptParserRULE_primary        = 6
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

func (s *ProgrammContext) AllIntDeclaration() []IIntDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIntDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IIntDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIntDeclarationContext); ok {
			tst[i] = t.(IIntDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *ProgrammContext) IntDeclaration(i int) IIntDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntDeclarationContext); ok {
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

	return t.(IIntDeclarationContext)
}

func (s *ProgrammContext) AllAssignment() []IAssignmentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignmentContext); ok {
			len++
		}
	}

	tst := make([]IAssignmentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignmentContext); ok {
			tst[i] = t.(IAssignmentContext)
			i++
		}
	}

	return tst
}

func (s *ProgrammContext) Assignment(i int) IAssignmentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
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

	return t.(IAssignmentContext)
}

func (s *ProgrammContext) AllExpression() []IExpressionContext {
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

func (s *ProgrammContext) Expression(i int) IExpressionContext {
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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SimpleScriptParserInt)|(1<<SimpleScriptParserIntegerLiteral)|(1<<SimpleScriptParserLeftParen)|(1<<SimpleScriptParserIdentifier))) != 0 {
		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(14)
				p.IntDeclaration()
			}

		case 2:
			{
				p.SetState(15)
				p.Assignment()
			}

		case 3:
			{
				p.SetState(16)
				p.Expression()
			}

		}

		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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

func (s *IntDeclarationContext) SemiColon() antlr.TerminalNode {
	return s.GetToken(SimpleScriptParserSemiColon, 0)
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
	p.EnterRule(localctx, 2, SimpleScriptParserRULE_intDeclaration)
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
		p.SetState(22)
		p.Match(SimpleScriptParserInt)
	}
	{
		p.SetState(23)
		p.Match(SimpleScriptParserIdentifier)
	}
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleScriptParserAssignmentOP {
		{
			p.SetState(24)
			p.Match(SimpleScriptParserAssignmentOP)
		}
		{
			p.SetState(25)
			p.Additive()
		}

	}
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleScriptParserSemiColon {
		{
			p.SetState(28)
			p.Match(SimpleScriptParserSemiColon)
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

func (s *AssignmentContext) SemiColon() antlr.TerminalNode {
	return s.GetToken(SimpleScriptParserSemiColon, 0)
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
	p.EnterRule(localctx, 4, SimpleScriptParserRULE_assignment)
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
		p.SetState(31)
		p.Match(SimpleScriptParserIdentifier)
	}
	{
		p.SetState(32)
		p.Match(SimpleScriptParserAssignmentOP)
	}
	{
		p.SetState(33)
		p.Additive()
	}
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleScriptParserSemiColon {
		{
			p.SetState(34)
			p.Match(SimpleScriptParserSemiColon)
		}

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

func (s *ExpressionContext) SemiColon() antlr.TerminalNode {
	return s.GetToken(SimpleScriptParserSemiColon, 0)
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
	p.EnterRule(localctx, 6, SimpleScriptParserRULE_expression)
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
		p.SetState(37)
		p.Additive()
	}
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleScriptParserSemiColon {
		{
			p.SetState(38)
			p.Match(SimpleScriptParserSemiColon)
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

func (s *AdditiveContext) AllMultiplicative() []IMultiplicativeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMultiplicativeContext); ok {
			len++
		}
	}

	tst := make([]IMultiplicativeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMultiplicativeContext); ok {
			tst[i] = t.(IMultiplicativeContext)
			i++
		}
	}

	return tst
}

func (s *AdditiveContext) Multiplicative(i int) IMultiplicativeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicativeContext); ok {
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

	return t.(IMultiplicativeContext)
}

func (s *AdditiveContext) AllPlus() []antlr.TerminalNode {
	return s.GetTokens(SimpleScriptParserPlus)
}

func (s *AdditiveContext) Plus(i int) antlr.TerminalNode {
	return s.GetToken(SimpleScriptParserPlus, i)
}

func (s *AdditiveContext) AllMinus() []antlr.TerminalNode {
	return s.GetTokens(SimpleScriptParserMinus)
}

func (s *AdditiveContext) Minus(i int) antlr.TerminalNode {
	return s.GetToken(SimpleScriptParserMinus, i)
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
	this := p
	_ = this

	localctx = NewAdditiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SimpleScriptParserRULE_additive)
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
		p.SetState(41)
		p.Multiplicative()
	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SimpleScriptParserPlus || _la == SimpleScriptParserMinus {
		{
			p.SetState(42)

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
			p.SetState(43)
			p.Multiplicative()
		}

		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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

func (s *MultiplicativeContext) AllPrimary() []IPrimaryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrimaryContext); ok {
			len++
		}
	}

	tst := make([]IPrimaryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrimaryContext); ok {
			tst[i] = t.(IPrimaryContext)
			i++
		}
	}

	return tst
}

func (s *MultiplicativeContext) Primary(i int) IPrimaryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
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

	return t.(IPrimaryContext)
}

func (s *MultiplicativeContext) AllStar() []antlr.TerminalNode {
	return s.GetTokens(SimpleScriptParserStar)
}

func (s *MultiplicativeContext) Star(i int) antlr.TerminalNode {
	return s.GetToken(SimpleScriptParserStar, i)
}

func (s *MultiplicativeContext) AllSlash() []antlr.TerminalNode {
	return s.GetTokens(SimpleScriptParserSlash)
}

func (s *MultiplicativeContext) Slash(i int) antlr.TerminalNode {
	return s.GetToken(SimpleScriptParserSlash, i)
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
	this := p
	_ = this

	localctx = NewMultiplicativeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SimpleScriptParserRULE_multiplicative)
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
		p.SetState(49)
		p.Primary()
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SimpleScriptParserStar || _la == SimpleScriptParserSlash {
		{
			p.SetState(50)

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
			p.SetState(51)
			p.Primary()
		}

		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.EnterRule(localctx, 12, SimpleScriptParserRULE_primary)

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

	p.SetState(63)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SimpleScriptParserIntegerLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(57)
			p.Match(SimpleScriptParserIntegerLiteral)
		}

	case SimpleScriptParserIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(58)
			p.Match(SimpleScriptParserIdentifier)
		}

	case SimpleScriptParserLeftParen:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(59)
			p.Match(SimpleScriptParserLeftParen)
		}
		{
			p.SetState(60)
			p.Expression()
		}
		{
			p.SetState(61)
			p.Match(SimpleScriptParserRightParen)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
