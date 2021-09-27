package lexer

import tk "eon/token"

func (l *Lexer) NextToken() tk.Token{
	var tok tk.Token

	l.skipWhitespace()

	switch l.ch{
	case '/':
		if l.peekChar() == '/' {
			tok = tk.Token{Cat: tk.COMMENT, Type: tk.COMMENT, Literal: l.readCommentLine()}
		} else if l.peekChar() == '*' {
			tok = tk.Token{Cat: tk.COMMENT, Type: tk.COMMENT, Literal: l.readCommentMultiline()}
		} else {
			tok = newToken(tk.ACCESS_OPERATOR, tk.SLASH, l.ch)
		}
	case '\n':
		tok = newToken(tk.EOL, tk.EOL, l.ch)
	case ':':
		if l.peekChar() == ':' {
			tok = newTokenFromSrc(tk.ASSIGN_OPERATOR, tk.SET_CONST, l.input, l.position, l.position+2)
			l.readChar()
		} else if l.peekChar() == '?' {
			tok = newTokenFromSrc(tk.ASSIGN_OPERATOR, tk.SET_WEAK, l.input, l.position, l.position+2)
			l.readChar()
		} else if l.peekChar() == '&' {
			tok = newTokenFromSrc(tk.ASSIGN_OPERATOR, tk.SET_BIND, l.input, l.position, l.position+2)
			l.readChar()
		}  else if l.peekChar() == '+' {
			tok = newTokenFromSrc(tk.ASSIGN_OPERATOR, tk.SET_PLUS, l.input, l.position, l.position+2)
			l.readChar()
		}  else if l.peekChar() == '-' {
			tok = newTokenFromSrc(tk.ASSIGN_OPERATOR, tk.SET_MINUS, l.input, l.position, l.position+2)
			l.readChar()
		}  else if l.peekChar() == '#' {
			tok = newTokenFromSrc(tk.ASSIGN_OPERATOR, tk.SET_TYPE, l.input, l.position, l.position+2)
			l.readChar()
		} else {
			tok = newToken(tk.ASSIGN_OPERATOR, tk.SET_VAL, l.ch)
		}
	case '.':
		tok = newToken(tk.ACCESS_OPERATOR, tk.DOT, l.ch)
	case '|':
		tok = newToken(tk.EVAL_OPERATOR, tk.PIPE, l.ch) 
	case '<':
		if l.peekChar() == '=' {
			tok = newTokenFromSrc(tk.EVAL_OPERATOR, tk.LT_EQ, l.input, l.position, l.position+2)
			l.readChar()
		} else {
			tok = newToken(tk.EVAL_OPERATOR, tk.LT, l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			tok = newTokenFromSrc(tk.EVAL_OPERATOR, tk.GT_EQ, l.input, l.position, l.position+2)
			l.readChar()
		} else {
			tok = newToken(tk.EVAL_OPERATOR, tk.GT, l.ch)
		}
	case '(':
		if l.peekChar() == ':' {
			tok = newTokenFromSrc(tk.OPEN_DELIMITER, tk.CPAREN , l.input, l.position, l.position+2)
			l.readChar()
		} else if l.peekChar() == '-' {
			tok = newTokenFromSrc(tk.OPEN_DELIMITER, tk.HPAREN , l.input, l.position, l.position+2)
			l.readChar()
		} else {
			tok = newToken(tk.OPEN_DELIMITER, tk.LPAREN, l.ch)
		}
	case ')':
		tok = newToken(tk.CLOSE_DELIMITER, tk.RPAREN, l.ch)
	case '{':
		if l.peekChar() == ':' {
			tok = newTokenFromSrc(tk.OPEN_DELIMITER, tk.SCURLY, l.input, l.position, l.position+2)
			l.readChar()
		} else {
			tok = newToken(tk.OPEN_DELIMITER, tk.LCURLY, l.ch)
		}
	case '}':
		tok = newToken(tk.CLOSE_DELIMITER, tk.RCURLY, l.ch)
	case '[':
		tok = newToken(tk.OPEN_DELIMITER, tk.LSQUAR, l.ch)
	case ']':
		tok = newToken(tk.CLOSE_DELIMITER, tk.RSQUAR, l.ch)
	case '+':
		if isDigit(l.peekChar()){
			tok.Type = tk.SINT
			tok.Cat = tk.PRIMITIVE
			l.readChar()
			tok.Literal = "+" + l.readNumber()
		}
	case '-':
		if isDigit(l.peekChar()){
			tok.Type = tk.SINT
			tok.Cat = tk.PRIMITIVE
			l.readChar()
			tok.Literal = "-" + l.readNumber()
		}
	case '$':
		tok = newToken(tk.KEYWORD, tk.DOLLAR, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = tk.EOF
		tok.Cat = tk.EOF
	default:
		if isLetter(l.ch){
			if l.peekLChar() == '<' {
				tok.Literal = l.readIdentifier()
				tok.Type = tk.TYPE
				tok.Cat = tk.TYPE
			} else {
				tok.Literal = l.readIdentifier()
				tok.Type = tk.LookupName(tok.Literal)
				if tk.IsKeyword(tok.Literal){
					tok.Cat = tk.KEYWORD
				}else{
					tok.Cat = tk.NAME
				}
			}
		} else if isDigit(l.ch){
			tok.Type = tk.UINT
			tok.Literal = l.readNumber()
			tok.Cat = tk.PRIMITIVE
		} else if isOpChar(l.ch){
			tok.Literal = l.readOperator()
			tok.Type = tk.LookupOp(tok.Literal)
			if tk.IsAccessOp(tok.Literal){
				tok.Cat = tk.ACCESS_OPERATOR
			} else if tk.IsAssignOp(tok.Literal){
				tok.Cat = tk.ASSIGN_OPERATOR
			} else if tk.IsEvalOp(tok.Literal){
				tok.Cat = tk.EVAL_OPERATOR
			}else{
				tok.Cat = tk.ILLEGAL
			}
		} else if isQuote(l.ch){
			tok.Type = tk.STR 
			tok.Literal = l.readString(l.ch)
			tok.Cat = tk.PRIMITIVE
		} else {
			tok = newToken(tk.ILLEGAL, tk.ILLEGAL, l.ch)
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