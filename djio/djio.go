// !windows

package djio

import (
	gpio "github.com/davecheney/gpio"
	"fmt"
)


func switchPin(p int, v int) {
	pin, err := gpio.OpenPin(p, gpio.ModeOutput)

	if err != nil {
		fmt.Printf("Error opening pin! %s\n", err)
		return
	}

	if v == 1 {
		pin.Set()
	} else {
		pin.Clear()
	}
}

func PinHigh(pin int) {
	switchPin(pin, 1)
}

func PinLow(pin int) {
	switchPin(pin, 0)
}