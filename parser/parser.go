package parser 

import (
	"eon/ast"
	"eon/lexer"
	"eon/token"
)

const (
	_ int = iota
	
)

var precedences = map[token.TokenType]int{
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn func(ast.Expression) ast.Expression
)

type Parser struct {
	l *lexer.Lexer 

	curToken token.Token
	peekToken token.Token

	shellEnv bool

	errors []string
}

func New(l *lexer.Lexer, sh bool) *Parser {
	p := &Parser{
		l: l,
		shellEnv: sh,
		errors: []string{},
	}

	// read 2 tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p 
}


func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Expression{}

	for !p.curTokenIs(token.EOF) {
		expr := p.parseExpression()
		if expr != nil {
			program.Statements = append(program.Statements, expr)
		}
		p.nextToken()
	}

	return program 
}

func (p *Parser) parseExpression() ast.Expression {
	switch p.curToken.Type {
	case token.NAME, token.INIT, token.DEST, token.OUT:
		switch p.peekToken.Type {
		case token.SET_VAL, token.SET_CONST, token.SET_WEAK, token.SET_BIND:
			return p.parseAssignExpression()
		default:
			return nil	
		}
	default:
		return nil
	}
}

func (p *Parser) parseAssignExpression() *ast.AssignmentExpr {
	expr := &ast.AssignmentExpr{Name: &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}}

	if !(p.expectPeek(token.SET_VAL) || p.expectPeek(token.SET_CONST) || p.expectPeek(token.SET_WEAK) || p.expectPeek(token.SET_BIND)){
		return nil
	}

	expr.Token = p.curToken

	// for now, skip past p.parseExpression() with this for loop
	for !p.curTokenIs(token.EOL){
		p.nextToken()
	}

	return expr 
	
	// <-- here
}
/*
// mutate for object
func (p *Parser) ParseObject(parent &Object) *object.Object {
	object := &object.Object{}
	object.Owner = parent
	
	for !p.curTokenIs(token.EOF) {
		switch p.curToken.Type {
			case token.TYPE:
				object.TypeValue = p.curToken.Value
			case token.NAME:
				p.parseIndex(object)
		}
		
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.nextToken()
	}

	return program 
} */

// for consult
// func (p *Parser) parseStatement() ast.Statement{
// 	switch p.curToken.Type {
// 	case token.LET:
// 		return p.parseLetStatement()
// 	case token.RETURN:
// 		return p.parseReturnStatement()
// 	default:
// 		return p.parseExpressionStatement()
// 	}
// }

/*
func (p *Parser) parseIndex(o &Object) {
	for !p.curTokenIs(token.EOF){
		switch tok.Type{
			
		}
	}
}
*/










