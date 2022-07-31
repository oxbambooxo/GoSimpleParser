// Code generated from SimpleScript.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SimpleScript

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSimpleScriptListener is a complete listener for a parse tree produced by SimpleScriptParser.
type BaseSimpleScriptListener struct{}

var _ SimpleScriptListener = &BaseSimpleScriptListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSimpleScriptListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSimpleScriptListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSimpleScriptListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSimpleScriptListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgramm is called when production programm is entered.
func (s *BaseSimpleScriptListener) EnterProgramm(ctx *ProgrammContext) {}

// ExitProgramm is called when production programm is exited.
func (s *BaseSimpleScriptListener) ExitProgramm(ctx *ProgrammContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseSimpleScriptListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseSimpleScriptListener) ExitStatement(ctx *StatementContext) {}

// EnterIntDeclaration is called when production intDeclaration is entered.
func (s *BaseSimpleScriptListener) EnterIntDeclaration(ctx *IntDeclarationContext) {}

// ExitIntDeclaration is called when production intDeclaration is exited.
func (s *BaseSimpleScriptListener) ExitIntDeclaration(ctx *IntDeclarationContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseSimpleScriptListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseSimpleScriptListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSimpleScriptListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSimpleScriptListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAdditive is called when production additive is entered.
func (s *BaseSimpleScriptListener) EnterAdditive(ctx *AdditiveContext) {}

// ExitAdditive is called when production additive is exited.
func (s *BaseSimpleScriptListener) ExitAdditive(ctx *AdditiveContext) {}

// EnterMultiplicative is called when production multiplicative is entered.
func (s *BaseSimpleScriptListener) EnterMultiplicative(ctx *MultiplicativeContext) {}

// ExitMultiplicative is called when production multiplicative is exited.
func (s *BaseSimpleScriptListener) ExitMultiplicative(ctx *MultiplicativeContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BaseSimpleScriptListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BaseSimpleScriptListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseSimpleScriptListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseSimpleScriptListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseSimpleScriptListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseSimpleScriptListener) ExitExpressionList(ctx *ExpressionListContext) {}
