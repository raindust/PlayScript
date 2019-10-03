package parser

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strings"
)

type calcVisitor struct {
	*BasePlayScriptVisitor

	stack strings.Builder
}

func (v *calcVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	switch ctx.GetBop().GetTokenType() {
	case PlayScriptLexerADD, PlayScriptLexerSUB, PlayScriptLexerMUL,
		PlayScriptLexerDIV:
		v.stack.WriteString(ctx.GetText())
	}

	return v.VisitChildren(ctx)
}

func (v *calcVisitor) VisitIntLiteral(ctx *IntLiteralContext) interface{} {
	v.stack.WriteString(ctx.GetText())
	return v.VisitChildren(ctx)
}

func ExampleNewPlayScriptParser() {
	// setup input
	is := antlr.NewInputStream("1 + 2 * 3")

	// create lexer
	lexer := NewPlayScriptLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// create the parser
	p := NewPlayScriptParser(stream)

	// finally parse the expression
	v := calcVisitor{BasePlayScriptVisitor: &BasePlayScriptVisitor{}}
	p.Expression().Accept(&v)

	fmt.Println(v.stack.String())
	// Output:
	// 1+2*3
}
