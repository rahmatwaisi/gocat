package writer

import (
	"fmt"
	"os"
)

type Writer struct{}

func New() *Writer {
	return &Writer{}
}

// WriteFile writes data to the specified file.
func (w *Writer) WriteFile(filename string, data []byte) error {
	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("write %s: %w", filename, err)
	}

	return nil
}

// WriteStdout writes data to stdout.
func (w *Writer) WriteStdout(data []byte) error {
	_, err := os.Stdout.Write(data)
	if err != nil {
		return fmt.Errorf("write stdout: %w", err)
	}

	return nil
}
