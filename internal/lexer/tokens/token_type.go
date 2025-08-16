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
	AMPERSAND
	PIPE

	LT
	LT_EQ
	GT
	GT_EQ
	EQ
	NOT_EQ
	AND
	OR

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
