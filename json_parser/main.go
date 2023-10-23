package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/mauricioabreu/codingchallenges/json_parser/lexer"
	"github.com/mauricioabreu/codingchallenges/json_parser/parser"
)

func main() {
	var input io.Reader

	filename := flag.Arg(0)

	if filename != "" {
		file, err := os.Open(filename)
		if err != nil {
			log.Fatalf("failed to read file: %s", err)
		}
		defer file.Close()
		input = file
	} else {
		input = os.Stdin
	}

	lxr := lexer.NewLexer(input)
	psr := parser.NewParser(lxr)

	if psr.Parse() {
		fmt.Println("Valid JSON")
		os.Exit(0)
	} else {
		fmt.Println("Invalid JSON")
		os.Exit(1)
	}
}
