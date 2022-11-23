package parser

import (
	"eon/ast"
	tk "eon/token"
)

func (p *Parser) parseCard() *ast.Card {
	// TODO: change to parse body as part of index with key '/'
	p.addTrace("START parseCard()")
	
	card := &ast.Card{Token: p.curToken}

	if p.peekTokenIs(tk.RCURLY) {
		p.nextToken()
	}else {
		p.inCard = true
		p.Depth++
		card.Index = []ast.Node{}
		p.nextToken()

		var exp ast.Node
		// parse index
		for !p.curTokenIs(tk.RCURLY) && p.curToken.Cat == tk.NAME && p.curToken.Literal != "/" {
			exp = p.parseExpression(LOWEST)
			card.Index = append(card.Index, exp)
			p.addTrace("__index expression last token: "+p.curToken.Literal)
			p.nextToken()
			p.addTrace("__index expression next token: "+p.curToken.Literal)
		}

		// parse body
		if p.curToken.Literal == "/" {
			p.nextToken()
			card.Body = p.parseExpression(LOWEST)
			p.nextToken()
		}
		
		// end card
		for p.curTokenIs(tk.EOL){
			p.nextToken()
		}
		if !p.curTokenIs(tk.RCURLY) {
			p.parsingErrAt("parseCard()")
			p.addTrace("END parse return nil")
			return nil
		}	
		p.Depth--
		p.inCard = false
	}
	
	p.addTrace("END parseCard()")
	return card
}