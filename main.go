//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"
	"syscall/js"

	"wasm-brainwave/pkg/binaural"
	"wasm-brainwave/pkg/isochronic"
)

var (
	// Channel to keep the program running
	done = make(chan struct{})
)

func generateBinauralBeat(this js.Value, args []js.Value) interface{} {
	if len(args) < 1 {
		fmt.Println("No buffer provided to generateBinauralBeat")
		return nil
	}

	buffer := args[0]
	leftChannel := buffer.Call("getChannelData", 0)
	rightChannel := buffer.Call("getChannelData", 1)
	
	baseFreq := 432.0  // Base frequency
	beatFreq := 7.0    // Beat frequency

	gen := binaural.NewGenerator(baseFreq, beatFreq)
	bufferLength := buffer.Get("length").Int()
	fmt.Printf("Generating binaural beat: base=%v Hz, beat=%v Hz, length=%v\n", baseFreq, beatFreq, bufferLength)

	// Generate samples
	for i := 0; i < bufferLength; i++ {
		leftSample, rightSample := gen.GenerateSample()
		leftChannel.SetIndex(i, js.ValueOf(leftSample))
		rightChannel.SetIndex(i, js.ValueOf(rightSample))
	}

	fmt.Println("Binaural beat generation complete")
	return nil
}

func generateIsochronicTone(this js.Value, args []js.Value) interface{} {
	if len(args) < 1 {
		fmt.Println("No buffer provided to generateIsochronicTone")
		return nil
	}

	buffer := args[0]
	channel := buffer.Call("getChannelData", 0)
	
	freq := 432.0     // Tone frequency
	pulseRate := 7.0  // Pulse rate

	gen := isochronic.NewGenerator(freq, pulseRate)
	bufferLength := buffer.Get("length").Int()
	fmt.Printf("Generating isochronic tone: freq=%v Hz, pulse=%v Hz, length=%v\n", freq, pulseRate, bufferLength)

	// Generate samples
	for i := 0; i < bufferLength; i++ {
		sample := gen.GenerateSample()
		channel.SetIndex(i, js.ValueOf(sample))
	}

	fmt.Println("Isochronic tone generation complete")
	return nil
}

func main() {
	fmt.Println("WebAssembly module initialized")
	
	// Export functions to JavaScript
	generateBinauralBeatFunc := js.FuncOf(generateBinauralBeat)
	generateIsochronicToneFunc := js.FuncOf(generateIsochronicTone)
	
	js.Global().Set("generateBinauralBeat", generateBinauralBeatFunc)
	js.Global().Set("generateIsochronicTone", generateIsochronicToneFunc)

	// Keep the program running
	fmt.Println("WebAssembly module ready")
	<-done
}
