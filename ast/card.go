package ast

import (
	"bytes"
	tk "eon/token"
)

type Card struct{
	Token tk.Token // the open delimiter token
	Type string // the type literal
	Index []Node // name and infix assign expressions
	Body Node // card body expression
}
// TODO: change to include body as part of index with key '/'

func (c *Card) TokenLiteral() string {
	return c.Token.Literal
}
func (c *Card) String() string {
	var out bytes.Buffer
	iLen := len(c.Index)

	if c.Type != "" {
		out.WriteString(c.Type + " ")
	}

	out.WriteString("{")

	if iLen != 0 {
		if iLen == 1 {
			out.WriteString(" ")
			out.WriteString(c.Index[0].String())
		}else{
			for _, e := range c.Index {
				out.WriteString("\n")
				out.WriteString(e.String())
			}
			out.WriteString("\n")
		}
	}

	if c.Body != nil {
		temp := c.Body.String()
		if iLen < 2 && len(temp) < 50 {
			// same line
			out.WriteString("/")
		} else {
			// next line
			out.WriteString("\n/")
		}
		out.WriteString(temp)
	}

	out.WriteString("}")
	return out.String()
}
