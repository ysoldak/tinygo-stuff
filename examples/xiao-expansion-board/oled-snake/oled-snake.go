package main

import (
	"machine"

	"image/color"
	"time"

	"tinygo.org/x/drivers/ssd1306"
)

type point struct {
	x, y, dx, dy int16
}

// head and tail points
var ht = []point{
	{0, 0, 1, 1}, {0, 0, 1, 1},
}

func move(p *point, display ssd1306.Device) {
	pixel := display.GetPixel(p.x, p.y)
	c := color.RGBA{255, 255, 255, 255}
	if pixel {
		c = color.RGBA{0, 0, 0, 255}
	}
	display.SetPixel(p.x, p.y, c)
	display.Display()

	p.x += p.dx
	p.y += p.dy

	if p.x == 0 || p.x == 127 {
		p.dx = -p.dx
	}

	if p.y == 0 || p.y == 63 {
		p.dy = -p.dy
	}
}

func main() {
	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_400KHZ,
	})

	display := ssd1306.NewI2C(machine.I2C0)
	display.Configure(ssd1306.Config{
		Address: 0x3C,
		Width:   128,
		Height:  64,
	})

	display.ClearDisplay()

	length := 20
	i := 0
	for {
		move(&ht[0], display)
		if i > length {
			move(&ht[1], display)
		}
		// if i < 200 {
		// 	display.ClearDisplay()
		// }
		i++
		time.Sleep(1 * time.Millisecond)
	}
}
