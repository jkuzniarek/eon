package parser

import (
	"eon/ast"
	"strconv"
	"fmt"
)

func (p *Parser) parseUInt() ast.Node {
	lit := &ast.UInt{Token: p.curToken}
	value, err := strconv.ParseUint(p.curToken.Literal, 10, 0)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as unsigned integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = uint(value)
	return lit
}