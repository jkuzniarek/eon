package ast 

import (
	"bytes"
)

// eventually rename to Root
type Program struct {
	Expressions []Node 
}
func (p *Program) TokenLiteral() string {
	if len(p.Expressions) > 0 {
		return p.Expressions[0].TokenLiteral()
	} else {
		return ""
	}
}
func (p *Program) String() string{
	var out bytes.Buffer

	for _, s := range p.Expressions{
		out.WriteString(s.String())
	}

	return out.String()
}
