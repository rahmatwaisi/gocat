package scanner

// FileInfo represents a discovered file.
type FileInfo struct {
	Path string
}

// ScanResult contains the complete filesystem scan.
type ScanResult struct {
	Files []FileInfo
	Dirs  []string
}
