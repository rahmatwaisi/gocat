package scanner

import (
	"io/fs"
	"path/filepath"

	"github.com/rahmatwaisi/gocat/utility/ignore"
)

type Scanner struct {
	matcher *ignore.Matcher
}

func New(matcher *ignore.Matcher) *Scanner {
	return &Scanner{
		matcher: matcher,
	}
}

// Scan recursively scans the given root directory.
func (s *Scanner) Scan(root string) (*ScanResult, error) {
	result := &ScanResult{}

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {

		if err != nil {
			return nil
		}

		if path == root {
			return nil
		}

		if !d.IsDir() && ignore.IsWhitelisted(path) {
			result.Files = append(result.Files, FileInfo{
				Path: filepath.ToSlash(path),
			})
			return nil
		}

		if s.matcher.ShouldIgnore(path, d.IsDir()) {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if d.IsDir() {
			result.Dirs = append(result.Dirs, filepath.ToSlash(path))
			return nil
		}

		result.Files = append(result.Files, FileInfo{
			Path: filepath.ToSlash(path),
		})

		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
