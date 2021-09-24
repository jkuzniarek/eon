package token

type TokenType int

type Token struct{
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
	INT // 12345
	STR 
	COMMENT 
	
	// Delimiters
	CPAREN 
	LPAREN 
	RPAREN 
	SCURLY 
	LCURLY 
	RCURLY 
	LSQUAR 
	RSQUAR


	// Operators
	DOT 
	SLASH 
	OCTO 
	STAR
	AT 
	PIPE 
	BANG 
	DOLLAR 
	PERCENT 
	CARET
	EQ
	SET_VAL 
	SET_CONST  // constant value
	SET_WEAK  // weak reference (still uses ARC)
	SET_BIND  // binding reference (constraint ref)
	SET_PLUS 
	SET_MINUS 
	SET_TYPE 
	TYPE_EQ 
	EQEQ 
	NOT_EQ 
	LT 
	GT 
	LT_EQ 
	GT_EQ 

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
)

var operators = map[string]TokenType {
	".": DOT,
	"/": SLASH,
	"#": OCTO,
	"*": STAR,
	"@": AT,
	"|": PIPE,
	"!": BANG,
	"$": DOLLAR,
	"%": PERCENT,
	"^": CARET,
	"=": EQ,
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
}

func LookupName(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok 
	}
	return NAME 
}

// checks if the input string is an operator
func LookupOp(op string) TokenType {
	if tok, ok := operators[op]; ok {
		return tok 
	}
	return ILLEGAL 
}

func (t TokenType) ToStr() string {
	switch t {
	case ILLEGAL:
		return "ILLEGAL"
	case EOF:
		return "EOF"
	case EOL: 
		return "EOL"
	
	// Identifiers & Literals
	case TYPE:
		return "TYPE"
	case NAME: 
		return "NAME"
	case INT:
		return "INT"
	case STR:
		return "STR"
	case COMMENT:
		return "COMMENT"
	
	// Delimiters
	case CPAREN:
		return "CPAREN"
	case LPAREN:
		return "LPAREN"
	case RPAREN:
		return "RPAREN"
	case SCURLY:
		return "SCURLY"
	case LCURLY:
		return "LCURLY"
	case RCURLY:
		return "RCURLY"
	case LSQUAR:
		return "LSQUAR"
	case RSQUAR:
		return "RSQUAR"
	}
	for k, v := range keywords {
		if t == v {
			return k
		}
	}
	for k, v := range operators {
		if t == v {
			return k
		}
	}
	return "Undefined Token"
}