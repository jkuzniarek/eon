package evaluator

import (
	"eon/token"
	"eon/object"
)

func Eval(tok token.Token) object.Object {
	switch tok.Type {
	case token.INT:
		return &object.Integer{}
	}
}
