package token


func IsKeyword(ident string) bool {
	switch ident {
	case "fn", "cfn", "pfn", "conc", "void",
	"esc", "try", "loop", "next", "key",
	"val", "init", "dest", "in", "out",
	"type", "src", "has", "os", "vol",
	"$", "sum", "dif", "mul", "div",
	"exp", "mod":
		return true
	default:
		return false
	}
}

func IsOperator(ident string) bool {
	switch ident {
	case ".", "/", "#", "*", "@",
	":", "::", ":&", ":?", ":+", ":-", ":#",
	"#=", "==", "!=", "<" , ">", "<=", ">=", "|":
		return true
	default:
		return false
	}
}

func GetOperatorType(ident string) TokenType {
	switch ident {
	case ".", "/", "#", "*", "@":
		return ACCESS_OPERATOR
	case ":", "::", ":&", ":?", ":+", ":-", ":#":
		return ASSIGN_OPERATOR
	case "#=", "==", "!=", "<" , ">", "<=", ">=", "|":
		return EVAL_OPERATOR
	default:
		return ILLEGAL
	}
}

func IsOpenDel(ident string) bool {
	switch ident {
	case "(", "{", "[", "(-", "(=", "{:":
		return true
	default:
		return false
	}
}

func IsCloseDel(ident string) bool {
	switch ident {
	case ")", "}", "]":
		return true
	default:
		return false
	}
}

func IsAccessOp(ident string) bool {
	switch ident {
	case ".", "/", "#", "*", "@":
		return true
	default:
		return false
	}
}

func IsAssignOp(ident string) bool {
	switch ident {
	case ":", "::", ":&", ":?", ":+", ":-", ":#":
		return true
	default:
		return false
	}
}

func IsEvalOp(ident string) bool {
	switch ident {
	case "#=", "==", "!=", "<" , ">", "<=", ">=", "|":
		return true
	default:
		return false
	}
}

func (t TokenType) ToStr() string {
	switch t {
	case ILLEGAL:
		return "ILLEGAL"
	case EOF:
		return "EOF"
	case EOL: 
		return "EOL"
	case BSLASH:
		return "BSLASH"
	
	// Identifiers & Literals
	case KEYWORD:
		return "KEYWORD"
	case OPERATOR:
		return "OPERATOR"
	case OPEN_DELIMITER:
		return "OPEN_DELIMITER"
	case CLOSE_DELIMITER:
		return "CLOSE_DELIMITER"
	case ACCESS_OPERATOR:
		return "ACCESS_OPERATOR"
	case ASSIGN_OPERATOR:
		return "ASSIGN_OPERATOR"
	case EVAL_OPERATOR:
		return "EVAL_OPERATOR"
	case PRIMITIVE:
		return "PRIMITIVE"
	case NAME: 
		return "NAME"
	case UINT:
		return "UINT"
	case SINT:
		return "SINT"
	case DEC:
		return "DEC"
	case BYTES:
		return "BYTES"
	case STR:
		return "STR"
	case COMMENT:
		return "COMMENT"
	
	// Delimiters
	case HPAREN:
		return "HPAREN"
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
	return "Undefined Token"
}