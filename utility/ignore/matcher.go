package ignore

import (
	"path/filepath"
	"strings"
)

type Matcher struct {
	patterns []string
}

func NewMatcher() (*Matcher, error) {
	patterns, err := LoadPatterns()
	if err != nil {
		return nil, err
	}

	return &Matcher{
		patterns: patterns,
	}, nil
}

// ShouldIgnore reports whether a path should be ignored.
func (m *Matcher) ShouldIgnore(path string, isDir bool) bool {
	if IsWhitelisted(path) {
		return false
	}

	path = filepath.ToSlash(path)
	parts := strings.Split(path, "/")

	for _, pattern := range m.patterns {

		pattern = filepath.ToSlash(pattern)
		pattern = strings.TrimSuffix(pattern, "/")

		// Exact directory/file name
		for _, part := range parts {
			if part == pattern {
				return true
			}
		}

		// *.ext
		if strings.HasPrefix(pattern, "*") {
			if strings.HasSuffix(path, pattern[1:]) {
				return true
			}
		}

		// path prefix
		if path == pattern {
			return true
		}

		if strings.HasPrefix(path, pattern+"/") {
			return true
		}
	}

	return false
}
