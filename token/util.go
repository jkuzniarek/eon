package token


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

func IsKeyword(ident string) bool {
	if tok, ok := keywords[ident]; ok {
		return true
	}
	return false
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
	
	// Identifiers & Literals
	case TYPE:
		return "TYPE"
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