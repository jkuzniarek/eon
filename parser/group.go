package parser

import (
	"eon/ast"
	tk "eon/token"
)

func (p *Parser) parseGroup() *ast.Group {
	p.addTrace("START parseGroup()'"+p.curToken.Literal+"'")
	p.Depth++
	group := &ast.Group{Token: p.curToken}
	group.Expressions = []ast.Node{}
	gType := p.curToken.Type
	var exp ast.Node
	var endTok tk.TokenType
	
	if gType == tk.HPAREN || gType == tk.CPAREN {
		// loop to eval expressions until RPAREN
		for !p.peekTokenIs(tk.RPAREN) {
			p.nextToken()
			// retain comments here only. they will be stripped during conversion to a function
			if p.curToken.Type == tk.EOL {
				continue
			}
			if p.curToken.Type == tk.COMMENT {
				exp = p.parseComment()
			} else {
				exp = p.parseExpression(LOWEST)
			}
			// append exp to expression list
			if exp != nil {
				group.Expressions = append(group.Expressions, exp)
			}
			if !p.peekTokenIs(tk.RPAREN) && !p.peekTokenIs(tk.EOL) {
				p.errors = append(p.errors, "unexpected end of group")
				p.addTrace("END parse return nil")
				return nil
			}
		}
	} else {
		switch gType {
		case tk.LSQUAR:
			endTok = tk.RSQUAR
		case tk.SCURLY:
			endTok = tk.RCURLY
		default:
			p.parsingErrAt("parseGroup()")
			p.addTrace("END parse return nil")
			return nil
		}
		// loop to eval expressions until group close delimiter
		for !p.peekTokenIs(endTok) && !p.peekTokenIs(tk.EOF) {
			p.nextToken()
			if p.curTokenIs(tk.EOL){
				continue
			}
			exp = p.parseExpression(LOWEST)
			// append exp to expression list
			if exp != nil {
				group.Expressions = append(group.Expressions, exp)
			}
		}
		if !p.peekTokenIs(endTok){
			p.errors = append(p.errors, "unexpected end of group")
			p.addTrace("END parse return nil")
			return nil
		}
	}
	p.nextToken()
	p.Depth--
	p.addTrace("END parseGroup()")
	return group
}