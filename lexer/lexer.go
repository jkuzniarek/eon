package lexer

import "eon/token"

func (l *Lexer) NextToken() token.Token{
	var tok token.Token

	l.skipWhitespace()

	switch l.ch{
	case '/':
		if l.peekChar() == '/' {
			tok = token.Token{Type: token.COMMENT, Literal: l.readCommentLine()}
		} else if l.peekChar() == '*' {
			tok = token.Token{Type: token.COMMENT, Literal: l.readCommentMultiline()}
		} else {
			tok = newToken(token.SLASH, l.ch)
		}
	case '\n':
		tok = newToken(token.EOL, l.ch)
	case ':':
		if l.peekChar() == ':' {
			tok = newTokenFromSrc(token.SET_CONST, l.input, l.position, l.position+2)
			l.readChar()
		} else if l.peekChar() == '?' {
			tok = newTokenFromSrc(token.SET_WEAK, l.input, l.position, l.position+2)
			l.readChar()
		} else if l.peekChar() == '&' {
			tok = newTokenFromSrc(token.SET_BIND, l.input, l.position, l.position+2)
			l.readChar()
		}  else if l.peekChar() == '+' {
			tok = newTokenFromSrc(token.SET_PLUS, l.input, l.position, l.position+2)
			l.readChar()
		}  else if l.peekChar() == '-' {
			tok = newTokenFromSrc(token.SET_MINUS, l.input, l.position, l.position+2)
			l.readChar()
		}  else if l.peekChar() == '#' {
			tok = newTokenFromSrc(token.SET_TYPE, l.input, l.position, l.position+2)
			l.readChar()
		} else {
			tok = newToken(token.SET_VAL, l.ch)
		}
	case '.':
		tok = newToken(token.DOT, l.ch)
	case '|':
		tok = newToken(token.PIPE, l.ch) 
	case '<':
		if l.peekChar() == '=' {
			tok = newTokenFromSrc(token.LT_EQ, l.input, l.position, l.position+2)
			l.readChar()
		} else {
			tok = newToken(token.LT, l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			tok = newTokenFromSrc(token.GT_EQ, l.input, l.position, l.position+2)
			l.readChar()
		} else {
			tok = newToken(token.GT, l.ch)
		}
	case '(':
		if l.peekChar() == ':' {
			tok = newTokenFromSrc(token.CPAREN , l.input, l.position, l.position+2)
			l.readChar()
		} else {
			tok = newToken(token.LPAREN, l.ch)
		}
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		if l.peekChar() == ':' {
			tok = newTokenFromSrc(token.SCURLY, l.input, l.position, l.position+2)
			l.readChar()
		} else {
			tok = newToken(token.LCURLY, l.ch)
		}
	case '}':
		tok = newToken(token.RCURLY, l.ch)
	case '[':
		tok = newToken(token.LSQUAR, l.ch)
	case ']':
		tok = newToken(token.RSQUAR, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch){
			if l.peekLChar() == '<' {
				tok.Literal = l.readIdentifier()
				tok.Type = token.TYPE
			} else {
				tok.Literal = l.readIdentifier()
				tok.Type = token.LookupName(tok.Literal)
			}
		} else if isDigit(l.ch){
			tok.Type = token.INT
			tok.Literal = l.readNumber()
		} else if isOpChar(l.ch){
			tok.Literal = l.readOperator()
			tok.Type = token.LookupOp(tok.Literal)
		} else if isQuote(l.ch){
			tok.Type = token.STR 
			tok.Literal = l.readString(l.ch)
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string{
	position := l.position
	for isIdentChar(l.peekChar()){
		l.readChar()
	}
	ident := l.input[position:(l.position+1)]
	return ident
}

func (l *Lexer) readOperator() string{
	op := ""
	// operators can only be up to 2 characters long
	if isOpChar(l.peekChar()) {
		op = string(l.ch) + string(l.peekChar())
		l.readChar()
	} else {
		op = string(l.ch)
	}
	return op
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

func newToken(tokenType token.TokenType, ch byte) token.Token{
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func newTokenFromSrc(tokenType token.TokenType, src string, start int, upto int) token.Token{
	return token.Token{Type: tokenType, Literal: src[start:upto] }
}

type Lexer struct{
	input string
	position int // current position in input (points to current char)
	readPosition int // current reading position in input (after current char)
	ch byte // current char under examination
	row int // current line number (# of \n consumed + 1)
	col int // current character index in line (# of ch bytes consumed + 1)
}

func New(input string) *Lexer{
	// L
	l := &Lexer{input: input, row: 1, col: 1}
	l.readChar()
	return l
}

func (l *Lexer) readChar(){
	// tracks row & col
	if l.peekLChar() == '\n' {
		l.row++
		l.col = 1
	} else {
		l.col++
	}

	if l.readPosition >= len(l.input){
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string{
	position := l.position
	for isDigit(l.peekChar()){
		l.readChar()
	}
	return l.input[position:l.position+1]
}

func (l *Lexer) readCommentLine() string{
	l.readChar()
	l.readChar()
	position := l.position
	for l.ch != '\n' {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readCommentMultiline() string{
	l.readChar()
	l.readChar()
	position := l.position
	for !(l.ch == '*' && l.peekChar() == '/'){
		l.readChar()
	}
	l.readChar()
	return l.input[position:l.position-1]
}

func (l *Lexer) readString(ch byte) string{
	position := l.position
	l.readChar()
	for l.ch != ch {
		l.readChar()
	}
	return l.input[position:l.position+1]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isQuote(ch byte) bool {
	return ch == '`' || ch == '"' || ch == '\''
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) peekLChar() byte {
	if l.position == 0 {
		return 0
	} else {
		return l.input[l.position-1]
	}
}