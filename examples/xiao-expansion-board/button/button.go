package main

import (
	"machine"
)

var buttonPin = machine.D1
var buttonState = false

func setup() {
	// initialize the LED pin as an output:
	pinMode(machine.LED, machine.PinOutput)
	// initialize the pushbutton pin as an input:
	pinMode(buttonPin, machine.PinInputPullup)

}

func loop() {
	// read the state of the pushbutton value:
	buttonState = buttonPin.Get()

	// check if the pushbutton is pressed. If it is, the buttonState is true:
	if buttonState == true {
		// turn LED on:
		machine.LED.High()
	} else {
		// turn LED off:
		machine.LED.Low()
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
