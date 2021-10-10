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
	program.Commands = []ast.Expression{}

	for !p.curTokenIs(tk.EOF) {
		expr := p.parseExpression(LOWEST)
		if expr != nil {
			program.Commands = append(program.Commands, expr)
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
		leftExp := p.parseGroup()
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
	return leftExp

	// // original
	// switch p.curToken.Type {
	// case tk.NAME, tk.INIT, tk.DEST, tk.OUT:
	// 	switch p.peekToken.Type {
	// 	case tk.SET_VAL, tk.SET_CONST, tk.SET_WEAK, tk.SET_BIND, tk.SET_PLUS, tk.SET_MINUS, tk.SET_TYPE:
	// 		return p.parseAssignExpression()
	// 	case tk.DOT, tk.SLASH, tk.OCTO, tk.STAR, tk.AT, tk.PIPE, tk.BANG, tk.DOLLAR, tk.PERCENT, tk.CARET,
	// 	tk.TYPE_EQ, tk.EQEQ, tk.NOT_EQ, tk.LT, tk.GT, tk.LT_EQ, tk.GT_EQ:
	// 		return p.parseInfixExpression()
	// 	}
	// }
	// return leftExp
}

func (p *Parser) parseName() ast.Expression {
	return &ast.Name{Token: p.curToken, Value: p.curToken.Literal}
}

p.nextToken()
expression.Right = p.parseExpression(LOWEST)
return expression
}

func (p *Parser) parseCard() ast.Expression {
// TODO
}

func (p *Parser) parseGroup() ast.Expression {
	group := &ast.Group{Token: p.curToken}
	group.Expressions = []ast.Expression{}
	gType := p.curToken.TokenType
	p.nextToken()
	
	if gType == tk.LPAREN {
		exp := p.parseExpression(LOWEST)
		if !p.expectPeek(tk.RPAREN) {
			return nil
		}
		return exp

	} else if gType == tk.HPAREN || gType == tk.CPAREN {
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

// FOR REFERENCE
// func (p *Parser) parseAssignExpression() *ast.AssignmentExpr {
// 	expr := &ast.AssignmentExpr{Name: &ast.Name{Token: p.curToken, Value: p.curToken.Literal}}

// 	if !( 
// 		p.peekTokenIs(tk.SET_VAL) || 
// 		p.peekTokenIs(tk.SET_CONST) || 
// 		p.peekTokenIs(tk.SET_WEAK) || 
// 		p.peekTokenIs(tk.SET_BIND) || 
// 		p.peekTokenIs(tk.SET_PLUS) || 
// 		p.peekTokenIs(tk.SET_MINUS) || 
// 		p.peekTokenIs(tk.SET_TYPE)
// 		){
// 		return nil
// 	}

// 	p.nextToken()
// 	expr.Token = p.curToken
// 	p.nextToken()

// 	expr.Value = p.parseExpression(LOWEST)

// 	return expr 
	
// }
