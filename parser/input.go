package parser

import (
	"eon/ast"
)

func (p *Parser) parseInput(left ast.Node) ast.Node {
	p.addTrace("parseInput()'"+p.curToken.Literal+"'")
	expression := &ast.Input{
		Left: left,
	}
	// p.nextToken()
	expression.Input = p.parseExpression(LOWEST)
	p.addTrace("END parseInput()")
	return expression
}