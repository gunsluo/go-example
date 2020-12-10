package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gunsluo/go-example/antlr/hello/parser"
)

type TreeShapeListener struct {
	*parser.BaseHelloListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	var input = "hello luoji"
	inputStream := antlr.NewInputStream(input)
	lexer := parser.NewHelloLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewHelloParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.R()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
