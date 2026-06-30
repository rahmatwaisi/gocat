package version

import "fmt"

var (
	// These values are overridden at build time using -ldflags.
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)

func String() string {
	return Version
}

func Full() string {
	return fmt.Sprintf(
		"%s (commit=%s, built=%s)",
		Version,
		Commit,
		Date,
	)
}
