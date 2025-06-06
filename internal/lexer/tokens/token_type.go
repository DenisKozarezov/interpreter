package tokens

type TokenType = int8

const (
	ILLEGAL TokenType = iota
	EOF

	// Идентификаторы
	INT
	IDENTIFIER

	// Арифметические операторы
	ASSIGN
	PLUS
	MINUS
	BANG
	ASTERISK
	SLASH

	LT
	GT
	EQ
	NOT_EQ

	// Разделители
	COMMA
	SEMICOLON

	LPAREN
	RPAREN
	LBRACE
	RBRACE
	LBRACKET
	RBRACKET

	// Ключевые слова
	FUNCTION
	LET
	TRUE
	FALSE
	IF
	ELSE
	RETURN
)

var tokenTypes = map[string]TokenType{
	"=": ASSIGN,
	"+": PLUS,
	"-": MINUS,
	"!": BANG,
	"*": ASTERISK,
	"/": SLASH,
	"<": LT,
	">": GT,

	"==": EQ,
	"!=": NOT_EQ,

	",": COMMA,
	";": SEMICOLON,
	"(": LPAREN,
	")": RPAREN,
	"{": LBRACE,
	"}": RBRACE,
	"[": LBRACKET,
	"]": RBRACKET,
}

func LookupTokenType(literal string) (TokenType, bool) {
	t, found := tokenTypes[literal]
	return t, found
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdentifierType(literal string) TokenType {
	if token, found := keywords[literal]; found {
		return token
	}
	return IDENTIFIER
}
