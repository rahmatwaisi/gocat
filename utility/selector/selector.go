package selector

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/rahmatwaisi/gocat/utility/scanner"
)

type Selector struct {
	result *scanner.ScanResult
}

func New(result *scanner.ScanResult) *Selector {
	return &Selector{
		result: result,
	}
}

// Select prompts the user and returns the selected files.
func (s *Selector) Select() ([]scanner.FileInfo, error) {
	fmt.Print("\nEnter numbers to select (e.g. 1 3 d2 d5): ")

	reader := bufio.NewScanner(os.Stdin)

	if !reader.Scan() {
		return nil, fmt.Errorf("failed to read input")
	}

	input := strings.ReplaceAll(reader.Text(), ",", " ")
	fields := strings.Fields(input)

	selected := make(map[int]struct{})

	for _, field := range fields {

		field = strings.ToLower(strings.TrimSpace(field))

		if strings.HasPrefix(field, "d") {

			dirNum, err := strconv.Atoi(field[1:])
			if err != nil {
				fmt.Printf("Skipping invalid directory selector: %q\n", field)
				continue
			}

			if dirNum < 1 || dirNum > len(s.result.Dirs) {
				fmt.Printf("Skipping invalid directory selector: %q\n", field)
				continue
			}

			dirPath := filepath.ToSlash(s.result.Dirs[dirNum-1]) + "/"

			for index, file := range s.result.Files {
				if strings.HasPrefix(filepath.ToSlash(file.Path), dirPath) {
					selected[index] = struct{}{}
				}
			}

			continue
		}

		fileNum, err := strconv.Atoi(field)
		if err != nil {
			fmt.Printf("Skipping invalid selector: %q\n", field)
			continue
		}

		if fileNum < 1 || fileNum > len(s.result.Files) {
			fmt.Printf("Skipping invalid selector: %q\n", field)
			continue
		}

		selected[fileNum-1] = struct{}{}
	}

	if len(selected) == 0 {
		return nil, fmt.Errorf("no valid files selected")
	}

	indices := make([]int, 0, len(selected))

	for index := range selected {
		indices = append(indices, index)
	}

	sort.Ints(indices)

	files := make([]scanner.FileInfo, 0, len(indices))

	for _, index := range indices {
		files = append(files, s.result.Files[index])
	}

	return files, nil
}
