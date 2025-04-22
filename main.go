package main

import (
	"fmt"
	"os"

	"monkey-lang/repl"
)

func main() {
	fmt.Println("Monkey programming language! Start typing...")
	repl.Start(os.Stdin, os.Stdout)
}
