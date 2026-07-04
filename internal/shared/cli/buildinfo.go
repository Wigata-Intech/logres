package cli

// Overridden via -ldflags -X at build time (see Makefile).
var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)
