package token

type TokenType int

type Token struct{
	Cat TokenType
	Type TokenType
	Literal string
}

const (
	_ = iota // makes the constants auto incremented integers
	ILLEGAL
	EOF
	EOL
	BSLASH
	
	// Identifiers & Literals
	KEYWORD 
	NAME // add, foobar, x, y, ...
	OPERATOR
	OPEN_DELIMITER
	CLOSE_DELIMITER
	PRIMITIVE
	UINT // 12345
	SINT // +1, -2, +0
	DEC // 1.2, +1.0, -2.00, +0.000
	BYTES // \x 1D 2A FF\, \d 255 23 0\, \b 00110111 11101000 00000000\
	STR 
	COMMENT 
	
	// Delimiters
	HPAREN
	CPAREN 
	LPAREN 
	RPAREN 
	SCURLY 
	LCURLY 
	RCURLY 
	LSQUAR 
	RSQUAR

	// Operator Types
	ACCESS_OPERATOR
	ASSIGN_OPERATOR
	EVAL_OPERATOR

)