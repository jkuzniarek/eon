package eval

import (
	"eon/card"
)

func boolToVoid(input bool) *card.Void {
	if input == true {
		return ANTIVOID
	}
	return VOID
}