package object

import (
	"fmt"
	"strings"
)

type ObjectType string

const (
	OBJ_VOID ObjectType = iota
	OBJ_EMPTY
	OBJ_INDEX
	OBJ_LLNODE
	OBJ_BYT
)

type Object interface{
	VMType() ObjectType
	Inspect() string
	UserType() string
	Register() bool // what does this do, is it necessary?
	GetIndex(string i) Object
	SetIndex(string i, Object o, bRef bool) bool
	GetBody() Object
	SetBody(Object o) bool
	
}



// GENERATIONAL REFERENCES

// update these on object creation
uint genCount = 1
genIndex := make(map[&Object]uint)
volIndex := make(map[&Object]bool)
// for compiler add a free-list

// these are used to aggregate code that should also be executed when a compiler would normally execute malloc() and free()
(o *Object) genMalloc(bool vol) {
	// for compiler pull from a free-list if possible. If free-list is empty, call malloc and initialize the generation number to 1
	genIndex[o] = genCount
	volIndex[o] = vol
}
(o *Object) genFree() {
	genCount++;
	// for compiler add code to update the allocation in a free-list
}

type Ref interface{
	Target() &Object 
	BRef() bool // reports if a binding ref or not
}

// Binding Reference
type BRef struct{
	Target &Object
}
(r BRef) Target() &Object{ return r.Target}
(r BRef) BRef() bool{ return true}

// Unbound Reference
type URef struct{
	Target &Object
	Gen uint
}
(r URef) Target() &Object{ return r.Target}
(r URef) BRef() bool{ return false}

// index of bound objects
BindingIndex := make(map[&Object]&BRef)


// EON OBJECT objects

type ObjIndex struct {
	TypeValue string // convert to constant from a map [string]constant
	Index map[string]&Ref
	Body &Object
}
func (o *ObjIndex) Inspect() string { 
	out := "<"
	hasType := false
	hasBody := false
	hasIndex := false
	eoo := true // end of object ">"

	if (len(o.TypeValue) > 0){
		hasType = true
	}
	if (o.Body != nil){
		hasBody = true
	}
	if (len(o.Index)){
		hasIndex = true
	}
	bType := o.Body.VMType() 

	if (hasType && hasBody && !hasIndex && bType == OBJ_BYT && o.TypeValue == "str"){
		out = fmt.Sprintf("\"%s\"", strings.Replace(o.Value, "'", "''", -1))
		eoo = false
	}else if (hasType && hasBody && !hasIndex && bType == OBJ_BYT && o.TypeValue == "int"){ // add handling for int8, uint8, int16, uint16, etc 
		out = fmt.Sprintf("%d", o.Value)
		eoo = false
	}else if (hasType && hasBody && !hasIndex && bType == OBJ_BYT && o.TypeValue == "dec"){ // add handling for dec8, udec8, dec16, udec16, etc
		// for dec8 the first byte (int8) is the number of digits after which the decimal appears
		// 0 indicates all digits are to the right of the decimal
		// sign of the first byte (int8) indicates the sign of the number

		// for udec16 the first 2 bytes (uint16) is the number of digits after which the decimal appears
		// 0 indicates all digits are to the right of the decimal
		// the u in the type udec16 indicates the first 2 bytes (uint8), and therefore the dec number, is unsigned
		out = fmt.Sprintf("%s", "<dec ?>")
		eoo = false
	}else if (hasType && hasBody && !hasIndex && bType == OBJ_BYT && o.TypeValue == "fra"){ // add handling for fra8, ufra8, fra16, ufra16, etc
		// for fra8 the first byte (int8) is the number of ints after which the denominator appears
		// 0 indicates all digits are in the denominator and the numerator is 1
		// sign of the first byte (int8) indicates the sign of the number

		// for ufra16 the first 2 bytes (uint16) is the number of digits after which the denominator appears
		// 0 indicates all digits are in the denominator and the numerator is 1
		// the u in the type ufra16 indicates the first 2 bytes (uint8), and therefore the number, is unsigned
		out = fmt.Sprintf("%s", "<fra ?>")
		eoo = false
	}else if (!hasType && !hasBody && !hasIndex){
		// do nothing
	}else if (hasType && !hasBody && !hasIndex){
		out = out + o.TypeValue
	}else{
		out = out + o.TypeValue + " "
		if hasIndex {
			t := OBJ_VOID
			for k, v := range o.Index {
				t = v.VMType()
				if(t == OBJ_EMPTY){
					out = out + fmt.Sprintf("\n %s ", k)
				}else{
					out = fmt.Sprintf("\n %s: %s ", k, v.Inspect())
				}
			}
			if hasBody {
				out += "\n;"
			}
		}
		if hasBody {
			if bType == OBJ_LLNODE {
				out += o.Body.InspectList()
			}else{
				out += o.Body.Inspect()
			}
		}
	}

	if eoo {
		return out + ">"
	}else{
		return out
	}
	
}
func (o *ObjIndex) VMType() ObjectType { 
	if (len(o.TypeValue) > 0 && len(o.Index) && o.Body != nil){
		return OBJ_EMPTY
	}else{
		return OBJ_INDEX
	}
}
func (o *ObjIndex) UserType() string { return o.TypeValue}
func (o *ObjIndex) Register() &Object { return o.Owner}

