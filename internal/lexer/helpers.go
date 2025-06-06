package lexer

func isLetter(symbol Symbol) bool {
	return 'a' <= symbol && symbol <= 'z' || 'A' <= symbol && symbol <= 'Z' || symbol == '_'
}

func isDigit(symbol Symbol) bool {
	return '0' <= symbol && symbol <= '9'
}
