package card

import (
	"eon/ast"
)

type Function struct {
	Body *ast.Group
	Env *Env
}

func (o *Function) IRType() CardType { return FUNCTION }
func (o *Function) String() string { return o.Body.String() }