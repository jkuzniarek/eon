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
	PIPESRC // 4 | function
	EQUALS // == or !=
	LESSGREATER // < or >
	CALL // myFunction x
)

var precedences = map[tk.TokenType]int{
	tk.PIPE: PIPESRC,
	tk.TYPE_EQ: EQUALS, 
	tk.EQEQ: EQUALS,
	tk.NOT_EQ: EQUALS,
	tk.LT: LESSGREATER,
	tk.GT: LESSGREATER,
	tk.LT_EQ: LESSGREATER,
	tk.GT_EQ: LESSGREATER,
}


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

