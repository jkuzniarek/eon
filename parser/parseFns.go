package parser

import (
	"eon/ast"
	"eon/lexer"
	tk "eon/token"
	"strconv"
	"fmt"
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

func (p *Parser) parseExpression(precedence int) ast.Expression {
	accessorStart := false
	switch p.curToken.Cat {
	case tk.NAME:
		leftExpr := p.parseName()
	case tk.OPEN_DELIMITER:
		if p.curToken.Type == tk.LPAREN {
			p.nextToken()
			leftExpr := p.parseExpression(LOWEST)
			if !p.expectPeek(tk.RPAREN) {
				p.parsingErrAt("parseExpression()")
				return nil
			}	
		} else {
			leftExp := p.parseGroup()
		}
	case tk.EVAL_OPERATOR:
		if p.curToken.Type == tk.LT {
			leftExp := p.parseCard()
		} else {
			p.parsingErrAt("parseExpression()")
			return nil
		}
	case tk.PRIMITIVE:
		switch p.curToken.Type {
			case SINT: 
				leftExp := parseSInt()
			case UINT:
				leftExp := parseUInt()
			case DEC:
				leftExp := parseDec()
			case STR:
				leftExp := parseStr()
			case BYTES:
				leftExp := parseBytes()
			default:
				p.parsingErrAt("parseExpression()")
				return nil
		}
	default:
		p.parsingErrAt("parseExpression()")
		return nil
	}

	// infix handling
	for !p.peekTokenIs(tk.EOL) && precedence < p.peekPrecedence() {
		if p, ok := precedences[p.peekToken.Type]; !ok {
			return leftExp
		}

		p.nextToken()
		leftExp = p.parseInfix(leftExp)
	}
	// input handling
	if !p.peekTokenIs(tk.EOL) {
		leftExp = p.parseInput(leftExp)
	}
	
	if !p.peekTokenIs(tk.EOL) {
		p.parsingErrAt("parseExpression()")
		return nil
	}
	return leftExp


func (p *Parser) parseName() ast.Expression {
	return &ast.Name{Token: p.curToken, Value: p.curToken.Literal}
}


func (p *Parser) parseCard() ast.Expression {
	card := &ast.Card{Token: p.curToken}

	if !p.peekTokenIs(tk.GT) {
		card.Index = []ast.Expression{}
		if p.peekTokenIs(tk.TYPE){
			p.nextToken()
			card.Type = p.curToken
			if p.peekTokenIs(tk.BSLASH) {
				p.nextToken()
				p.nextToken()
				card.Size = p.parseExpression()
			}
		}

		for !p.peekTokenIs(tk.GT){
			p.nextToken()
			if p.curToken.Cat == tk.NAME {
				exp := p.parseExpression()
				card.Index = append(card.Index, exp)
				p.nextToken()
			} else if p.curTokenIs(tk.SLASH){
				p.nextToken()
				card.Body = p.parseExpression()
			}
		}
		p.nextToken()
	}

	return card
}

func (p *Parser) parseGroup() ast.Expression {
	group := &ast.Group{Token: p.curToken}
	group.Expressions = []ast.Expression{}
	gType := p.curToken.TokenType
	p.nextToken()
	
	if gType == tk.HPAREN || gType == tk.CPAREN {
		// loop to eval expressions until RPAREN
		for !p.curTokenIs(tk.RPAREN) {
			// retain comments here only. they will be stripped during conversion to a function
			if p.curToken.TokenType == tk.EOL {
				p.nextToken()
				continue
			} else if p.curToken.TokenType == tk.COMMENT {
				exp := p.parseComment()
			} else {
				exp := p.parseExpression()
			}
			// append exp to expression list
			if(exp != nil){
				group.Expressions = append(group.Expressions, exp)
			}
			if p.peekToken.TokenType != tk.RPAREN && p.peekToken.TokenType != tk.EOL {
				msg := fmt.Sprintf("expected next token to be EOL or RPAREN, got %s instead", p.peekToken.Type.ToStr())
				p.errors = append(p.errors, msg)
				return nil
			} else {
				p.nextToken()
			}
		}
		return group
		
	} else {
		switch p.curToken.TokenType {
		case tk.LSQUAR:
			endTok := tk.RSQUAR
		case tk.LCURLY, tk.SCURLY:
			endTok := tk.RCURLY
		default:
			p.parsingErrAt("parseExpression()")
			return nil
		}
		// loop to eval expressions until group close delimiter
		for !p.curTokenIs(endTok) {
			exp := p.parseExpression()
			// append exp to expression list
			if(exp != nil){
				group.Expressions = append(group.Expressions, exp)
			}
			p.nextToken()
			return group
	}
}

func (p *Parser) parseInfix(left ast.Expression) ast.Expression {
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
	expression := &ast.Input{
		Left: left,
	}
	p.nextToken()
	expression.Input = p.parseExpression(LOWEST)
	return expression
}