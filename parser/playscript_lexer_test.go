package parser

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func ExampleNewPlayScriptLexer() {
	is := antlr.NewInputStream("1 + 2 * 3")

	lexer := NewPlayScriptLexer(is)

	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}

		fmt.Printf("%s (%q)\n", lexer.SymbolicNames[t.GetTokenType()],
			t.GetText())
	}
	// Output:
	//DECIMAL_LITERAL ("1")
	//WS (" ")
	//ADD ("+")
	//WS (" ")
	//DECIMAL_LITERAL ("2")
	//WS (" ")
	//MUL ("*")
	//WS (" ")
	//DECIMAL_LITERAL ("3")
}
