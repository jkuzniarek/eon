package parser

import (
	"eon/ast"
	"fmt"
	ssDec "github.com/shopspring/decimal"
)

func (p *Parser) parseDec() ast.Node {
	lit := &ast.Dec{Token: p.curToken}
	value, err := ssDec.NewFromString(p.curToken.Literal)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as decimal", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = value
	return lit
}