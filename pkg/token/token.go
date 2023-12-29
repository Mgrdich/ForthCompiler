package token

// Token is the set of lexical tokens of our created language
type Token int

// The list of tokens.

const (
	// Special tokens

	ILLEGAL Token = iota

	//  Identifiers and basic type literals
	// (these tokens stand for classes of literals)
	literalBeg
	PRINT
	IDENT
	numBeg
	INTEGER // 12345
	FLOAT
	numEnd // 123.323
	literalEnd

	keywordBeg

	keywordEnd

	// Operators and delimiters
	operatorBeg

	operatorSimpleBeg
	ADD
	SUB
	MUL
	QUO
	operatorSimpleEnd

	operatorKeywordBeg
	MOD
	NEGATE
	ABS
	MAX
	MIN
	DUP
	SWAP
	ROT
	DROP
	NIP
	TUCK
	OVER
	ROLL
	PICK
	operatorKeywordEnd

	COLON
	DOT
	SEMICOLON
	operatorEnd
)

var tokens = [...]string{

	ILLEGAL: "ILLEGAL",

	PRINT: ".s",
	IDENT: "IDENT",

	ADD:       "+",
	SUB:       "-",
	MUL:       "*",
	QUO:       "/",
	MOD:       "mod",
	NEGATE:    "negate",
	ABS:       "abs",
	MAX:       "max",
	MIN:       "min",
	DUP:       "dup",
	SWAP:      "swap",
	ROT:       "rot",
	DROP:      "drop",
	NIP:       "nip",
	TUCK:      "tuck",
	OVER:      "over",
	ROLL:      "roll",
	PICK:      "pick",
	DOT:       ".",
	SEMICOLON: ";",
	COLON:     ":",
}

var keywords map[string]Token

func init() {
	keywords = make(map[string]Token)
	for i := keywordBeg + 1; i < keywordEnd; i++ {
		keywords[tokens[i]] = i
	}

	for i := operatorKeywordBeg + 1; i < operatorKeywordEnd; i++ {
		keywords[tokens[i]] = i
	}
}

// Lookup maps an identifier to its keyword token or IDENT (if not a keyword).
func Lookup(ident string) Token {
	if tok, isKeyword := keywords[ident]; isKeyword {
		return tok
	}
	return IDENT
}

// IsLiteral returns true for tokens corresponding to identifiers
// and basic type literals; it returns false otherwise.
func (tok Token) IsLiteral() bool { return literalBeg < tok && tok < literalEnd }

// IsOperator returns true for tokens corresponding to operators and
// delimiters; it returns false otherwise.
func (tok Token) IsOperator() bool {
	return operatorBeg < tok && tok < operatorEnd
}

func (tok Token) IsSimpleOperator() bool {
	return operatorSimpleBeg < tok && tok < operatorSimpleEnd
}

func (tok Token) IsKeywordOperator() bool {
	return operatorKeywordBeg < tok && tok < operatorKeywordEnd
}

// IsNumber returns true for tokens corresponding to a numbers
func (tok Token) IsNumber() bool {
	return numBeg < tok && tok < numEnd
}

// IsKeyword returns true for tokens corresponding to keywords;
// it returns false otherwise.
func (tok Token) IsKeyword() bool { return keywordBeg < tok && tok < keywordEnd }

func (tok Token) IsIllegal() bool {
	return tok == ILLEGAL
}
