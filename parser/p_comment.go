package parser

import (
	"eon/ast"
)

func (p *Parser) parseComment() ast.Node {
	lit := &ast.Comment{Token: p.curToken}
	src := p.curToken.Literal
	srcLen := len(src)
	if src[1] == '*' {
		lit.Multiline = true
	}else{
		lit.Multiline = false
	}
	lit.Value = src[2:srcLen]
	return lit
}