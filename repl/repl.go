package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/jmlattanzi/interp/token"

	"github.com/jmlattanzi/interp/lexer"
)

/*

The REPL

The REPL (Read Evaluate Print Loop) takes input and processes it as it is submitted.
Our REPL here is simple. We scan buffered input, pass it to an instance of our Lexer,
and output the generated tokens to standard out.

*/

// PROMPT : The string we are prompted with to enter a command
const PROMPT = "¯\\_(ツ)_/¯ >> "

// Start : Starts the REPL and creates a new scanner. Takes an io.Reader and an io.Writer.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	// while the program is running....
	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		// p := parser.New(l)
		// output := p.ParseProgram()
		// fmt.Println(output.Statements)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}

	}

}
