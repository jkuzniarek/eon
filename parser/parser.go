package parser 

import (
	"eon/lexer"
	tk "eon/token"
)

const (
	_ int = iota
	LOWEST
	EQUALS // EVAL_OPERATOR (including PIPE)
	// LESSGREATER // < or >
	ASSIGN // ASSIGN_OPERATOR
	CALL // myFunction x and ACCESS_OPERATOR
)


type Parser struct {
	l *lexer.Lexer 
	errors []string

	curToken tk.Token
	peekToken tk.Token

	shellEnv bool
	Trace string
	inCard bool
	Depth int
}

func New(l *lexer.Lexer, sh bool) *Parser {
	p := &Parser{
		l: l,
		shellEnv: sh,
		errors: []string{},
		Trace: "",
		inCard: false,
		Depth: 0,
	}


	
	// read 2 tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p 
}

func (p *Parser) addTrace(s string) {
	p.Trace += "\n"+s
}