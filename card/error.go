package card


type Error struct {
	Message string
}

func (o *Error) IRType() CardType { return ERROR }
func (o *Error) Inspect() string { return "!Error: "+ o.Message}