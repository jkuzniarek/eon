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
	msg := fmt.Sprintf("could not parse %d in %s @ line %d, col %d", p.curToken.Type, location, p.l.GetRow(), p.l.GetCol())
	p.errors = append(p.errors, msg)
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

func isHexChar(ch byte) bool{
	if 48 <= ch && ch <= 57 {
		return true
	} else if 65 <= ch && ch <= 70 {
		return true
	}
	return false
}

func isIntChar(ch byte) bool{
	if 48 <= ch && ch <= 57 {
		return true
	}
	return false
}

func isBinChar(ch byte) bool{
	if 48 <= ch && ch <= 49 {
		return true
	}
	return false
}

func hexToByte(src string) byte {
	i := 0
	out := make([]byte, 2)
	for i < 2 {
		switch src[i] {
		case 48:
			out[i] = 0
		case 49:
			out[i] = 1
		case 50:
			out[i] = 2
		case 51:
			out[i] = 3
		case 52:
			out[i] = 4
		case 53:
			out[i] = 5
		case 54:
			out[i] = 6
		case 55:
			out[i] = 7
		case 56:
			out[i] = 8
		case 57:
			out[i] = 9
		case 65:
			out[i] = 10
		case 66:
			out[i] = 11
		case 67:
			out[i] = 12
		case 68:
			out[i] = 13
		case 69:
			out[i] = 14
		case 70:
			out[i] = 15
		}
		i += 1
	}
	return (src[0]*16)+src[1]
}

func decToByte(src string) byte {
	i := 0
	out := make([]byte, 3)
	for i < 3 {
		switch src[i] {
		case 48:
			out[i] = 0
		case 49:
			out[i] = 1
		case 50:
			out[i] = 2
		case 51:
			out[i] = 3
		case 52:
			out[i] = 4
		case 53:
			out[i] = 5
		case 54:
			out[i] = 6
		case 55:
			out[i] = 7
		case 56:
			out[i] = 8
		case 57:
			out[i] = 9
		}
		i += 1
	}
	return (src[0]*100)+(src[1]*10)+src[2]
}

func binToByte(src string) byte {
	i := 0
	out := make([]byte, 8)
	for i < 8 {
		switch src[i] {
		case 48:
			out[i] = 0
		case 49:
			out[i] = 1
		}
		i += 1
	}
	return (src[0]*128)+(src[1]*64)+(src[2]*32)+(src[3]*16)+(src[4]*8)+(src[5]*4)+(src[6]*2)+src[7]
}