package parser

import (
	"eon/ast"
)

func (p *Parser) parseName() *ast.Name {
	p.addTrace("parseName()'"+p.curToken.Literal+"'")
	return &ast.Name{Token: p.curToken, Value: p.curToken.Literal}
}