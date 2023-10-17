package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	c := flag.Bool("c", false, "Count bytes")

	flag.Parse()

	if *c {
		data, err := readFile(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d\t%s\n", len(data), flag.Arg(0))
	}
}

func readFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}
