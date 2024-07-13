//go:build !unix

package hardware

import (
	"image/color"
)

// Hardware provides a stub for non-unix platforms
type Hardware struct {
}

func New() (*Hardware, error) {
	return &Hardware{}, nil
}

func (h *Hardware) SetPixel(i int, c color.Color) {}
func (h *Hardware) WritePixels() error            { return nil }
func (h *Hardware) Close()                        {}
