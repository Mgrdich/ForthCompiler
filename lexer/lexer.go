package lexer

import (
	"CompilerPlayground/token"
	"bufio"
	"fmt"
	"os"
)

type LexToken struct {
	tok token.Token
	lit string
}

type Lexer struct {
	directory string
	scanner   *bufio.Scanner
	tokens    []LexToken
}

// Tokenize will tokenize the whole file
func (lex *Lexer) Tokenize() {
	if len(lex.directory) == 0 {
		panic("directory is not defined")
	}

	f, err := os.Open(lex.directory)
	if err != nil {
		panic("something wrong with the provided directory")
	}

	// This is used since forth is very simple language and everything is seperated by newline or space
	scanner := bufio.NewScanner(f)
	// This already work with spaces under the hood
	scanner.Split(bufio.ScanWords)

	lex.scanner = scanner
	lex.scanStart()
}

func (lex *Lexer) scanStart() {
	for lex.scanner.Scan() {
		tok, lit := lex.getNextToken()
		lex.tokens = append(lex.tokens, LexToken{
			tok: tok,
			lit: lit,
		})
	}
}

func (lex *Lexer) getNextToken() (token.Token, string) {
	ch := lex.scanner.Bytes() // TODO maybe keep the bytes in the array for next peek functionality

	fmt.Println(ch, lex.scanner.Text())

	return token.INTEGER, ""
}

func GetLexer(directory string) Lexer {
	if len(directory) == 0 {
		panic("directory is not defined")
	}

	return Lexer{directory: directory}
}
