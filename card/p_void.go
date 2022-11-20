package card


type Void struct {
	Value Card
}
func (o *Void) Inspect() string { 
	if o.Value == nil {
		return "void"
	}else if (o.Value).IRType() == VOID {
		return "{}"
	} 
	return "void " + (o.Value).Inspect()
}
func (o *Void) IRType() CardType { 
	if o.Value == nil {
		return VOID
	}else if (o.Value).IRType() == VOID {
		return ANTIVOID
	} 
	return VOID 
}