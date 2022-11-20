package card

import(
	"fmt"
)

type UInt struct {
	Value uint
}
func (o *UInt) Inspect() string { return fmt.Sprintf("%d", o.Value)}
func (o *UInt) IRType() CardType { return UINT }