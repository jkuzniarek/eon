package card

/*
// NOT GENERATIONAL REFERENCES
// USE ARC INSTEAD WITH AUTO STRONG/WEAK ASSIGNMENT

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

*/