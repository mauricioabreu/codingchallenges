package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	c := flag.Bool("c", false, "Count bytes")
	l := flag.Bool("l", false, "Count lines")

	flag.Parse()

	data, err := readFile(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	if *c {
		fmt.Printf("%d\t%s\n", len(data), flag.Arg(0))
	}

	if *l {
		lines, err := countLines(bytes.NewReader(data))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d\t%s\n", lines, flag.Arg(0))
	}
}

func readFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func countLines(data io.Reader) (int, error) {
	var (
		lines int
		read  int
		err   error
		sep   = []byte("\n")
	)

	buf := make([]byte, 32*1024)

	for {
		read, err = data.Read(buf)
		if err != nil {
			break
		}

		lines += bytes.Count(buf[:read], sep)
	}

	if err == io.EOF {
		return lines, nil
	}

	return 0, nil
}
