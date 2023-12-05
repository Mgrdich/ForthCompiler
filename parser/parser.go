package parser

import (
	"CompilerPlayground/lexer"
	"errors"
	"strings"
)

type Parser struct {
	tokens []lexer.LexToken
}

func (parser *Parser) Parse() error {
	if len(parser.tokens) == 0 {
		panic("no items in the tokens")
	}

	var errorBuffer strings.Builder
	for _, item := range parser.tokens {
		if item.Tok.IsIllegal() {
			errorBuffer.WriteString(item.Lit)
		}
	}

	if errorBuffer.Len() == 0 {
		return nil
	}

	return errors.New("Illegal Tokens: " + errorBuffer.String())
}

func GetParser(tokens []lexer.LexToken) *Parser {
	return &Parser{tokens: tokens}
}
