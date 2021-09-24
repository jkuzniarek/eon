package run

import (
	"fmt"
	"eon/lexer"
	"eon/token"
)

func main(in string){
	
	l := lexer.New(in)
	
	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		// print token for debug
		fmt.Fprintf(out, "%+v\n", tok)
		
	}
}