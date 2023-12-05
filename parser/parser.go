package parser

import "CompilerPlayground/lexer"

type Parser struct {
	tokens []lexer.LexToken
}

func (parser *Parser) Parse() {

}

func GetParser(tokens []lexer.LexToken) *Parser {
	return &Parser{tokens: tokens}
}
