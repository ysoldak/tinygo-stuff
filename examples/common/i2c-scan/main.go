// I2C Scanner for TinyGo
// Inspired by https://playground.arduino.cc/Main/I2cScanner/
//
// Algorithm
// 1. Send I2C Start condition
// 2. Send a single byte representing the address, and get the ACK/NAK
// 3. Send the stop condition.
// https://electronics.stackexchange.com/a/76620
//
// Learn more about I2C
// https://learn.sparkfun.com/tutorials/i2c/all

package main

import (
	"fmt"
	"machine"
	"time"
)

func main() {

	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_100KHZ,
	})

	// Wait for user to open serial console
	for !machine.UART0.DTR() {
		time.Sleep(100 * time.Microsecond)
	}

	w, r := []byte{}, []byte{0}
	nDevices := 0

	println("Scanning...")
	for address := uint16(1); address < 127; address++ {
		if err := machine.I2C0.Tx(address, w, r); err == nil { // try read a byte from current address
			fmt.Printf("I2C device found at address %#X !\n", address)
			nDevices++
		}
	}

	if nDevices == 0 {
		println("No I2C devices found")
	} else {
		println("Done")
	}

	// procrastinate for an hour to ensure everything was printed out and board does not die
	time.Sleep(1 * time.Hour)

}
