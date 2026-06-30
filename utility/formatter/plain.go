package formatter

import (
	"bytes"
	"fmt"
	"path/filepath"

	"github.com/rahmatwaisi/gocat/utility/concat"
)

type PlainFormatter struct{}

func NewPlain() *PlainFormatter {
	return &PlainFormatter{}
}

func (f *PlainFormatter) Format(documents []concat.Document) ([]byte, error) {
	var out bytes.Buffer

	for i, document := range documents {

		fmt.Fprintf(&out, "%s\n\n", filepath.ToSlash(document.Path))
		out.Write(document.Content)

		if len(document.Content) == 0 || document.Content[len(document.Content)-1] != '\n' {
			out.WriteByte('\n')
		}

		if i != len(documents)-1 {
			out.WriteString("\n---\n\n")
		}
	}

	return out.Bytes(), nil
}
