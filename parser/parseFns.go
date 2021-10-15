package parser

import (
	"eon/ast"
	tk "eon/token"
	"fmt"
	sc "strconv"
)


func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Expressions = []ast.Expression{}

	for !p.curTokenIs(tk.EOF) {
		expr := p.parseExpression(LOWEST)
		if expr != nil {
			program.Expressions = append(program.Expressions, expr)
		}
		p.nextToken()
	}

	return program 
}

func (p *Parser) ParseShell() *ast.Program {
	program := &ast.Program{}
	program.Expressions = []ast.Expression{}

	for !p.curTokenIs(tk.EOF) {
		expr := p.parseGroup()
		if expr != nil {
			program.Expressions = append(program.Expressions, expr)
		}
		p.nextToken()
	}

	return program 
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	p.addTrace("parseExpression("+sc.Itoa(precedence)+")'"+p.curToken.Literal+"'")
	var leftExp ast.Expression
	switch p.curToken.Cat {
	case tk.NAME:
		leftExp = p.parseName()
	case tk.OPEN_DELIMITER:
		if p.curToken.Type == tk.LPAREN {
			p.nextToken()
			if p.closeCard {
				p.closeCard = false
				leftExp = p.parseExpression(LOWEST)
				p.closeCard = true
				if !p.expectPeek(tk.RPAREN) {
					p.parsingErrAt("parseExpression() 1")
					return nil
				}	
			}else{
				leftExp = p.parseExpression(LOWEST)
				if !p.expectPeek(tk.RPAREN) {
					p.parsingErrAt("parseExpression() 1")
					return nil
				}	
			}
		} else {
			leftExp = p.parseGroup()
		}
	case tk.EVAL_OPERATOR:
		if p.curToken.Type == tk.LT {
			leftExp = p.parseCard()
		} else {
			p.parsingErrAt("parseExpression() 2")
			return nil
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
				return nil
		}
	case tk.CLOSE_DELIMITER, tk.EOF:
		msg := fmt.Sprintf("unexpected end of expression @ line %d", p.l.GetRow())
		p.errors = append(p.errors, msg)
		return nil
	default:
		msg := fmt.Sprintf("unexpected expression @ line %d", p.l.GetRow())
		p.errors = append(p.errors, msg)
		return nil
	}

	// infix handling
	for !p.peekTokenIs(tk.EOL) && !p.peekTokenIs(tk.EOF) && !(p.closeCard && p.peekTokenIs(tk.GT)) && precedence < p.peekPrecedence() {
		if _, ok := precedences[p.peekToken.Type]; !ok {
			return leftExp
		}

		p.nextToken()
		leftExp = p.parseInfix(leftExp)
	}
	// input handling
	if !p.peekTokenIs(tk.EOL) && !p.peekTokenIs(tk.EOF) && p.peekToken.Cat != tk.CLOSE_DELIMITER && !(p.closeCard && p.peekTokenIs(tk.GT)) {
		leftExp = p.parseInput(leftExp)
	}
	return leftExp
}

func (p *Parser) parseName() *ast.Name {
	p.addTrace("parseName()'"+p.curToken.Literal+"'")
	return &ast.Name{Token: p.curToken, Value: p.curToken.Literal}
}


func (p *Parser) parseCard() *ast.Card {
	p.addTrace("parseCard()")
	p.closeCard = true
	card := &ast.Card{Token: p.curToken}

	if !p.peekTokenIs(tk.GT) {
		card.Index = []ast.Expression{}
		p.nextToken()
		// parse type
		if p.curTokenIs(tk.TYPE){
			card.Type = p.curToken.Literal
			p.nextToken()
			if p.curTokenIs(tk.BSLASH) {
				p.nextToken()
				card.Size = p.parseExpression(LOWEST)
			}
		}

		var exp ast.Expression
		// parse index
		for !p.curTokenIs(tk.GT) && !p.curTokenIs(tk.SLASH) && p.curToken.Cat == tk.NAME {
			exp = p.parseExpression(LOWEST)
			card.Index = append(card.Index, exp)
			p.nextToken()
		}

		// parse body
		if p.curTokenIs(tk.SLASH) {
			p.nextToken()
			card.Body = p.parseExpression(LOWEST)
		}
	}
	p.closeCard = false
	return card
}

func (p *Parser) parseGroup() *ast.Group {
	p.addTrace("parseGroup()'"+p.curToken.Literal+"'")
	group := &ast.Group{Token: p.curToken}
	group.Expressions = []ast.Expression{}
	gType := p.curToken.Type
	var exp ast.Expression
	var endTok tk.TokenType
	p.nextToken()
	
	if gType == tk.HPAREN || gType == tk.CPAREN {
		// loop to eval expressions until RPAREN
		for !p.curTokenIs(tk.RPAREN) {
			// retain comments here only. they will be stripped during conversion to a function
			if p.curToken.Type == tk.EOL {
				p.nextToken()
				continue
			} else if p.curToken.Type == tk.COMMENT {
				exp = p.parseComment()
			} else {
				exp = p.parseExpression(LOWEST)
			}
			// append exp to expression list
			if exp != nil {
				group.Expressions = append(group.Expressions, exp)
			}
			if p.peekToken.Type != tk.RPAREN && p.peekToken.Type != tk.EOL {
				p.errors = append(p.errors, "unexpected end of group")
				return nil
			} else {
				p.nextToken()
			}
		}
	} else {
		switch gType {
		case tk.LSQUAR:
			endTok = tk.RSQUAR
		case tk.LCURLY, tk.SCURLY:
			endTok = tk.RCURLY
		default:
			p.parsingErrAt("parseGroup()")
			return nil
		}
		// loop to eval expressions until group close delimiter
		for !p.curTokenIs(endTok) {
			exp = p.parseExpression(LOWEST)
			// append exp to expression list
			if(exp != nil){
				group.Expressions = append(group.Expressions, exp)
			}
			p.nextToken()
		}
	}
	return group
}

func (p *Parser) parseInfix(left ast.Expression) ast.Expression {
	p.addTrace("parseInfix()'"+p.curToken.Literal+"'")
	expression := &ast.Infix{
		Token: p.curToken,
		Operator: p.curToken.Literal,
		Left: left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)
	return expression
}

func (p *Parser) parseInput(left ast.Expression) ast.Expression {
	p.addTrace("parseInput()'"+p.curToken.Literal+"'")
	expression := &ast.Input{
		Left: left,
	}
	p.nextToken()
	expression.Input = p.parseExpression(LOWEST)
	return expression
}