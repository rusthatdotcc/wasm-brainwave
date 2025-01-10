package binaural

import (
	"wasm-brainwave/pkg/audio"
)

// Generator represents a binaural beat generator
type Generator struct {
	leftGen  *audio.Generator
	rightGen *audio.Generator
	baseFreq float64
	beatFreq float64
}

// NewGenerator creates a new binaural beat generator
func NewGenerator(baseFreq, beatFreq float64) *Generator {
	return &Generator{
		leftGen:  audio.NewGenerator(baseFreq),
		rightGen: audio.NewGenerator(baseFreq + beatFreq),
		baseFreq: baseFreq,
		beatFreq: beatFreq,
	}
}

// GenerateSample generates the next stereo sample pair
func (g *Generator) GenerateSample() (float64, float64) {
	return g.leftGen.GenerateSample(), g.rightGen.GenerateSample()
}

// SetFrequencies updates both the base and beat frequencies
func (g *Generator) SetFrequencies(baseFreq, beatFreq float64) {
	g.baseFreq = baseFreq
	g.beatFreq = beatFreq
	g.leftGen.SetFrequency(baseFreq)
	g.rightGen.SetFrequency(baseFreq + beatFreq)
}

// Reset resets both generators
func (g *Generator) Reset() {
	g.leftGen.Reset()
	g.rightGen.Reset()
}

// GetFrequencies returns the current base and beat frequencies
func (g *Generator) GetFrequencies() (float64, float64) {
	return g.baseFreq, g.beatFreq
}
