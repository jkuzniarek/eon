package card


type OutVal struct {
	Value Card
}

func (o *OutVal) IRType() CardType { return OUT_VAL }
func (o *OutVal) Inspect() string { return o.Value.Inspect() }
