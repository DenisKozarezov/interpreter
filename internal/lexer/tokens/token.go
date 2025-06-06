package tokens

// Token структура данных, составляющая минимальную смысловую единицу исходного кода.
// Текстовая программа разбивается на множество простейших компонентов:
//   - константы;
//   - идентификаторы (переменные);
//   - ключевые слова;
//   - спецсимволы и т.п.
type Token struct {
	// Type тип токена: константа, ключевое слово и т.д.
	Type TokenType

	// Literal значение токена. Например: int a = 5, где `int` - ключевое слово, `а` - идентификатор, а `5` - литерал.
	Literal string
}

func NewToken(tokenType TokenType, literal string) Token {
	return Token{Type: tokenType, Literal: literal}
}
