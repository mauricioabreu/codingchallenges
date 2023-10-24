package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mauricioabreu/codingchallenges/compressor/compress"
)

func main() {
	flag.Parse()

	filename := flag.Arg(0)
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read file %s: %s", filename, err)
	}

	fmt.Println(compress.Count(data))
}
