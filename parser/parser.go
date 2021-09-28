package parser 

import (
	"eon/ast"
	"eon/lexer"
	tk "eon/token"
	"strconv"
	"fmt"
)

const (
	_ int = iota
	LOWEST
	EQUALS // == or !=
	LESSGREATER // < or >
	CALL // myFunction(x)
)

var precedences = map[tk.TokenType]int{
}

type (
	infixParseFn func(ast.Expression) ast.Expression
)

type Parser struct {
	l *lexer.Lexer 
	errors []string

	curToken tk.Token
	peekToken tk.Token

	shellEnv bool

	infixParseFns map[tk.TokenType]infixParseFn
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


func (p *Parser) registerInfix(tokenType tk.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}


func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Commands = []ast.Expression{}

	for !p.curTokenIs(tk.EOF) {
		expr := p.parseExpression(LOWEST)
		if expr != nil {
			program.Commands = append(program.Commands, expr)
		}
		p.nextToken()
	}

	return program 
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	switch p.curToken.Cat {
	case tk.NAME, tk.KEYWORD:
		return p.parseAccessor()
	case tk.OPEN_DELIMITER:
		return p.parseGroup()
	case tk.EVAL_OPERATOR:
		if p.curToken.Type == tk.LT {
			return p.parseCard()
		} else {
			p.parsingErrAt("parseExpression()")
			return nil
		}
	case tk.PRIMITIVE:
		switch p.curToken.Type {
			case SINT: 
				return parseSInt()
			case UINT:
				return parseUInt()
			case SDEC:
				return parseSDec()
			case UDEC:
				return parseUDec()
			case STR:
				return parseStr()
			case BYTES:
				return parseBytes()
			default:
				p.parsingErrAt("parseExpression()")
				return nil
		}
	default:
		p.parsingErrAt("parseExpression()")
		return nil
	}


	switch p.curToken.Type {
	case tk.NAME, tk.INIT, tk.DEST, tk.OUT:
		switch p.peekToken.Type {
		case tk.SET_VAL, tk.SET_CONST, tk.SET_WEAK, tk.SET_BIND, tk.SET_PLUS, tk.SET_MINUS, tk.SET_TYPE:
			return p.parseAssignExpression()
		case tk.DOT, tk.SLASH, tk.OCTO, tk.STAR, tk.AT, tk.PIPE, tk.BANG, tk.DOLLAR, tk.PERCENT, tk.CARET,
		tk.TYPE_EQ, tk.EQEQ, tk.NOT_EQ, tk.LT, tk.GT, tk.LT_EQ, tk.GT_EQ:
			return p.parseInfixExpression()
		}
	}
	
	return leftExp
}

func (p *Parser) parseAccessor() *ast.Accessor {

}

func (p *Parser) parseCard() *ast.Group {

}

func (p *Parser) parseSInt() *ast.Group {

}

func (p *Parser) parseUInt() *ast.Group {

}

func (p *Parser) parseSDec() *ast.Group {

}

func (p *Parser) parseUDec() *ast.Group {

}

func (p *Parser) parseStr() *ast.Group {

}

func (p *Parser) parseBytes() *ast.Group {

}


// FOR REFERENCE
func (p *Parser) parseAssignExpression() *ast.AssignmentExpr {
	expr := &ast.AssignmentExpr{Name: &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}}

	if !(
		p.expectPeek(tk.SET_VAL) || 
		p.expectPeek(tk.SET_CONST) || 
		p.expectPeek(tk.SET_WEAK) || 
		p.expectPeek(tk.SET_BIND) || 
		p.expectPeek(tk.SET_PLUS) || 
		p.expectPeek(tk.SET_MINUS) || 
		p.expectPeek(tk.SET_TYPE)){
		return nil
	}

	expr.Token = p.curToken

	// for now, skip past p.parseExpression() with this for loop
	for !p.curTokenIs(tk.EOL){
		p.nextToken()
	}

	return expr 
	
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseUIntegerLiteral() ast.Expression {
	lit := &ast.UIntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseUint(p.curToken.Literal, 0, 0)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as unsigned integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = uint(value)

	return lit
}

func (p *Parser) parseSIntegerLiteral() ast.Expression {
	lit := &ast.SIntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 0)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as signed integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = int(value)

	return lit
}












