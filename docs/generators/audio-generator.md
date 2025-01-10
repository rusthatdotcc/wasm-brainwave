# Base Audio Generator

The base audio generator (`pkg/audio/generator.go`) provides the fundamental building blocks for creating audio waveforms. It serves as the foundation for both the binaural beat and isochronic tone generators.

## Technical Overview

### Constants

```go
const SampleRate = 44100
```

The sample rate determines how many samples are generated per second. 44100 Hz is the standard CD-quality audio rate.

### Generator Structure

```go
type Generator struct {
    Phase     float64
    Frequency float64
}
```

- `Phase`: Current phase of the waveform (0.0 to 1.0)
- `Frequency`: Frequency of the waveform in Hz

### Key Methods

#### NewGenerator

```go
func NewGenerator(freq float64) *Generator
```

Creates a new generator instance with the specified frequency.

Example:
```go
gen := audio.NewGenerator(432.0) // Create a 432 Hz generator
```

#### GenerateSample

```go
func (g *Generator) GenerateSample() float64
```

Generates the next sample in the waveform sequence. Uses sine wave generation:
1. Calculates sine value based on current phase
2. Updates phase for next sample
3. Keeps phase between 0 and 1

Example:
```go
sample := gen.GenerateSample() // Get next sample value
```

#### SetFrequency

```go
func (g *Generator) SetFrequency(freq float64)
```

Updates the generator's frequency while maintaining phase continuity.

Example:
```go
gen.SetFrequency(440.0) // Change to 440 Hz
```

#### Reset

```go
func (g *Generator) Reset()
```

Resets the generator's phase to 0, useful when starting a new sequence.

## Implementation Details

### Phase Calculation

The phase increment per sample is calculated as:
```
phase_increment = frequency / sample_rate
```

This ensures accurate frequency reproduction regardless of the sample rate.

### Sine Wave Generation

The sine wave is generated using:
```go
sample := math.Sin(2 * math.Pi * phase)
```

This produces values between -1 and 1, suitable for audio output.

## Usage Examples

### Basic Usage

```go
// Create a 432 Hz generator
gen := audio.NewGenerator(432.0)

// Generate 1 second of audio
for i := 0; i < audio.SampleRate; i++ {
    sample := gen.GenerateSample()
    // Use sample value...
}
```

### Frequency Modulation

```go
gen := audio.NewGenerator(432.0)

// Change frequency over time
for i := 0; i < audio.SampleRate; i++ {
    if i == audio.SampleRate/2 {
        gen.SetFrequency(440.0) // Change frequency halfway through
    }
    sample := gen.GenerateSample()
    // Use sample value...
}
```

## Best Practices

1. **Sample Rate Consistency**
   - Always use the same sample rate throughout your application
   - Don't mix different sample rates without proper conversion

2. **Phase Continuity**
   - Use SetFrequency instead of creating new generators when changing frequency
   - This prevents audio clicks and pops

3. **Value Range**
   - Output values are between -1 and 1
   - Scale appropriately for your audio system

4. **Memory Management**
   - Generators are lightweight and can be created/destroyed frequently
   - Reuse generators when possible for better performance
