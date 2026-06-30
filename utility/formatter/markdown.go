package formatter

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/rahmatwaisi/gocat/utility/concat"
)

type MarkdownFormatter struct{}

func NewMarkdown() *MarkdownFormatter {
	return &MarkdownFormatter{}
}

func (f *MarkdownFormatter) Format(documents []concat.Document) ([]byte, error) {
	var out bytes.Buffer

	for i, document := range documents {

		ext := strings.TrimPrefix(filepath.Ext(document.Path), ".")

		if ext == "" {
			ext = "text"
		}

		fmt.Fprintf(&out, "# %s\n\n", filepath.ToSlash(document.Path))
		fmt.Fprintf(&out, "```%s\n", ext)

		out.Write(document.Content)

		if len(document.Content) == 0 || document.Content[len(document.Content)-1] != '\n' {
			out.WriteByte('\n')
		}

		out.WriteString("```\n")

		if i != len(documents)-1 {
			out.WriteString("\n---\n\n")
		}
	}

	return out.Bytes(), nil
}
