# Binaural Beat Generator

The binaural beat generator (`pkg/binaural/generator.go`) creates binaural beats by managing two audio generators that produce slightly different frequencies for each ear.

## Technical Overview

### Generator Structure

```go
type Generator struct {
    leftGen  *audio.Generator
    rightGen *audio.Generator
    baseFreq float64
    beatFreq float64
}
```

- `leftGen`: Generator for the left ear channel
- `rightGen`: Generator for the right ear channel
- `baseFreq`: Base frequency (carrier wave)
- `beatFreq`: Beat frequency (difference between channels)

### Key Methods

#### NewGenerator

```go
func NewGenerator(baseFreq, beatFreq float64) *Generator
```

Creates a new binaural beat generator with specified frequencies.

Example:
```go
gen := binaural.NewGenerator(432.0, 7.0) // 432 Hz carrier, 7 Hz beat
```

#### GenerateSample

```go
func (g *Generator) GenerateSample() (float64, float64)
```

Generates the next stereo sample pair:
- Left channel at baseFreq
- Right channel at baseFreq + beatFreq

Example:
```go
leftSample, rightSample := gen.GenerateSample()
```

#### SetFrequencies

```go
func (g *Generator) SetFrequencies(baseFreq, beatFreq float64)
```

Updates both the base and beat frequencies while maintaining phase continuity.

Example:
```go
gen.SetFrequencies(440.0, 10.0) // Change to 440 Hz carrier, 10 Hz beat
```

## Implementation Details

### Frequency Calculation

The two channels are generated at:
- Left channel: baseFreq
- Right channel: baseFreq + beatFreq

This creates a perceived beat frequency equal to the difference between the channels.

### Phase Management

Each channel maintains its own phase, ensuring:
- Continuous waveforms
- Stable beat frequency
- No audio artifacts when changing frequencies

## Usage Examples

### Basic Usage

```go
// Create a generator for alpha wave (10 Hz) using 432 Hz carrier
gen := binaural.NewGenerator(432.0, 10.0)

// Generate stereo samples
for i := 0; i < audio.SampleRate; i++ {
    left, right := gen.GenerateSample()
    // Use samples in audio buffer...
}
```

### Frequency Modulation

```go
gen := binaural.NewGenerator(432.0, 7.0)

// Gradually change beat frequency
for i := 0; i < audio.SampleRate; i++ {
    if i % (audio.SampleRate/10) == 0 {
        currentBase, currentBeat := gen.GetFrequencies()
        gen.SetFrequencies(currentBase, currentBeat + 0.1)
    }
    left, right := gen.GenerateSample()
    // Use samples...
}
```

## Best Practices

### 1. Frequency Selection

- Base frequency: 300-600 Hz works best for most people
- Beat frequency: 0.5-40 Hz depending on desired brain state
- Keep beat frequency stable for at least 30 seconds

### 2. Audio Output

- Always use stereo output
- Requires headphones for proper effect
- Maintain equal volume in both channels

### 3. Frequency Changes

- Change frequencies gradually
- Use SetFrequencies for smooth transitions
- Allow time for brain to synchronize

### 4. Safety

- Start with higher beat frequencies (>10 Hz)
- Gradually move to lower frequencies
- Monitor user comfort and response

## Common Applications

1. **Meditation**
   - Theta waves (4-8 Hz)
   - Alpha waves (8-13 Hz)

2. **Focus**
   - Beta waves (13-30 Hz)
   - Low gamma waves (30-40 Hz)

3. **Relaxation**
   - Alpha waves (8-13 Hz)
   - High theta waves (6-8 Hz)

4. **Sleep**
   - Delta waves (0.5-4 Hz)
   - Low theta waves (4-6 Hz)
