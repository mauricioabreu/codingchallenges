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

	display.Display(input, os.Stdout, *n)
}
