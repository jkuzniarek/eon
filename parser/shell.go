package parser

import (
	"eon/ast"
	tk "eon/token"
)


func (p *Parser) ParseShell() *ast.Program {
	program := &ast.Program{}
	program.Expressions = []ast.Node{}

	for !p.curTokenIs(tk.EOF) {
		if p.curTokenIs(tk.EOL){
			p.nextToken()
			continue
		}
		expr := p.parseExpression(LOWEST)
		if expr != nil {
			program.Expressions = append(program.Expressions, expr)
		}
		p.nextToken()
	}

	return program 
}