package formatter

import (
	"bytes"
	"fmt"
	"path/filepath"

	"github.com/rahmatwaisi/gocat/utility/concat"
)

type XMLFormatter struct{}

func NewXML() *XMLFormatter {
	return &XMLFormatter{}
}

func (f *XMLFormatter) Format(documents []concat.Document) ([]byte, error) {
	var out bytes.Buffer

	out.WriteString("<documents>\n")

	for _, document := range documents {

		fmt.Fprintf(
			&out,
			`  <file path="%s">`+"\n",
			filepath.ToSlash(document.Path),
		)

		out.Write(document.Content)

		if len(document.Content) == 0 || document.Content[len(document.Content)-1] != '\n' {
			out.WriteByte('\n')
		}

		out.WriteString("  </file>\n")
	}

	out.WriteString("</documents>\n")

	return out.Bytes(), nil
}
