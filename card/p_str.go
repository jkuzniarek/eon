package card

import (
	"strings"
)
type Str struct {
	Value string
}
func (o *Str) String() string { return strings.Replace(o.Value, "'", "''", -1)}
func (o *Str) IRType() CardType { return STR }
