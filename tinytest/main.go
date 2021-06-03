package main

import "machine"

func main() {
	machine.LED.Configure(machine.PinConfig{Mode: machine.PinOutput})

	b := &Banana{}
	b.Inc()

	for {
	}
}