// LISTS

func (o *Object) InspectList() string{
	out := ""
	switch o.Body.VMType(){
	case OBJ_LLNODE:
		//
	}
	return out
}
// linked list
type ObjLLNode struct {
	FirstSibling &ObjLLNode
	PrevSibling &ObjLLNode
	NextSibling &ObjLLNode
	NodeObj &Object
}
func (o *ObjNoBody) Inspect() string { return o.InspectList()}
func (o *ObjNoBody) VMType() ObjectType { return OBJ_LLNODE}
func (o *ObjNoBody) UserType() string { return o.TypeValue}

// expression list
// EOL delimited () is stored as {} of strings until an EX or FN keyword converts it to bytecode/function
// () with no EOL characters (or ones that have been commented out) immediately evaluates the contained expression

// FIXED OBJECTS

// struct

// array

// PRIMITIVE

// TODO
type ObjBytes struct {
	Value []byte
}
func (o *ObjBytes) Inspect() string { 
	out := "<byte [ "
	for _, v := range o.Value {
		out += fmt.Sprintf("%d ", v)
	}
	out += "]"
	return out
}
func (o *ObjBytes) VMType() ObjectType { return OBJ_BYT }
func (o *ObjBytes) UserType() string { return "byte"}

// type ObjInt struct {
// 	Value int64
// }
// func (o *ObjInt) Inspect() string { return fmt.Sprintf("%d", o.Value)}
// func (o *ObjInt) VMType() ObjectType { return OBJ_INT }
// func (o *ObjInt) UserType() string { return "str"}

// type ObjStr struct {
// 	Value string
// }
// func (o *ObjStr) Inspect() string { return strings.Replace(o.Value, "'", "''", -1)}
// func (o *ObjStr) VMType() ObjectType { return OBJ_STR }
// func (o *ObjStr) UserType() string { return "str"}

// type ObjDec struct {
// 	// up to 34 digits; equivalent to decimal128 IEEE-754
// 	decPlace int8 // sign indicates sign of total val, number indicates decimal distance from rightmost digit where 0 = no digit < 1
// 	skyVal uint8
// 	topVal uint16
// 	midVal uint32
// 	botVal uint64
// }
// func (o *ObjDec) value() string { return "" } // TODO
// func (o *ObjDec) Inspect() string { 
// 	return fmt.Sprintf("%d", o.Value)
// }
// func (o *ObjDec) VMType() ObjectType { return OBJ_DEC }
// func (o *ObjDec) UserType() string { return "str"}