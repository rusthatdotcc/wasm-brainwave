package audio

import "math"

const (
	// SampleRate is the number of samples per second
	SampleRate = 44100
)

// Generator represents a basic audio waveform generator
type Generator struct {
	Phase     float64
	Frequency float64
}

// NewGenerator creates a new audio generator with the specified frequency
func NewGenerator(freq float64) *Generator {
	return &Generator{
		Frequency: freq,
	}
}

// GenerateSample generates the next sample in the waveform
func (g *Generator) GenerateSample() float64 {
	sample := math.Sin(2 * math.Pi * g.Phase)
	g.Phase += g.Frequency / SampleRate
	if g.Phase >= 1.0 {
		g.Phase -= 1.0
	}
	return sample
}

// SetFrequency updates the generator's frequency
func (g *Generator) SetFrequency(freq float64) {
	g.Frequency = freq
}

// Reset resets the generator's phase
func (g *Generator) Reset() {
	g.Phase = 0
}
