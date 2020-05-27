// REPL stands for "Read Eval Print Loop".
// The REPL reads input, sends it to the
// interpreter for evaluation, prints the result/output
// of the interpreter and starts again

package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/thewebdevel/monkey-interpreter/lexer"
	"github.com/thewebdevel/monkey-interpreter/token"
)

// PROMPT to the user inside the console or REPL
const PROMPT = ">> "

// Start function reads the input and print out the tokens
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	// Read from the input until encountering a new line
	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		// Take the read line and pass it to an instance of our lexer
		line := scanner.Text()
		l := lexer.New(line)
		// Print all the token the lexer gives us until EOF
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
