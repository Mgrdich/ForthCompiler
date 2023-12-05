package main

import "CompilerPlayground/lexer"

func main() {
	lex := lexer.GetLexer("./test1")

	lex.Tokenize()
}
