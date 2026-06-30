package configs

const (
	FolderIcon = "📁"
	FileIcon   = "📄"
)

const (
	// AppName Application
	AppName = "gocat"
	Version = "dev"

	// DefaultOutputFile Output
	DefaultOutputFile = "output.md"

	// GitIgnoreFile Supported ignore files (loaded if present)
	GitIgnoreFile    = ".gitignore"
	DockerIgnoreFile = ".dockerignore"
	AIIgnoreFile     = ".aiignore"
	GocatIgnoreFile  = ".gocatignore"
)

var IgnoreFiles = []string{
	GitIgnoreFile,
	DockerIgnoreFile,
	AIIgnoreFile,
	GocatIgnoreFile,
}

var WhitelistPatterns = []string{
	"Dockerfile",
	".dockerignore",
	"docker-compose.yml",
	"docker-compose.*.yml",
	"*.Dockerfile",
}
