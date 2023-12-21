package display_test

import (
	"bufio"
	"bytes"
	"os"
	"testing"

	"github.com/mauricioabreu/codingchallenges/cchead/display"
)

func TestDisplay(t *testing.T) {
	file, err := os.Open("../text.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	var buffer bytes.Buffer

	if err := display.Display(bufio.NewReader(file), &buffer, 1); err != nil {
		t.Fatal(err)
	}

	expected := "The Project Gutenberg eBook of The Art of War\n"
	if buffer.String() != expected {
		t.Fatalf("expected %q, got %q", expected, buffer.String())
	}
}
