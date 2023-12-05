package parser

import (
	"CompilerPlayground/lexer"
	"errors"
	"strings"
)

type Parser struct {
	Tokens []lexer.LexToken
}

func (parser *Parser) Parse() error {
	if len(parser.Tokens) == 0 {
		panic("no items in the tokens")
	}

	var errorBuffer strings.Builder
	for _, item := range parser.Tokens {
		if item.Tok.IsIllegal() {
			errorBuffer.WriteString(item.Lit)
			errorBuffer.WriteString(" ,")
		}
	}

	if errorBuffer.Len() == 0 {
		return nil
	}

	return errors.New("Illegal Tokens: " + errorBuffer.String()[:errorBuffer.Len()-1])
}

func GetParser() *Parser {
	return &Parser{}
}
