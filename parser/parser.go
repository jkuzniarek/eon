package parser 

import (
	"eon/lexer"
	tk "eon/token"
	"fmt"
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

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t tk.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t.ToStr(), p.peekToken.Type.ToStr())
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	if(p.curTokenIs(tk.EOF)){
		return
	}
	p.peekToken = p.l.NextToken()
}

func (p *Parser) curTokenIs(t tk.TokenType) bool {
	return p.curToken.Type == t 
}

func (p *Parser) peekTokenIs(t tk.TokenType) bool {
	return p.peekToken.Type == t 
}

func (p *Parser) expectPeek(t tk.TokenType) bool{
	if p.peekTokenIs(t){
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) parsingErrAt(location string) {
	msg := fmt.Sprintf("could not parse %s %s in %s @ line %d", p.curToken.Cat.ToStr(), p.curToken.Type.ToStr(), location, p.l.GetRow())
	p.errors = append(p.errors, msg)
}

func (p *Parser) peekPrecedence() int {
	return getPrecedence(p.peekToken.Type)
}

func (p *Parser) curPrecedence() int {
	return getPrecedence(p.curToken.Type)
}