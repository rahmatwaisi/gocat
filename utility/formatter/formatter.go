package formatter

import "github.com/rahmatwaisi/gocat/utility/concat"

type Formatter interface {
	Format(documents []concat.Document) ([]byte, error)
}
