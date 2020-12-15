package main

import (
	"fmt"
	"math/big"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gunsluo/go-example/antlr/calc/parser"
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

	bf := listener.pop()
	f, _ := bf.Float64()
	return f, nil
}

type calcListener struct {
	*parser.BaseCalcListener

	stack []*big.Float
}

func (l *calcListener) push(f *big.Float) {
	// fmt.Println("push->", f, l.stack)
	l.stack = append(l.stack, f)
	// fmt.Println("push after->", f, l.stack)
}

func (l *calcListener) pop() *big.Float {
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
		f1 := &big.Float{}
		f1.Mul(left, right)
		l.push(f1)
	case parser.CalcParserDIV:
		f1 := &big.Float{}
		f1.Quo(left, right)
		l.push(f1)
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
		f1 := &big.Float{}
		f1.Add(left, right)
		l.push(f1)
		//l.push(left + right)
	case parser.CalcParserSUB:
		f1 := &big.Float{}
		f1.Sub(left, right)
		l.push(f1)
		//l.push(left - right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

func (l *calcListener) ExitNumber(c *parser.NumberContext) {
	fmt.Println("-->", c.GetText())
	f, _, err := big.ParseFloat(c.GetText(), 10, 0, big.ToNearestEven)
	if err != nil {
		panic(err.Error())
	}

	l.push(f)
}
