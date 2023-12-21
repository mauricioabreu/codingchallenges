package display_test

import (
	"bufio"
	"bytes"
	"os"
	"testing"

	"github.com/mauricioabreu/codingchallenges/cchead/display"
)

func TestDisplayLines(t *testing.T) {
	file, err := os.Open("../text.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	var buffer bytes.Buffer

	if err := display.DisplayLines(bufio.NewReader(file), &buffer, 1); err != nil {
		t.Fatal(err)
	}

	expected := "The Project Gutenberg eBook of The Art of War\n"
	if buffer.String() != expected {
		t.Fatalf("expected %q, got %q", expected, buffer.String())
	}
}

func TestDisplayBytes(t *testing.T) {
	file, err := os.Open("../text.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	var buffer bytes.Buffer

	if err := display.DisplayBytes(bufio.NewReader(file), &buffer, 3); err != nil {
		t.Fatal(err)
	}

	expected := "The"
	if buffer.String() != expected {
		t.Fatalf("expected %q, got %q", expected, buffer.String())
	}
}
