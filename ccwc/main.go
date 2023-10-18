package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Opts struct {
	countBytes bool
	countLines bool
	countWords bool
}

func main() {
	c := flag.Bool("c", false, "Count bytes")
	l := flag.Bool("l", false, "Count lines")
	w := flag.Bool("w", false, "Count words")

	flag.Parse()

	opts := Opts{
		countBytes: *c,
		countLines: *l,
		countWords: *w,
	}

	var (
		file *os.File
		err  error
	)

	filename := flag.Arg(0)
	input := os.Stdin

	if filename != "" {
		file, err = os.Open(filename)
		if err != nil {
			log.Fatalf("failed to read file: %s", err)
		}
		defer file.Close()
		input = file
	}

	stats, err := fetchStats(input)
	if err != nil {
		log.Fatalf("failed to fetch stats: %s", err)
	}

	fmt.Print(formatStats(opts, filename, stats))
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
		prevWS    int
	)

	reader := bufio.NewReader(f)
	stats := Stats{}

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			return Stats{}, err
		}

		stats.nbytes++
		stats.lines += int(linePattern[b])
		stats.words += int(wsPattern[b]) & prevWS
		prevWS = int(wsPattern[b]) ^ 1
	}

	stats.words += prevWS

	return stats, nil
}

func formatStats(opts Opts, filename string, stats Stats) string {
	v := make([]string, 0)

	if opts.countLines {
		v = append(v, fmt.Sprint(stats.lines))
	}
	if opts.countWords {
		v = append(v, fmt.Sprint(stats.words))
	}
	if opts.countBytes {
		v = append(v, fmt.Sprint(stats.nbytes))
	}
	if !opts.countLines && !opts.countWords && !opts.countBytes {
		v = append(v, fmt.Sprint(stats.lines, stats.words, stats.nbytes))
	}

	output := strings.Join(v, "\t")

	if filename != "" {
		output += "\t" + filename
	}

	return output + "\n"
}
