package parser

import (
	"eon/token"
	"fmt"
)


func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t.ToStr(), p.peekToken.Type.ToStr())
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	if(p.curTokenIs(token.EOF)){
		return
	}
	p.peekToken = p.l.NextToken()
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t 
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t 
}

func (p *Parser) expectPeek(t token.TokenType) bool{
	if p.peekTokenIs(t){
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}