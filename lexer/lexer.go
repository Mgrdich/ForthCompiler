package lexer

import (
	"CompilerPlayground/token"
	"bufio"
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
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

func lower(ch rune) rune     { return ('a' - 'A') | ch } // returns lower-case ch iff ch is ASCII letter
func isDecimal(ch rune) bool { return '0' <= ch && ch <= '9' }
func isHex(ch rune) bool     { return '0' <= ch && ch <= '9' || 'a' <= lower(ch) && lower(ch) <= 'f' }

func isLetter(ch rune) bool {
	return 'a' <= lower(ch) && lower(ch) <= 'z' || ch == '_' || ch >= utf8.RuneSelf && unicode.IsLetter(ch)
}

func isDigit(ch rune) bool {
	return isDecimal(ch) || ch >= utf8.RuneSelf && unicode.IsDigit(ch)
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

func getRune(bytes []byte) ([]rune, error) {
	var runes []rune

	for len(bytes) > 0 {
		r, size := utf8.DecodeRune(bytes)
		if r == utf8.RuneError && size == 1 {
			return nil, fmt.Errorf("something in getting rune error")
		}
		runes = append(runes, r)
		bytes = bytes[size:]
	}

	return runes, nil
}

func (lex *Lexer) getNextToken() (token.Token, string) {
	ch, err := getRune(lex.scanner.Bytes())

	if err != nil {
		return token.ILLEGAL, ""
	}

	fmt.Println(ch, lex.scanner.Text())

	return token.INTEGER, ""
}

func GetLexer(directory string) Lexer {
	if len(directory) == 0 {
		panic("directory is not defined")
	}

	return Lexer{directory: directory}
}
