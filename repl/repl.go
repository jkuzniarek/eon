package repl

import (
	"bufio"
	"fmt"
	"io"
	"eon/lexer"
	"eon/parser"
)

const PROMPT = ">_ "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := "(-"+scanner.Text()+")"
		l := lexer.New(line)
		p := parser.New(l, true)

		program := p.ParseShell()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors(), p.Trace)
			continue
		}else{
			// reset trace
			p.Trace = ""
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string, t string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
	io.WriteString(out, "\tTrace: "+t+"\n")
}