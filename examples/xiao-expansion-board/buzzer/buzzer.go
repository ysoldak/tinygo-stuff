package main

import (
	"machine"
	"time"
)

var speakerPin = machine.A3
var length = 28 // the number of notes
var notes = "GGAGcB GGAGdc GGxecBA yyecdc"
var beats = []int{2, 2, 8, 8, 8, 16, 1, 2, 2, 8, 8, 8, 16, 1, 2, 2, 8, 8, 8, 8, 16, 1, 2, 2, 8, 8, 8, 16}
var tempo = 150

func playTone(tone int, duration int) {
	for i := 0; i < duration*1000; i += tone * 2 {
		speakerPin.High()
		delayMicroseconds(tone)
		speakerPin.Low()
		delayMicroseconds(tone)
	}
}

func playNote(note byte, duration int) {
	names := "CDEFGABcdefgabxy"
	tones := []int{
		1915, 1700, 1519, 1432, 1275, 1136, 1014,
		956, 834, 765, 593, 468, 346, 224,
		655, 715,
	}
	SPEE := 5

	// play the tone corresponding to the note name

	for i := 0; i < 16; i++ {
		if names[i] == note {
			newduration := duration / SPEE
			playTone(tones[i], newduration)
		}
	}
}

func setup() {
	pinMode(speakerPin, machine.PinOutput)
}

func loop() {
	for i := 0; i < length; i++ {
		if notes[i] == ' ' {
			delay(beats[i] * tempo) // rest
		} else {
			playNote(notes[i], beats[i]*tempo)
		}
		// pause between notes
		delay(tempo)
	}
}

// --- end of line-to-line parity ---

func main() {
	setup()
	for {
		loop()
	}
}

func pinMode(pin machine.Pin, mode machine.PinMode) {
	pin.Configure(machine.PinConfig{Mode: mode})
}

func delay(t int) {
	time.Sleep(time.Duration(t) * time.Millisecond)
}

func delayMicroseconds(t int) {
	time.Sleep(time.Duration(t) * time.Microsecond)
}
