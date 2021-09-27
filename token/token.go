package token

type TokenType int
type CatType int

type Token struct{
	Cat CatType
	Type TokenType
	Literal string
}

const (
	_ = iota // makes the constants auto incremented integers
	ILLEGAL
	EOF
	EOL

	// Identifiers & Literals
	TYPE
	NAME // add, foobar, x, y, ...
	KEYWORD
	OPEN_DELIMITER
	CLOSE_DELIMITER
	ACCESS_OPERATOR
	ASSIGN_OPERATOR
	EVAL_OPERATOR
	PRIMITIVE
	UINT // 12345
	SINT // +1, -2, +0
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


	// Access Operators
	DOT 
	SLASH 
	OCTO 
	STAR
	AT 

	// Assign Operators
	SET_VAL 
	SET_CONST  // constant value
	SET_WEAK  // weak reference (still uses ARC)
	SET_BIND  // binding reference (constraint ref)
	SET_PLUS 
	SET_MINUS 
	SET_TYPE 

	// Eval Operators
	TYPE_EQ 
	EQEQ 
	NOT_EQ 
	LT 
	GT 
	LT_EQ 
	GT_EQ 
	PIPE 

	// Keywords
	FN 
	CFN
	PFN
	CONC
	VOID 
	ESC 
	TRY 
	LOOP 
	NEXT 
	KEY 
	VAL 
	INIT 
	DEST 
	IN 
	OUT 
	TYPE_DEF 
	SRC 
	HAS 
	OS 
	VOL
	DOLLAR 
	SUM 
	DIF 
	MUL 
	DIV 
	EXP 
	MOD 
)

var operators = map[string]TokenType {
	".": DOT,
	"/": SLASH,
	"#": OCTO,
	"*": STAR,
	"@": AT,
	"|": PIPE,
	"$": DOLLAR,
	":": SET_VAL,
	"::": SET_CONST,
	":?": SET_WEAK,
	":&": SET_BIND,
	":+": SET_PLUS,
	":-": SET_MINUS,
	":#": SET_TYPE,
	"#=": TYPE_EQ,
	"==": EQEQ,
	"!=": NOT_EQ,
	"<": LT,
	">": GT,
	"<=": LT_EQ,
	">=": GT_EQ, 
}

var keywords = map[string]TokenType {
	"fn": FN,
	"cfn": CFN,
	"pfn": PFN,
	"conc": CONC,
	"void": VOID,
	"esc": ESC,
	"try": TRY,
	"loop": LOOP,
	"next": NEXT,
	"key": KEY,
	"val": VAL,
	"init": INIT,
	"dest": DEST,
	"in": IN,
	"out": OUT,
	"type": TYPE_DEF,
	"src": SRC,
	"has": HAS,
	"os": OS,
	"vol": VOL,
	"$": DOLLAR,
	"sum": SUM,
	"dif": DIF,
	"mul": MUL,
	"div": DIV,
	"exp": EXP,
	"mod": MOD,
}
