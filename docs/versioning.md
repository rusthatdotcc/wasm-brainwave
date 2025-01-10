# Versioning Guide

This project follows [Semantic Versioning 2.0.0](https://semver.org/). This document explains how we version our releases and what users and contributors can expect from version updates.

## Version Format

We use the format `MAJOR.MINOR.PATCH` (e.g., `1.2.3`):

- **MAJOR** version increments indicate incompatible API changes
- **MINOR** version increments indicate new backwards-compatible functionality
- **PATCH** version increments indicate backwards-compatible bug fixes

## Version Categories

### Major Versions (X.0.0)
Changes that require updates to client code:
- Removing or renaming public APIs
- Changing function signatures
- Modifying core generator behavior
- Updating minimum Go version requirements

### Minor Versions (0.X.0)
New features that don't break existing code:
- Adding new generator types
- Introducing optional parameters
- Adding new utility functions
- Expanding documentation

### Patch Versions (0.0.X)
Bug fixes and small improvements:
- Performance optimizations
- Documentation corrections
- Bug fixes in existing functionality
- Minor code improvements

## Version Guarantees

### API Stability
- Major version 0 (0.y.z): No stability guarantees
- Major version 1+: API stability within major versions
- Experimental features: Tagged with `experimental` in docs

### Compatibility
- Go modules: Compatible with the version specified in `go.mod`
- WebAssembly: Tested with major browsers
- Audio API: Following Web Audio API standards

## Release Process

1. **Preparation**
   ```bash
   # Update CHANGELOG.md
   # Update version numbers in documentation
   # Run tests
   go test ./...
   ```

2. **Version Update**
   ```bash
   # Update go.mod version
   # Commit changes
   git add .
   git commit -m "chore: prepare release vX.Y.Z"
   ```

3. **Tagging**
   ```bash
   # Create annotated tag
   git tag -a vX.Y.Z -m "Release vX.Y.Z"
   ```

## Using Versions

### In Go Projects
```go
module example.com/myapp

require (
    github.com/user/wasm-brainwave v1.2.3
)
```

### WebAssembly Build
```bash
# Always use tagged versions for builds
git checkout v1.2.3
GOOS=js GOARCH=wasm go build -o main.wasm
```

## Version History

See [CHANGELOG.md](../CHANGELOG.md) for detailed version history.

Current stable version: v0.1.0

## Pre-1.0 Development

During pre-1.0 development (0.x.y):
- APIs may change without major version bump
- Focus on gathering feedback and stabilizing interfaces
- Regular minor version updates for new features
- Patch versions for critical fixes

## Long-term Support (LTS)

- Major versions receive security updates for 1 year
- Only latest minor version of each major version supported
- Critical security fixes may be backported to older versions

## Deprecation Policy

1. **Feature Deprecation**
   - Marked as deprecated in documentation
   - Maintained for one major version
   - Removed in next major version

2. **Migration Support**
   - Migration guides provided for major changes
   - Compatibility layers when possible
   - Tools for automated updates when applicable

## Version Checking

### Runtime Version Check
```go
package main

import "wasm-brainwave/version"

func main() {
    // Get current version
    ver := version.Current()
    
    // Check compatibility
    if !version.IsCompatible("1.2.3") {
        log.Fatal("Incompatible version")
    }
}
```

### Build-time Version Info
```go
var (
    Version   = "1.2.3"
    BuildTime = "2025-01-10T16:12:45+01:00"
    GitCommit = "abc123"
)
```

## Contributing

When contributing:
1. Document breaking changes clearly
2. Update CHANGELOG.md
3. Follow semantic versioning rules
4. Include version compatibility in PRs
5. Add migration notes if needed
