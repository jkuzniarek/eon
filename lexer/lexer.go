package lexer

import tk "eon/token"

type Lexer struct{
	input string
	position int // current position in input (points to current char)
	readPosition int // current reading position in input (after current char)
	ch byte // current char under examination
	row int // current line number (# of \n consumed + 1)
	col int // current character index in line (# of ch bytes consumed + 1)
	Depth int
}

func (l *Lexer) NextToken() tk.Token{
	var tok tk.Token

	l.skipWhitespace()

	switch l.ch{
	case '/':
		if l.peekChar() == '/' {
			l.Depth++
			tok = tk.Token{Cat: tk.COMMENT, Type: tk.COMMENT, Literal: l.readCommentLine()}
		} else if l.peekChar() == '*' {
			l.Depth++
			tok = tk.Token{Cat: tk.COMMENT, Type: tk.COMMENT, Literal: l.readCommentMultiline()}
		} else {
			tok = newToken(tk.OPERATOR, tk.ACCESS_OPERATOR, l.ch)
		}
	case '\n':
		tok = newToken(tk.EOL, tk.EOL, l.ch)
	case ':':
		if l.peekChar() == ':' {
			tok = newTokenFromSrc(tk.OPERATOR, tk.ASSIGN_OPERATOR, &l.input, l.position, l.position+2)
			l.readChar()
		} else if l.peekChar() == '?' {
			tok = newTokenFromSrc(tk.OPERATOR, tk.ASSIGN_OPERATOR, &l.input, l.position, l.position+2)
			l.readChar()
		} else if l.peekChar() == '&' {
			tok = newTokenFromSrc(tk.OPERATOR, tk.ASSIGN_OPERATOR, &l.input, l.position, l.position+2)
			l.readChar()
		}  else if l.peekChar() == '+' {
			tok = newTokenFromSrc(tk.OPERATOR, tk.ASSIGN_OPERATOR, &l.input, l.position, l.position+2)
			l.readChar()
		}  else if l.peekChar() == '-' {
			tok = newTokenFromSrc(tk.OPERATOR, tk.ASSIGN_OPERATOR, &l.input, l.position, l.position+2)
			l.readChar()
		}  else if l.peekChar() == '#' {
			tok = newTokenFromSrc(tk.OPERATOR, tk.ASSIGN_OPERATOR, &l.input, l.position, l.position+2)
			l.readChar()
		} else {
			tok = newToken(tk.OPERATOR, tk.ASSIGN_OPERATOR, l.ch)
		}
	case '.':
		tok = newToken(tk.OPERATOR, tk.ACCESS_OPERATOR, l.ch)
	case '|':
		tok = newToken(tk.OPERATOR, tk.EVAL_OPERATOR, l.ch) 
	case '<':
		if l.peekChar() == '=' {
			tok = newTokenFromSrc(tk.OPERATOR, tk.EVAL_OPERATOR, &l.input, l.position, l.position+2)
			l.readChar()
		} else {
			tok = newToken(tk.OPERATOR, tk.EVAL_OPERATOR, l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			tok = newTokenFromSrc(tk.OPERATOR, tk.EVAL_OPERATOR, &l.input, l.position, l.position+2)
			l.readChar()
		} else {
			tok = newToken(tk.OPERATOR, tk.EVAL_OPERATOR, l.ch)
		}
	case '(':
		if l.peekChar() == ':' {
			tok = newTokenFromSrc(tk.OPEN_DELIMITER, tk.CPAREN , &l.input, l.position, l.position+2)
			l.readChar()
		} else if l.peekChar() == '-' {
			tok = newTokenFromSrc(tk.OPEN_DELIMITER, tk.HPAREN , &l.input, l.position, l.position+2)
			l.readChar()
		} else {
			tok = newToken(tk.OPEN_DELIMITER, tk.LPAREN, l.ch)
		}
	case ')':
		tok = newToken(tk.CLOSE_DELIMITER, tk.RPAREN, l.ch)
	case '{':
		if l.peekChar() == ':' {
			tok = newTokenFromSrc(tk.OPEN_DELIMITER, tk.SCURLY, &l.input, l.position, l.position+2)
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
		if isWS(l.peekLChar()) && isDigit(l.peekChar()){
			tok = l.newSNumber()
		}
	case '-':
		if isWS(l.peekLChar()) && isDigit(l.peekChar()){
			tok = l.newSNumber()
		}
	case '$':
		tok = newToken(tk.NAME, tk.KEYWORD, l.ch)
	case '\\':
		if l.peekChar() == 'x' || l.peekChar() == 'd' || l.peekChar() == 'b' {
			l.Depth++
			tok = tk.Token{Cat: tk.PRIMITIVE, Type: tk.BYTES, Literal: l.readBytes()}
		} else {
			tok = newToken(tk.BSLASH, tk.BSLASH, l.ch)
		}
	case 0:
		tok.Literal = ""
		tok.Type = tk.EOF
		tok.Cat = tk.EOF
	default:
		if isLetter(l.ch){
			tok.Literal = l.readIdentifier()
			tok.Cat = tk.NAME
			if tk.IsKeyword(tok.Literal){
				tok.Type = tk.KEYWORD
			}else{
				tok.Type = tk.NAME
			}
		} else if isDigit(l.ch){
			tok = l.newUNumber()
		} else if isOpChar(l.ch){
			tok.Literal = l.readOperator()
			if tk.IsOperator(tok.Literal){
				tok.Cat = tk.OPERATOR
				if tk.IsAccessOp(tok.Literal){
					tok.Type = tk.ACCESS_OPERATOR
				} else if tk.IsAssignOp(tok.Literal){
					tok.Type = tk.ASSIGN_OPERATOR
				} else if tk.IsEvalOp(tok.Literal){
					tok.Type = tk.EVAL_OPERATOR
				}else{
					tok.Type = tk.ILLEGAL
				}
			}else{
				tok.Type = tk.ILLEGAL
				tok.Cat = tk.ILLEGAL
			}
		} else if isQuote(l.ch){
			l.Depth++
			tok.Type = tk.STR 
			tok.Literal = l.readString()
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

func New(input string) *Lexer{
	// L
	l := &Lexer{input: input, row: 1, col: 1, Depth: 0}
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

func (l *Lexer) newUNumber() tk.Token{
	position := l.position
	for isDigit(l.peekChar()){
		l.readChar()
	}

	if l.peekChar() == '.'{
		l.readChar()
		for isDigit(l.peekChar()){
			l.readChar()
		}
		return newTokenFromSrc(tk.PRIMITIVE, tk.DEC, &l.input, position, l.position+1)
	}else{
		return newTokenFromSrc(tk.PRIMITIVE, tk.UINT, &l.input, position, l.position+1)
	}
}

func (l *Lexer) newSNumber() tk.Token{
	position := l.position
	for isDigit(l.peekChar()){
		l.readChar()
	}

	if l.peekChar() == '.'{
		l.readChar()
		for isDigit(l.peekChar()){
			l.readChar()
		}
		return newTokenFromSrc(tk.PRIMITIVE, tk.DEC, &l.input, position, l.position+1)
	}else{
		return newTokenFromSrc(tk.PRIMITIVE, tk.SINT, &l.input, position, l.position+1)
	}
}

func (l *Lexer) readCommentLine() string{
	l.readChar()
	l.readChar()
	position := l.position
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
	if l.ch != 0 {
		l.Depth--
	}
	return l.input[position:l.position]
}

func (l *Lexer) readCommentMultiline() string{
	l.readChar()
	l.readChar()
	position := l.position
	for !((l.ch == '*' && l.peekChar() == '/') || l.ch == 0) {
		l.readChar()
	}
	if l.ch != 0 {
		l.readChar()
		l.Depth--
	}
	return l.input[position:l.position-1]
}

func (l *Lexer) readString() string{
	ch := l.ch
	position := l.position
	another := true
	l.readChar()
	for another {
		for l.ch != ch {
			l.readChar()
		}
		if l.ch == ch && l.peekChar() != ch {
			another = false
		}
	}
	l.Depth--
	return l.input[position:l.position+1]
}

func (l *Lexer) readBytes() string{
	position := l.position
	l.readChar()
	for l.ch != '\\' {
		l.readChar()
	}
	l.Depth--
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