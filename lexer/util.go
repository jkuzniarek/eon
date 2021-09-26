package lexer

import "eon/token"


func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isQuote(ch byte) bool {
	return ch == '`' || ch == '"' || ch == '\''
}

func newToken(tokenType token.TokenType, ch byte) token.Token{
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func newTokenFromSrc(tokenType token.TokenType, src string, start int, upto int) token.Token{
	return token.Token{Type: tokenType, Literal: src[start:upto] }
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z')
}

func isIdentChar(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || (ch == '_') || ('0' <= ch && ch <= '9')
}

// checks if the input character is an op character NOT if a string is an operator
func isOpChar(ch byte) bool {
	switch ch {
	case '.':
		return true
	case '/':
		return true
	case '#':
		return true
	case '*':
		return true
	case '@':
		return true
	case '|':
		return true
	case '!':
		return true
	case '$':
		return true
	case '%':
		return true
	case '^':
		return true
	case '=':
		return true
	case ':':
		return true
	case '?':
		return true
	case '&':
		return true
	case '+':
		return true
	case '-':
		return true
	case '<':
		return true
	case '>':
		return true
	default:
		return false
	}
}