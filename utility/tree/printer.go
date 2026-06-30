package tree

import (
	"fmt"
	"github.com/rahmatwaisi/gocat/utility/configs"
	"path/filepath"
	"sort"
	"strings"

	"github.com/rahmatwaisi/gocat/utility/scanner"
)

type Tree struct {
	root *Node
}

func New(result *scanner.ScanResult) *Tree {
	root := &Node{
		Name:     ".",
		Children: make(map[string]*Node),
	}

	dirIndex := make(map[string]int)

	for i, dir := range result.Dirs {
		dirIndex[filepath.ToSlash(dir)] = i + 1
	}

	for fileIndex, file := range result.Files {

		parts := strings.Split(filepath.ToSlash(file.Path), "/")

		current := root

		for i, part := range parts {

			isLast := i == len(parts)-1

			node, exists := current.Children[part]
			if !exists {

				node = &Node{
					Name:     part,
					Children: make(map[string]*Node),
				}

				if isLast {

					node.FileNum = fileIndex + 1

				} else {

					partial := strings.Join(parts[:i+1], "/")

					if dirNum, ok := dirIndex[partial]; ok {
						node.DirNum = dirNum
					}
				}

				current.Children[part] = node
			}

			current = node
		}
	}

	return &Tree{
		root: root,
	}
}

func (t *Tree) Print() {
	printNode(t.root, "", true, true)
}

func printNode(node *Node, indent string, isLast bool, isRoot bool) {

	if !isRoot {

		fmt.Print(indent)

		if isLast {
			fmt.Print("└── ")
			indent += "    "
		} else {
			fmt.Print("├── ")
			indent += "│   "
		}

		switch {

		case node.FileNum > 0:
			fmt.Printf("[%d] %s %s\n",
				node.FileNum,
				configs.FileIcon,
				node.Name,
			)

		case node.DirNum > 0:
			fmt.Printf("[d%d] %s %s/\n",
				node.DirNum,
				configs.FolderIcon,
				node.Name,
			)

		default:
			fmt.Printf("%s/\n", node.Name)
		}
	}

	keys := make([]string, 0, len(node.Children))

	for k := range node.Children {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for i, k := range keys {
		printNode(
			node.Children[k],
			indent,
			i == len(keys)-1,
			false,
		)
	}
}
