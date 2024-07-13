//go:build unix

package hardware

import (
	"image/color"

	"github.com/Jon-Bright/ledctl/pixarray"
)

// Hardware provides access to the hardware on the system.
type Hardware struct {
	ledStrip pixarray.LEDStrip
}

func New() (*Hardware, error) {
	l, err := pixarray.NewWS281x(48, 3, pixarray.GRB, 800000, 10, []int{18})
	if err != nil {
		return nil, err
	}
	return &Hardware{
		ledStrip: l,
	}, nil
}

func (h *Hardware) SetPixel(i int, c color.Color) {
	r, g, b, _ := c.RGBA()
	h.ledStrip.SetPixel(
		i,
		pixarray.Pixel{
			R: int(r),
			G: int(g),
			B: int(b),
		},
	)
}

func (h *Hardware) WritePixels() error {
	return h.ledStrip.Write()
}

func (h *Hardware) Close() {
	//...
}
