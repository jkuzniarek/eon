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

	switch p.curToken.Cat {
	case tk.NAME, tk.KEYWORD:
		leftExp := p.parseAccessor()
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
			case SDEC:
				leftExp := parseSDec()
			case UDEC:
				leftExp := parseUDec()
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

func (p *Parser) parseAccessor() *ast.Accessor {

}

func (p *Parser) parseCard() *ast.Group {

}

func (p *Parser) parseSInt() *ast.Group {

}

func (p *Parser) parseUInt() *ast.Group {

}

func (p *Parser) parseSDec() *ast.Group {

}

func (p *Parser) parseUDec() *ast.Group {

}

func (p *Parser) parseStr() *ast.Group {

}

func (p *Parser) parseBytes() *ast.Group {

}


// FOR REFERENCE
func (p *Parser) parseAssignExpression() *ast.AssignmentExpr {
	expr := &ast.AssignmentExpr{Name: &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}}

	if !(
		p.expectPeek(tk.SET_VAL) || 
		p.expectPeek(tk.SET_CONST) || 
		p.expectPeek(tk.SET_WEAK) || 
		p.expectPeek(tk.SET_BIND) || 
		p.expectPeek(tk.SET_PLUS) || 
		p.expectPeek(tk.SET_MINUS) || 
		p.expectPeek(tk.SET_TYPE)){
		return nil
	}

	expr.Token = p.curToken

	// for now, skip past p.parseExpression() with this for loop
	for !p.curTokenIs(tk.EOL){
		p.nextToken()
	}

	return expr 
	
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseUIntegerLiteral() ast.Expression {
	lit := &ast.UIntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseUint(p.curToken.Literal, 0, 0)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as unsigned integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = uint(value)

	return lit
}

func (p *Parser) parseSIntegerLiteral() ast.Expression {
	lit := &ast.SIntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 0)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as signed integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = int(value)

	return lit
}

func (p *Parser) parseInfix(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token: p.curToken,
		Operator: p.curToken.Literal,
		Left: left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)
	return expression
}