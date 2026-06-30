package ignore

import (
	"path/filepath"

	"github.com/rahmatwaisi/gocat/utility/configs"
)

// IsWhitelisted reports whether a file should always be included
// regardless of ignore rules.
func IsWhitelisted(path string) bool {
	name := filepath.Base(path)

	for _, pattern := range configs.WhitelistPatterns {

		if ok, err := filepath.Match(pattern, name); err == nil && ok {
			return true
		}

		if ok, err := filepath.Match(pattern, filepath.ToSlash(path)); err == nil && ok {
			return true
		}
	}

	return false
}
