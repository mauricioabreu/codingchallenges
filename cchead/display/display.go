package display

import (
	"bufio"
	"fmt"
	"io"
)

func Display(r io.Reader, w io.Writer, lines int) error {
	bufReader := bufio.NewReader(r)
	for i := 0; i < lines; i++ {
		line, err := bufReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if _, err := fmt.Fprint(w, line); err != nil {
			return err
		}
	}

	return nil
}
