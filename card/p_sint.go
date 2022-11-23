package card

import(
	"fmt"
)

type SInt struct {
	Value int
}
func (o *SInt) String() string { 
	if o.Value >= 0 {
		return fmt.Sprintf("+%d", o.Value)
	}else{
		return fmt.Sprintf("-%d", o.Value)
	}
}
func (o *SInt) IRType() CardType { return SINT }