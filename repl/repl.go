package repl

import (
	"bufio"
	"fmt"
	"io"

	"monkey-lang/lexer"
	"monkey-lang/token"
)

const PROMPT = ">> "

// Start launches a simple token-printing REPL over the provided streams.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		if !scanner.Scan() {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		// Tokenize the line and print each token until EOF.
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}


