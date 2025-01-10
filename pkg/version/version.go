package version

import (
	"fmt"
	"runtime/debug"
	"strings"
)

var (
	// Version is the current version of the package
	Version = "v0.1.0"
	
	// BuildTime is the time the binary was built
	BuildTime string
	
	// GitCommit is the git commit hash of the build
	GitCommit string
)

// Info represents version information
type Info struct {
	Version   string
	BuildTime string
	GitCommit string
	GoVersion string
}

// Get returns the current version information
func Get() Info {
	var info Info
	
	// Get build info
	if bi, ok := debug.ReadBuildInfo(); ok {
		info.GoVersion = bi.GoVersion
		for _, s := range bi.Settings {
			switch s.Key {
			case "vcs.revision":
				info.GitCommit = s.Value
			case "vcs.time":
				info.BuildTime = s.Value
			}
		}
	}
	
	// Use hardcoded values if available
	if Version != "" {
		info.Version = Version
	}
	if BuildTime != "" {
		info.BuildTime = BuildTime
	}
	if GitCommit != "" {
		info.GitCommit = GitCommit
	}
	
	return info
}

// String returns a string representation of version info
func (i Info) String() string {
	return fmt.Sprintf("Version: %s\nBuild Time: %s\nGit Commit: %s\nGo Version: %s",
		i.Version, i.BuildTime, i.GitCommit, i.GoVersion)
}

// IsCompatible checks if the given version is compatible with the current version
func IsCompatible(reqVersion string) bool {
	// Remove 'v' prefix if present
	current := strings.TrimPrefix(Version, "v")
	required := strings.TrimPrefix(reqVersion, "v")
	
	// Split versions into parts
	currentParts := strings.Split(current, ".")
	requiredParts := strings.Split(required, ".")
	
	// Compare major versions
	if currentParts[0] != requiredParts[0] {
		return false
	}
	
	// In 0.x.y versions, minor version changes can break compatibility
	if currentParts[0] == "0" {
		return currentParts[1] == requiredParts[1]
	}
	
	return true
}
