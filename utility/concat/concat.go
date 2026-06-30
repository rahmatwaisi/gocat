package concat

import (
	"os"

	"github.com/rahmatwaisi/gocat/utility/scanner"
)

type Document struct {
	Path    string
	Content []byte
}

type Concatenator struct{}

func New() *Concatenator {
	return &Concatenator{}
}

// Read loads the contents of every selected file.
func (c *Concatenator) Read(files []scanner.FileInfo) ([]Document, error) {
	documents := make([]Document, 0, len(files))

	for _, file := range files {

		content, err := os.ReadFile(file.Path)
		if err != nil {
			return nil, err
		}

		documents = append(documents, Document{
			Path:    file.Path,
			Content: content,
		})
	}

	return documents, nil
}
