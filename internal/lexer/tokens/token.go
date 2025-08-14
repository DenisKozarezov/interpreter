package tokens

// Token is a data structure that represents the minimum semantic unit of every source code, the lexeme.
// A text program is divided into a set of simple lexemes of the language:
// - constants;
// - identifiers (variables);
// - keywords;
// - special characters, etc.
//
// Just as punctuation marks, prepositions, and words form a sentence, programming language constructs are also formed:
//
//	let x = 5;
//	^   ^ ^ ^^
//	1   2 3 45
//
// Where:
//  1. 'let' is a keyword;
//  2. 'x' is a variable identifier;
//  3. '=' is an assignment operator;
//  4. '5' is an expression returning an integer constant;
//  5. ';' is the end of the statement.
type Token struct {
	// Type is a type of token: constant, keyword, etc.
	Type TokenType

	// Literal is the meaning of a token. For example: int a = 5, where `int` is a keyword, `a` is
	// an identifier, and `5` is a literal.
	Literal string
}

func NewToken(tokenType TokenType, literal string) Token {
	return Token{Type: tokenType, Literal: literal}
}
