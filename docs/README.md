# Brainwave Generator Documentation

Welcome to the documentation for the WebAssembly-based Brainwave Generator project. This documentation will help you understand the project's architecture, components, and underlying concepts.

## Table of Contents

### Concepts
- [Understanding Brainwave Entrainment](concepts/brainwave-entrainment.md)
- [Binaural Beats Explained](concepts/binaural-beats.md)
- [Isochronic Tones Explained](concepts/isochronic-tones.md)

### Technical Documentation
- [Project Architecture](api/architecture.md)
- [WebAssembly Integration](api/wasm-integration.md)
- [Audio Processing](api/audio-processing.md)

### Generator Documentation
- [Base Audio Generator](generators/audio-generator.md)
- [Binaural Beat Generator](generators/binaural-generator.md)
- [Isochronic Tone Generator](generators/isochronic-generator.md)

## Quick Start

To use this project in your own application:

1. Import the desired generator package:
```go
import (
    "github.com/user/wasm-brainwave/pkg/binaural"
    "github.com/user/wasm-brainwave/pkg/isochronic"
)
```

2. Create and use a generator:
```go
// For binaural beats
binauralGen := binaural.NewGenerator(432.0, 7.0)
leftSample, rightSample := binauralGen.GenerateSample()

// For isochronic tones
isochronicGen := isochronic.NewGenerator(432.0, 7.0)
sample := isochronicGen.GenerateSample()
```

See the individual documentation pages for more detailed information about each component.
