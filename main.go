package main

import (
	"fmt"
	"os"

	"monkey-lang/repl"
)

// main starts a simple REPL that prints tokens for each input line.
func main() {
	fmt.Println("Monkey programming language! Start typing...")
	repl.Start(os.Stdin, os.Stdout)
}
