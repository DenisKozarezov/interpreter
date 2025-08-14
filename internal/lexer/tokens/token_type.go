package tokens

type TokenType = int8

const (
	ILLEGAL TokenType = iota
	EOF

	// Identifiers
	INT
	STRING
	IDENTIFIER

	// Arithmetic operators
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

	// Separators
	COMMA
	SEMICOLON

	LPAREN
	RPAREN
	LBRACE
	RBRACE
	LBRACKET
	RBRACKET

	// Keywords
	FUNCTION
	LET
	TRUE
	FALSE
	IF
	ELSE
	RETURN
)

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
