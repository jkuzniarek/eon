package card

type CardType int

const (
	// special
	VOID CardType = iota
	ANTIVOID
	ERROR
	FUNCTION
	OUT_VAL

	// primitives
	UINT
	SINT
	DEC
	STR

	// composites
	INDEX
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
	case FUNCTION:
		return "FUNCTION"
	case OUT_VAL:
		return "OUT_VAL"
	default:
		return "UNDEFINED_CARD_TYPE"
	}
}