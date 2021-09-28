package lexer

import tk "eon/token"


func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isQuote(ch byte) bool {
	return ch == '`' || ch == '"' || ch == '\''
}

func newToken(cat tk.TokenType, typ tk.TokenType, ch byte) tk.Token{
	return tk.Token{Cat: cat, Type: typ, Literal: string(ch)}
}

func newTokenFromSrc(cat tk.TokenType, typ tk.TokenType, src *string, start int, upto int) tk.Token{
	return tk.Token{Cat: cat, Type: typ, Literal: src[start:upto] }
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z')
}

func isIdentChar(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || (ch == '_') || ('0' <= ch && ch <= '9') || ch == '$'
}

// checks if the input character is an op character NOT if a string is an operator
func isOpChar(ch byte) bool {
	switch ch {
	case '.', '/', '#', '*', '@', '|', '!', '%', '=', ':', '?', '&', '+', '-', '<', '>':
		return true
	default:
		return false
	}
}