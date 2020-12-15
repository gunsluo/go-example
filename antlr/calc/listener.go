package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gunsluo/go-example/antlr/calc/parser"
	"github.com/shopspring/decimal"
)

func main() {
	// Setup the input
	result, err := calc("(1.0 + 2.2) * 3 + 4")
	if err != nil {
		panic(err)
	}
	fmt.Println("calc result:", result)
}

// calc takes a string expression and returns the evaluated result.
func calc(input string) (float64, error) {
	// Setup the input
	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := parser.NewCalcLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewCalcParser(stream)

	// Finally parse the expression (by walking the tree)
	var listener calcListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())

	f := listener.pop()
	num, _ := f.Float64()
	return num, nil
}

type calcListener struct {
	*parser.BaseCalcListener

	stack []decimal.Decimal
}

func (l *calcListener) push(f decimal.Decimal) {
	// fmt.Println("push->", f, l.stack)
	l.stack = append(l.stack, f)
	// fmt.Println("push after->", f, l.stack)
}

func (l *calcListener) pop() decimal.Decimal {
	fmt.Println("pop->", l.stack)
	if len(l.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := l.stack[len(l.stack)-1]

	// Remove the last element from the stack.
	l.stack = l.stack[:len(l.stack)-1]

	return result
}

func (l *calcListener) ExitMulDiv(c *parser.MulDivContext) {
	right, left := l.pop(), l.pop()

	switch c.GetOp().GetTokenType() {
	case parser.CalcParserMUL:
		l.push(left.Mul(right))
	case parser.CalcParserDIV:
		l.push(left.Div(right))
		//l.push(left / right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

func (l *calcListener) ExitAddSub(c *parser.AddSubContext) {
	right, left := l.pop(), l.pop()
	// fmt.Println("ExitAddSub->", right, left)

	switch c.GetOp().GetTokenType() {
	case parser.CalcParserADD:
		l.push(left.Add(right))
		//l.push(left + right)
	case parser.CalcParserSUB:
		l.push(left.Sub(right))
		//l.push(left - right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

func (l *calcListener) ExitNumber(c *parser.NumberContext) {
	fmt.Println("-->", c.GetText())
	f, err := decimal.NewFromString(c.GetText())
	if err != nil {
		panic(err.Error())
	}

	l.push(f)
}
