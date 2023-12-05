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
	Tok token.Token
	Lit string
}

type Lexer struct {
	directory string
	scanner   *bufio.Scanner
	Tokens    []LexToken
	chars     []rune // current Token position
	index     int    // current Token index
}

func lower(ch rune) rune     { return ('a' - 'A') | ch } // returns lower-case ch iff ch is ASCII letter
func isDecimal(ch rune) bool { return '0' <= ch && ch <= '9' }
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
		lex.Tokens = append(lex.Tokens, LexToken{
			Tok: tok,
			Lit: lit,
		})
	}

	fmt.Println(lex.Tokens)
}

func convertByteToRune(bytes []byte) ([]rune, error) {
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

func (lex *Lexer) initializeRune(chars []rune) {
	lex.chars = chars
	lex.index = -1
}

func (lex *Lexer) getRune() rune {
	if len(lex.chars) == 0 {
		panic("character is not defined")
	}

	return lex.chars[lex.index]
}

func (lex *Lexer) nextCh() bool {
	if len(lex.chars) == 0 {
		panic("character is not defined")
	}

	if len(lex.chars) == lex.index+1 {
		return false
	}

	lex.index = lex.index + 1

	return true
}

func (lex *Lexer) peekNextCh() rune {
	if len(lex.chars) == 0 {
		panic("character is not defined")
	}

	if len(lex.chars) == lex.index+1 {
		return 0
	}

	return lex.chars[lex.index+1]
}

func (lex *Lexer) scanIdentifier() token.Token {
	builtWord := []rune{lex.getRune()}

	for lex.nextCh() && (isLetter(lex.getRune()) || isDigit(lex.getRune())) {
		builtWord = append(builtWord, lex.getRune())
	}

	str := string(builtWord)

	return token.Lookup(str)
}

func (lex *Lexer) digits(builtNumber *[]rune, state token.Token) token.Token {
	tokenState := state
	m := rune('0' + 10)
	for lex.nextCh() {
		if lex.getRune() >= m || !isDecimal(lex.getRune()) {
			tokenState = token.ILLEGAL
		}
		*builtNumber = append(*builtNumber, lex.getRune())
	}
	return tokenState
}

func (lex *Lexer) scanNumber() token.Token {
	tok := token.ILLEGAL
	builtNumber := []rune{lex.getRune()} // TODO research where we can keep it as byte

	peekedNextCh := lex.peekNextCh()

	if peekedNextCh != '.' {
		tok = token.INTEGER
	}

	if peekedNextCh == '.' {
		builtNumber = append(builtNumber, lex.getRune()) // add the dot notation
		lex.nextCh()
		tok = token.FLOAT
	}

	tok = lex.digits(&builtNumber, tok)

	return tok
}

// getNextToken this will tokenize by iterating between spaces
func (lex *Lexer) getNextToken() (token.Token, string) {
	chars, err := convertByteToRune(lex.scanner.Bytes())
	tok := token.ILLEGAL

	if err != nil {
		return tok, lex.scanner.Text()
	}

	if len(chars) == 1 {
		c := chars[0]
		switch c {
		case '+':
			tok = token.ADD
		case '-':
			tok = token.SUB
		case '*':
			tok = token.MUL
		case '/':
			tok = token.QUO
		case '.':
			tok = token.DOT
		default:
			if isDecimal(c) {
				tok = token.INTEGER
			}
		}

		return tok, lex.scanner.Text()
	}

	// bigger tokens
	lex.initializeRune(chars)
	lex.nextCh()

	switch c := lex.getRune(); {
	case isLetter(c):
		tok = lex.scanIdentifier()
	case isDecimal(c) || (c == '.' && isDecimal(lex.peekNextCh())):
		tok = lex.scanNumber()
	case c == '.':
		// move
		lex.nextCh()
		if lex.getRune() == 's' {
			tok = token.PRINT
		}
	case c == '-':
		lex.nextCh()
		if isDecimal(lex.getRune()) {
			tok = lex.scanNumber()
		}
	}

	return tok, lex.scanner.Text()
}

func GetLexer(directory string) *Lexer {
	if len(directory) == 0 {
		panic("directory is not defined")
	}

	return &Lexer{directory: directory}
}
