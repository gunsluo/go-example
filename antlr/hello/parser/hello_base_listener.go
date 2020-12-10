// Code generated from Hello.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Hello

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseHelloListener is a complete listener for a parse tree produced by HelloParser.
type BaseHelloListener struct{}

var _ HelloListener = &BaseHelloListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseHelloListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseHelloListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseHelloListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseHelloListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterR is called when production r is entered.
func (s *BaseHelloListener) EnterR(ctx *RContext) {}

// ExitR is called when production r is exited.
func (s *BaseHelloListener) ExitR(ctx *RContext) {}
