package tree

type Node struct {
	Name     string
	FileNum  int
	DirNum   int
	Children map[string]*Node
}
