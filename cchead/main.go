// cchead
// My own version of Unix head tool
package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"

	"github.com/mauricioabreu/codingchallenges/cchead/display"
)

func main() {
	n := flag.Int("n", 10, "Number of lines")
	c := flag.Int("c", 0, "Number of bytes")

	flag.Parse()

	var (
		file *os.File
		err  error
	)

	filename := flag.Arg(0)
	var input io.Reader = os.Stdin

	if filename != "" {
		file, err = os.Open(filename)
		if err != nil {
			log.Fatalf("failed to read file: %s", err)
		}
		defer file.Close()
		input = bufio.NewReader(file)
	}

	if *c > 0 {
		if err := display.DisplayBytes(input, os.Stdout, *c); err != nil {
			log.Fatal(err)
		}
	} else {
		if display.DisplayLines(input, os.Stdout, *n); err != nil {
			log.Fatal(err)
		}
	}
}
