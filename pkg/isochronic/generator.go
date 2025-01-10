package isochronic

import (
	"math"

	"wasm-brainwave/pkg/audio"
)

// Generator represents an isochronic tone generator
type Generator struct {
	toneGen   *audio.Generator
	toneFreq  float64
	pulseRate float64
	time      float64
}

// NewGenerator creates a new isochronic tone generator
func NewGenerator(toneFreq, pulseRate float64) *Generator {
	return &Generator{
		toneGen:   audio.NewGenerator(toneFreq),
		toneFreq:  toneFreq,
		pulseRate: pulseRate,
	}
}

// GenerateSample generates the next amplitude-modulated sample
func (g *Generator) GenerateSample() float64 {
	// Generate the base tone
	sample := g.toneGen.GenerateSample()
	
	// Apply amplitude modulation
	modulation := 0.5 * (1 + math.Sin(2*math.Pi*g.pulseRate*g.time))
	g.time += 1.0 / audio.SampleRate
	
	return sample * modulation
}

// SetFrequencies updates both the tone frequency and pulse rate
func (g *Generator) SetFrequencies(toneFreq, pulseRate float64) {
	g.toneFreq = toneFreq
	g.pulseRate = pulseRate
	g.toneGen.SetFrequency(toneFreq)
}

// Reset resets the generator
func (g *Generator) Reset() {
	g.toneGen.Reset()
	g.time = 0
}

// GetFrequencies returns the current tone frequency and pulse rate
func (g *Generator) GetFrequencies() (float64, float64) {
	return g.toneFreq, g.pulseRate
}
