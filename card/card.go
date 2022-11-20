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
)

type Card interface{
	IRType() CardType
	Inspect() string
	// ByteArray() []byte // everything should be convertible into a byte array
}
