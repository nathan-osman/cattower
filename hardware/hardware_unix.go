//go:build unix

package hardware

import (
	"github.com/stianeikeland/go-rpio/v4"
)

// Hardware provides access to the hardware on the system.
type Hardware struct{}

func New() (*Hardware, error) {
	if err := rpio.Open(); err != nil {
		return nil, err
	}
	return &Hardware{}, nil
}

func (h *Hardware) InitPin(pin uint8, direction Direction) {
	switch direction {
	case Input:
		rpio.Pin(pin).Input()
	case Output:
		rpio.Pin(pin).Output()
	}
}

func (h *Hardware) ReadPin(pin uint8) bool {
	return rpio.Pin(pin).Read() == rpio.High
}

func (h *Hardware) WritePin(pin uint8, v bool) {
	if v {
		rpio.Pin(pin).Write(rpio.High)
	} else {
		rpio.Pin(pin).Write(rpio.Low)
	}
}

func (h *Hardware) Close() {
	rpio.Close()
}
