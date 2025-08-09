package tokens

type TokenType = int8

const (
	ILLEGAL TokenType = iota
	EOF

	// Идентификаторы
	INT
	STRING
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
	COMMENT_LINE
	COMMENT_BEGIN
	COMMENT_END

	// Ключевые слова
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
