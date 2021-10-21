package main

import (
	"fmt"
	"os"
	"eon/repl"
)

func main() {
	if(len(os.Args) == 1){
		fmt.Printf("Welcome to the eon Shell\n")
		repl.Shell(os.Stdin, os.Stdout)
	}
	// else if(os.Args[1] == "bild"){
	// 	// compile
	// }
	// else if(os.Args[1] == "run"){
	// 	// interpret
	// }
}