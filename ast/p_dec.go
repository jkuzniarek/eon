package ast

import (
	tk "eon/token"
	"bytes"
	ssDec "github.com/shopspring/decimal"
)

// dec
type Dec struct{
	Token tk.Token 
	Value ssDec.Decimal
}

func (il *Dec) TokenLiteral() string {
	return il.Token.Literal
}
func (il *Dec) String() string{
	return il.Value.String()
}
