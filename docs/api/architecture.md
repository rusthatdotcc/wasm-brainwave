# Project Architecture

This document describes the overall architecture of the Brainwave Generator project, explaining how different components interact and the design decisions behind them.

## Project Structure

```
wasm-brainwave/
├── pkg/
│   ├── audio/        # Base audio generation
│   ├── binaural/     # Binaural beat generation
│   └── isochronic/   # Isochronic tone generation
├── docs/
│   ├── api/          # Technical documentation
│   ├── concepts/     # Theoretical background
│   └── generators/   # Generator documentation
├── main.go           # WebAssembly entry point
└── web/
    ├── index.html    # Web interface
    └── wasm_exec.js  # WebAssembly support
```

## Component Overview

### 1. Core Audio Generation (pkg/audio)

The foundation of all audio processing, providing:
- Basic waveform generation
- Sample rate management
- Phase tracking
- Frequency control

### 2. Specialized Generators

#### Binaural Beat Generator (pkg/binaural)
- Manages two audio generators
- Coordinates left/right channel frequencies
- Handles beat frequency calculations

#### Isochronic Tone Generator (pkg/isochronic)
- Single tone generation
- Amplitude modulation
- Pulse rate control

### 3. WebAssembly Interface (main.go)

Bridges the Go code with JavaScript:
- Exports generator functions
- Handles audio buffer management
- Provides error handling
- Manages WebAssembly lifecycle

### 4. Web Interface (web/)

Provides user interaction:
- Frequency control
- Generator selection
- Real-time parameter adjustment
- Audio playback control

## Data Flow

1. **User Input**
   ```
   Web Interface → WebAssembly → Generator Configuration
   ```

2. **Audio Generation**
   ```
   Generator → Samples → WebAssembly → AudioContext → Speakers
   ```

## Design Principles

### 1. Modularity
- Each generator is independent
- Common functionality in base package
- Clear interfaces between components

### 2. Reusability
- Generators can be used standalone
- No hard dependencies between packages
- Clean API design

### 3. Performance
- Efficient sample generation
- Minimal memory allocation
- WebAssembly optimization

### 4. Maintainability
- Clear separation of concerns
- Well-documented code
- Consistent coding style

## WebAssembly Integration

### Memory Management
- Audio buffers shared between JS and Go
- Careful handling of garbage collection
- Proper cleanup of resources

### JavaScript Bridge
- Type conversion handling
- Error propagation
- Async operation management

## Future Extensibility

The architecture supports easy addition of:
1. New generator types
2. Additional audio effects
3. Different output formats
4. Alternative user interfaces

## Performance Considerations

### Audio Processing
- Sample generation optimized for real-time
- Minimal allocations in audio path
- Efficient buffer management

### WebAssembly
- Careful memory management
- Optimized data transfer
- Efficient type conversions

## Error Handling

1. **Generator Level**
   - Parameter validation
   - Resource management
   - State consistency

2. **WebAssembly Level**
   - Type conversion errors
   - Memory allocation errors
   - JavaScript interop errors

3. **User Interface Level**
   - Input validation
   - Error display
   - Graceful degradation
