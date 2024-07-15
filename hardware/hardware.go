//go:build !unix

package hardware

// Hardware provides a stub for non-unix platforms
type Hardware struct {
}

func New() (*Hardware, error) {
	return &Hardware{}, nil
}

func (h *Hardware) InitPin(pin uint8, direction Direction) {}
func (h *Hardware) ReadPin(pin uint8) bool                 { return false }
func (h *Hardware) Close()                                 {}
