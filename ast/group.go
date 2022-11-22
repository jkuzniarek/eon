package ast

import (
	tk "eon/token"
	"bytes"
)

type Group struct{
	Token tk.Token // the open delimiter token
	Expressions []Node 
}

func (g *Group) TokenLiteral() string {
	return g.Token.Literal
}
func (g *Group) String() string {
	var out bytes.Buffer
	iLen := len(g.Expressions)

	out.WriteString(g.TokenLiteral())

	if iLen != 0 {
		if iLen == 1 {
			out.WriteString(g.Expressions[0].String())
		}else{
			out.WriteString("\n")
			for _, e := range g.Expressions {
				out.WriteString(e.String())
				out.WriteString("\n")
			}
		}
	}
	
	switch g.Token.Type {
	case tk.HPAREN, tk.CPAREN:
		out.WriteString(")")
	case tk.LSQUAR:
		out.WriteString("]")
	case tk.LCURLY, tk.SCURLY:
		out.WriteString("}")
	default:
		out.WriteString(" ")
	}
	return out.String()
}
