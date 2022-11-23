package card


type OutVal struct {
	Value Card
}

func (o *OutVal) IRType() CardType { return OUT_VAL }
func (o *OutVal) String() string { return o.Value.String() }
