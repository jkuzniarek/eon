package card

/*
type IndexCard struct {
	TypeConstant int 
	Index map[string]*Card
}
func (o *Index) String() string { 
	out := "{"
	hasType := false
	hasBody := false
	hasIndex := false
	eoo := true // end of card "}"

	if (len(o.TypeValue) > 0){
		hasType = true
	}
	if (o.Body != nil){
		hasBody = true
	}
	if (len(o.Index)){
		hasIndex = true
	}
	bType := o.Body.IRType() 

	if (hasType && hasBody && !hasIndex && bType == BYT && o.TypeValue == "str"){
		out = fmt.Sprintf("\"%s\"", strings.Replace(o.Value, "'", "''", -1))
		eoo = false
	}else if (hasType && hasBody && !hasIndex && bType == BYT && o.TypeValue == "int"){ // add handling for int8, uint8, int16, uint16, etc 
		out = fmt.Sprintf("%d", o.Value)
		eoo = false
	}else if (hasType && hasBody && !hasIndex && bType == BYT && o.TypeValue == "dec"){ // add handling for dec8, udec8, dec16, udec16, etc
		// for dec8 the first byte (int8) is the number of digits after which the decimal appears
		// 0 indicates all digits are to the right of the decimal
		// sign of the first byte (int8) indicates the sign of the number

		// for udec16 the first 2 bytes (uint16) is the number of digits after which the decimal appears
		// 0 indicates all digits are to the right of the decimal
		// the u in the type udec16 indicates the first 2 bytes (uint8), and therefore the dec number, is unsigned
		out = fmt.Sprintf("%s", "dec ?")
		eoo = false
	}else if (hasType && hasBody && !hasIndex && bType == BYT && o.TypeValue == "fra"){ // add handling for fra8, ufra8, fra16, ufra16, etc
		// for fra8 the first byte (int8) is the number of ints after which the denominator appears
		// 0 indicates all digits are in the denominator and the numerator is 1
		// sign of the first byte (int8) indicates the sign of the number

		// for ufra16 the first 2 bytes (uint16) is the number of digits after which the denominator appears
		// 0 indicates all digits are in the denominator and the numerator is 1
		// the u in the type ufra16 indicates the first 2 bytes (uint8), and therefore the number, is unsigned
		out = fmt.Sprintf("%s", "fra ?")
		eoo = false
	}else if (!hasType && !hasBody && !hasIndex){
		// do nothing
	}else if (hasType && !hasBody && !hasIndex){
		out = o.TypeValue + " " + out
	}else{
		out = o.TypeValue + " " + out
		if hasIndex {
			t := VOID
			for k, v := range o.Index {
				t = v.IRType()
				if(t == EMPTY){
					out = out + fmt.Sprintf("\n %s ", k)
				}else{
					out = fmt.Sprintf("\n %s: %s ", k, v.String())
				}
			}
			if hasBody {
				out += "\n/"
			}
		}
		if hasBody {
			if bType == LLNODE {
				out += o.Body.InspectList()
			}else{
				out += o.Body.String()
			}
		}
	}

	if eoo {
		return out + "}"
	}else{
		return out
	}
	
}
func (o *Index) IRType() CardType { 
	return INDEX
}


// TODO
// GetCardType() string
// SetUserType() string
// GetIndex(string i) Card
// SetIndex(string i, Card o, bRef bool) bool
// GetBody() Card
// SetBody(Card o) bool

*/