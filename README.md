# Brainwave Generator

A WebAssembly-powered application for generating binaural beats and isochronic tones, built with Go.

## Features

- Generate binaural beats with customizable frequencies
- Create isochronic tones with adjustable pulse rates
- Real-time audio synthesis
- Web-based interface for easy access

## Development

### Prerequisites

- Go 1.21 or later
- Modern web browser with WebAssembly support

### Building

```bash
# Build the WebAssembly binary
GOOS=js GOARCH=wasm go build -o main.wasm

# Copy the WebAssembly support files
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```

### Running

Serve the directory with a web server and open index.html in your browser.

## License

MIT License
