package parser

import (
	"eon/ast"
)

func (p *Parser) parseInfix(left ast.Node) ast.Node {
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