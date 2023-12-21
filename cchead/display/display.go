package display

import (
	"bufio"
	"fmt"
	"io"
)

func Display(r *bufio.Reader, w io.Writer, lines int) error {
	for i := 0; i < lines; i++ {
		line, err := r.ReadString('\n')
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
