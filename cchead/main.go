// cchead
// My own version of Unix head tool
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/mauricioabreu/codingchallenges/cchead/display"
)

func main() {
	n := flag.Int("n", 10, "Number of lines")
	c := flag.Int("c", 0, "Number of bytes")

	flag.Parse()
	filenames := flag.Args()

	if len(filenames) == 0 {
		handleStdIn(*n, *c)
	} else {
		for _, filename := range filenames {
			handleFile(filename, len(filenames) > 1, *n, *c)
		}
	}
}

func handleStdIn(nLines, nBytes int) {
	if err := handle(os.Stdin, "standard input", nLines, nBytes); err != nil {
		log.Fatalf("Error processing stdin: %s", err)
	}
}

func handleFile(filename string, addHeader bool, nLines, nBytes int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error processing file '%s': %s", filename, err)
	}
	defer file.Close()

	if addHeader {
		fmt.Printf("==> %s <==\n", filename)
	}

	if err := handle(bufio.NewReader(file), filename, nLines, nBytes); err != nil {
		log.Fatalf("Error processing file '%s': %s", filename, err)
	}
}

func handle(input io.Reader, filename string, nLines int, nBytes int) error {
	if nBytes > 0 {
		return display.DisplayBytes(input, os.Stdout, nBytes)
	} else {
		return display.DisplayLines(input, os.Stdout, nLines)
	}
}
