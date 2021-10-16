package parser 

import (
	"eon/lexer"
	tk "eon/token"
)

const (
	_ int = iota
	LOWEST
	PIPESRC // 4 | function
	EQUALS // == or !=
	LESSGREATER // < or >
	ASSIGN // var: 1
	CALL // myFunction x
)

var precedences = map[tk.TokenType]int{
	tk.PIPE: PIPESRC,
	tk.TYPE_EQ: EQUALS, 
	tk.EQEQ: EQUALS,
	tk.NOT_EQ: EQUALS,
	tk.LT: LESSGREATER,
	tk.GT: LESSGREATER,
	tk.LT_EQ: LESSGREATER,
	tk.GT_EQ: LESSGREATER,
	tk.SET_VAL: ASSIGN,
	tk.SET_CONST: ASSIGN,
	tk.SET_WEAK: ASSIGN,
	tk.SET_BIND: ASSIGN,
	tk.SET_PLUS: ASSIGN,
	tk.SET_MINUS: ASSIGN,
	tk.SET_TYPE: ASSIGN,
	// tk.NAME: CALL, // TODO: confirm that this is handled by the parse functions correctly, expect that it will be handled automatically
	tk.DOT: CALL,
	tk.SLASH: CALL,
	tk.OCTO: CALL,
	tk.STAR: CALL,
	tk.AT: CALL,
}


type Parser struct {
	l *lexer.Lexer 
	errors []string

	curToken tk.Token
	peekToken tk.Token

	shellEnv bool
	Trace string
	inCard bool
}

func New(l *lexer.Lexer, sh bool) *Parser {
	p := &Parser{
		l: l,
		shellEnv: sh,
		errors: []string{},
		Trace: "",
		inCard: false,
	}


	
	// read 2 tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p 
}

func (p *Parser) addTrace(s string) {
	p.Trace += "\n"+s
}