package parser

import (
	"eon/ast"
	"strconv"
	"fmt"
)


func (p *Parser) parseSInt() ast.Node {
	lit := &ast.SInt{Token: p.curToken}
	value, err := strconv.ParseInt(p.curToken.Literal, 10, 0)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as signed integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = int(value)
	return lit
}