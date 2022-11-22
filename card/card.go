package card

// import (
// 	"fmt"
// 	"strings"
// )

type CardType int

const (
	VOID CardType = iota
	ANTIVOID
	UINT
	INDEX
	ERROR
)

type Card interface{
	IRType() CardType
	Inspect() string
	// ByteArray() []byte // everything should be convertible into a byte array
}


func (o CardType) String() string {
	switch o {
	case VOID:
		return "VOID"
	case ANTIVOID:
		return "ANTIVOID"
	case UINT:
		return "UINT"
	case INDEX:
		return "INDEX"
	case ERROR:
		return "ERROR"
	default:
		return "UNDEFINED_CARD_TYPE"
	}
}