package card

import (
	ssDec "github.com/shopspring/decimal"
)

type Dec struct {
	Value ssDec.Decimal
}
func (o *Dec) Inspect() string { 
	return o.Value.String()
}
func (o *Dec) IRType() CardType { return DEC }