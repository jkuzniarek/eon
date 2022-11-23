package parser

import (
	"eon/ast"
)

func (p *Parser) parseStr() ast.Node {
	lit := &ast.Str{Token: p.curToken}
	src := p.curToken.Literal
	srcLen := len(src)
	ch := src[0]
	position := 1
	strStart := 1
	strEnd := 1
	another := true
	value := ""
	for another {
		strStart = position
		strEnd = position
		for src[position] != ch {
			position++
			strEnd++
		}
		value += src[strStart:strEnd]
		if position == (srcLen-1) {
			another = false
		} else if position < (srcLen-1) && src[position+1] == ch {
			value += string(ch)
			position = position + 2
		}
	}
	lit.Value = value
	return lit
}