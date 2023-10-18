package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	flag.Parse()

	var (
		f   *os.File
		err error
	)

	if flag.Arg(0) != "" {
		f, err = os.Open(flag.Arg(0))
		if err != nil {
			log.Fatal("failed to read file: ", err)
		}
		defer f.Close()
	} else {
		f = os.Stdin
	}

	stats, err := fetchStats(f)
	if err != nil {
		log.Fatal("failed to fetch stats: ", err)
	}

	fmt.Printf("%d %d %d\n", stats.lines, stats.words, stats.nbytes)
}

type Stats struct {
	words  int
	lines  int
	nbytes int
}

func fetchStats(f *os.File) (Stats, error) {
	var (
		linePattern = [256]uint8{'\n': 1}
		// https://en.cppreference.com/w/cpp/string/wide/iswspace
		wsPattern = [256]uint8{' ': 1, '\f': 1, '\n': 1, '\r': 1, '\t': 1, '\v': 1}
	)

	words, lines, nbytes, prevWS := 0, 0, 0, 0
	reader := bufio.NewReader(f)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			return Stats{words: 0, lines: 0, nbytes: 0}, err
		}

		nbytes++
		lines += int(linePattern[b])
		words += int(wsPattern[b]) & prevWS
		prevWS = int(wsPattern[b]) ^ 1
	}

	return Stats{
		words:  words + prevWS,
		lines:  lines,
		nbytes: nbytes,
	}, nil
}
