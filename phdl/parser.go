package phdl

import (
	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/participle/lexer/ebnf"
)

var (
	Lexer = lexer.Must(ebnf.New(`
		OneLineComment = "//" { "\u0000"…"\uffff"-"\n" } .
		MultiLineComment = "/*" { "\u0000"…"\uffff"-"*/" } "*/" .
		Ident1 = ( block | test ) ( alpha | "_" | digit ) { "_" | alpha | digit } .
		Ident2 = "d" { digit } ( alpha | "_" ) { "_" | alpha | digit } .
		Type = "d" digit { digit } .
		Block = block .
		Test = test .
		Ident3 =  ( alpha | "_" ) { "_" | alpha | digit } .
		Number = [ "-" ] digit [ "x" | "o" | "b" ] { hexdig } .
		Whitespace = " " | "\t" | "\n" | "\r" .
		Lparen = "(" .
		Rparen = ")" .
		Lbrace = "{" .
		Rbrace = "}" .
		Lbrak = "[" .
		Rbrak = "]" .
		Comma = "," .
		Ellipsis = ".." .
		Semicolon = ";" .
		Arrow = "->" .
		TestArrow = "==>" .

		block = "block" .
		test = "test" .
		hexdig = digit | "a"…"f" | "A"…"F" .
		alpha = "a"…"z" | "A"…"Z" .
		digit = "0"…"9" .
	`))
	Parser = participle.MustBuild(
		&File{},
		participle.Lexer(Lexer),
		participle.Elide("Whitespace", "OneLineComment", "MultiLineComment"),
	)

)
