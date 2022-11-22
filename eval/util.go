package eval

import (
	"eon/card"
	"fmt"
)

func boolToVoid(input bool) *card.Void {
	if input == true {
		return ANTIVOID
	}
	return VOID
}

func newError(format string, a ...interface{}) *card.Error {
	return &card.Error{Message: fmt.Sprintf(format, a...)}
}

func isError(o card.Card) bool{
	if o != nil {
		return o.IRType() == card.ERROR
	}
	return false
}
