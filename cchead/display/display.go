package display

import (
	"bufio"
	"fmt"
	"io"
)

func DisplayLines(r io.Reader, w io.Writer, nLines int) error {
	bufReader := bufio.NewReader(r)
	for i := 0; i < nLines; i++ {
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

func DisplayBytes(r io.Reader, w io.Writer, nBytes int) error {
	content := make([]byte, nBytes)

	bufReader := bufio.NewReader(r)
	bytesRead, err := bufReader.Read(content)
	if err != nil {
		return err
	}

	if _, err := fmt.Fprint(w, string(content[:bytesRead])); err != nil {
		return err
	}

	return nil
}
