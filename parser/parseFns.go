package parser

import (
	"eon/ast"
	tk "eon/token"
	"fmt"
	sc "strconv"
)


func (p *Parser) ParseShell() *ast.Program {
	program := &ast.Program{}
	program.Expressions = []ast.Expression{}

	for !p.curTokenIs(tk.EOF) {
		if p.curTokenIs(tk.EOL){
			p.nextToken()
			continue
		}
		expr := p.parseExpression(LOWEST)
		if expr != nil {
			program.Expressions = append(program.Expressions, expr)
		}
		p.nextToken()
	}

	return program 
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	p.addTrace("START parseExpression("+sc.Itoa(precedence)+")'"+p.curToken.Literal+"'")
	for p.curTokenIs(tk.EOL){
		p.nextToken()
	}
	var leftExp ast.Expression
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

func (p *Parser) parseName() *ast.Name {
	p.addTrace("parseName()'"+p.curToken.Literal+"'")
	return &ast.Name{Token: p.curToken, Value: p.curToken.Literal}
}


func (p *Parser) parseCard() *ast.Card {
	// TODO: change to parse body as part of index with key '/'
	p.addTrace("START parseCard()")
	
	card := &ast.Card{Token: p.curToken}

	if p.peekTokenIs(tk.RCURLY) {
		p.nextToken()
	}else {
		p.inCard = true
		p.Depth++
		card.Index = []ast.Expression{}
		p.nextToken()

		var exp ast.Expression
		// parse index
		for !p.curTokenIs(tk.RCURLY) && p.curToken.Cat == tk.NAME && p.curToken.Literal != "/" {
			exp = p.parseExpression(LOWEST)
			card.Index = append(card.Index, exp)
			p.addTrace("__index expression last token: "+p.curToken.Literal)
			p.nextToken()
			p.addTrace("__index expression next token: "+p.curToken.Literal)
		}

		// parse body
		if p.curToken.Literal == "/" {
			p.nextToken()
			card.Body = p.parseExpression(LOWEST)
			p.nextToken()
		}
		
		// end card
		for p.curTokenIs(tk.EOL){
			p.nextToken()
		}
		if !p.curTokenIs(tk.RCURLY) {
			p.parsingErrAt("parseCard()")
			p.addTrace("END parse return nil")
			return nil
		}	
		p.Depth--
		p.inCard = false
	}
	
	p.addTrace("END parseCard()")
	return card
}

func (p *Parser) parseGroup() *ast.Group {
	p.addTrace("START parseGroup()'"+p.curToken.Literal+"'")
	p.Depth++
	group := &ast.Group{Token: p.curToken}
	group.Expressions = []ast.Expression{}
	gType := p.curToken.Type
	var exp ast.Expression
	var endTok tk.TokenType
	
	if gType == tk.HPAREN || gType == tk.CPAREN {
		// loop to eval expressions until RPAREN
		for !p.peekTokenIs(tk.RPAREN) {
			p.nextToken()
			// retain comments here only. they will be stripped during conversion to a function
			if p.curToken.Type == tk.EOL {
				continue
			}
			if p.curToken.Type == tk.COMMENT {
				exp = p.parseComment()
			} else {
				exp = p.parseExpression(LOWEST)
			}
			// append exp to expression list
			if exp != nil {
				group.Expressions = append(group.Expressions, exp)
			}
			if !p.peekTokenIs(tk.RPAREN) && !p.peekTokenIs(tk.EOL) {
				p.errors = append(p.errors, "unexpected end of group")
				p.addTrace("END parse return nil")
				return nil
			}
		}
	} else {
		switch gType {
		case tk.LSQUAR:
			endTok = tk.RSQUAR
		case tk.SCURLY:
			endTok = tk.RCURLY
		default:
			p.parsingErrAt("parseGroup()")
			p.addTrace("END parse return nil")
			return nil
		}
		// loop to eval expressions until group close delimiter
		for !p.peekTokenIs(endTok) && !p.peekTokenIs(tk.EOF) {
			p.nextToken()
			if p.curTokenIs(tk.EOL){
				continue
			}
			exp = p.parseExpression(LOWEST)
			// append exp to expression list
			if exp != nil {
				group.Expressions = append(group.Expressions, exp)
			}
		}
		if !p.peekTokenIs(endTok){
			p.errors = append(p.errors, "unexpected end of group")
			p.addTrace("END parse return nil")
			return nil
		}
	}
	p.nextToken()
	p.Depth--
	p.addTrace("END parseGroup()")
	return group
}

func (p *Parser) parseInfix(left ast.Expression) ast.Expression {
	p.addTrace("START parseInfix()'"+p.curToken.Literal+"'")
	expression := &ast.Infix{
		Token: p.curToken,
		Operator: p.curToken.Literal,
		Left: left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)
	p.addTrace("END parseInfix()")
	return expression
}

func (p *Parser) parseInput(left ast.Expression) ast.Expression {
	p.addTrace("parseInput()'"+p.curToken.Literal+"'")
	expression := &ast.Input{
		Left: left,
	}
	// p.nextToken()
	expression.Input = p.parseExpression(LOWEST)
	p.addTrace("END parseInput()")
	return expression
}