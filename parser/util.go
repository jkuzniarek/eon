package parser

import (
	tk "eon/token"
	"fmt"
)


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
	msg := fmt.Sprintf("could not parse %q in %s", p.curToken.Literal, location)
	p.errors = append(p.errors, msg)
}

func (p *Parser) registerInfix(tokenType tk.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p 
	}
	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p 
	}
	return LOWEST
}