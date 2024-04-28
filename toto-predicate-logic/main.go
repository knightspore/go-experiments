package main

import (
	"encoding/json"
	"fmt"
)

const SOURCE_PATH = "samples/test.toto"

func main() {
	l := NewLexer(SOURCE_PATH)
	fmt.Printf("Parsing Source: %q (%d chars)\n", SOURCE_PATH, len(SOURCE_PATH))

	tokens := l.Tokens()
	fmt.Printf("> Created %d tokens\n", len(tokens))

	s := ParseStatements(l)
	fmt.Printf("> Created %d statements\n", len(s))

	fmt.Printf("> Printing statement tree:\n")
	pretty, _ := json.MarshalIndent(s, "\t", "  ")
	fmt.Printf("%s\n", pretty)
}
