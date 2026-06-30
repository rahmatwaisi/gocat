package ignore

import (
	"os"
	"strings"

	"github.com/rahmatwaisi/gocat/utility/configs"
)

// LoadPatterns reads all supported ignore files and returns
// a combined list of ignore patterns.
func LoadPatterns() ([]string, error) {
	var patterns []string

	for _, filename := range configs.IgnoreFiles {
		data, err := os.ReadFile(filename)
		if err != nil {
			continue
		}

		lines := strings.Split(string(data), "\n")

		for _, line := range lines {
			line = strings.TrimSpace(line)

			if line == "" {
				continue
			}

			if strings.HasPrefix(line, "#") {
				continue
			}

			patterns = append(patterns, line)
		}
	}

	return patterns, nil
}
