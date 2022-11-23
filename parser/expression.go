package parser

import (
	"eon/ast"
	tk "eon/token"
	"fmt"
	sc "strconv"
)

func (p *Parser) parseExpression(precedence int) ast.Node {
	p.addTrace("START parseExpression("+sc.Itoa(precedence)+")'"+p.curToken.Literal+"'")
	for p.curTokenIs(tk.EOL){
		p.nextToken()
	}
	var leftExp ast.Node
	switch p.curToken.Cat {
	case tk.NAME:
		leftExp = p.parseName()
	case tk.OPEN_DELIMITER:
		if p.curToken.Type == tk.LPAREN {
			p.Depth++
			p.nextToken()
			if p.inCard {
				p.inCard = false
				leftExp = p.parseExpression(LOWEST)
				p.inCard = true
				if !p.expectPeek(tk.RPAREN) {
					p.parsingErrAt("parseExpression() 1")
					p.addTrace("END parse return nil")
					return nil
				}	
			}else{
				leftExp = p.parseExpression(LOWEST)
				if !p.expectPeek(tk.RPAREN) {
					p.parsingErrAt("parseExpression() 1")
					p.addTrace("END parse return nil")
					return nil
				}	
			}
			p.Depth--
		} else if p.curToken.Type == tk.LCURLY {
			leftExp = p.parseCard()
		}else {
			leftExp = p.parseGroup()
		}
	case tk.PRIMITIVE:
		switch p.curToken.Type {
			case tk.SINT: 
				leftExp = p.parseSInt()
			case tk.UINT:
				leftExp = p.parseUInt()
			case tk.DEC:
				leftExp = p.parseDec()
			case tk.STR:
				leftExp = p.parseStr()
			case tk.BYTES:
				leftExp = p.parseBytes()
			default:
				p.parsingErrAt("parseExpression() 3")
				p.addTrace("END parse return nil")
				return nil
		}
	case tk.CLOSE_DELIMITER, tk.EOF:
		msg := fmt.Sprintf("unexpected end of expression @ line %d", p.l.GetRow())
		p.errors = append(p.errors, msg)
		p.addTrace("END parse return nil")
		return nil
	default:
		msg := fmt.Sprintf("unexpected expression @ line %d", p.l.GetRow())
		p.errors = append(p.errors, msg)
		p.addTrace("END parse return nil")
		return nil
	}

	// infix handling
	for !p.peekTokenIs(tk.EOL) && !p.peekTokenIs(tk.EOF) && !(p.inCard && p.peekToken.Literal == "/") && precedence < p.peekPrecedence() {
		if p.peekToken.Cat != tk.OPERATOR {
			p.addTrace("END parseExpression()")
			return leftExp
		}

		p.nextToken()
		leftExp = p.parseInfix(leftExp)
	}
	// input handling
	if !p.peekTokenIs(tk.EOL) && !p.peekTokenIs(tk.EOF) && p.peekToken.Cat != tk.CLOSE_DELIMITER && !(p.inCard && p.peekToken.Literal == "/") {
		p.nextToken()
		leftExp = p.parseInput(leftExp)
	}
	p.addTrace("END parseExpression()")
	return leftExp
}