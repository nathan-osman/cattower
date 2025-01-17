package server

import (
	"image/color"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/icza/gox/imagex/colorx"
)

const (
	cmdTopOn    = "top_on"
	cmdTopOff   = "top_off"
	cmdSidesOn  = "sides_on"
	cmdSidesOff = "sides_off"
)

type apiLedsSetColorsParams struct {
	Command string `json:"command"`
	Color   string `json:"color"`
}

func (s *Server) fillPixels(start, end int, c color.Color) {
	for i := start; i < end; i++ {
		s.leds.SetPixel(i, c)
	}
}

func (s *Server) apiLedsSetColors(c *gin.Context) {
	v := &apiLedsSetColorsParams{}
	if err := c.ShouldBindJSON(v); err != nil {
		panic(err)
	}
	var onColor color.Color = color.White
	if len(v.Color) > 0 {
		if c, err := colorx.ParseHexColor(v.Color); err == nil {
			onColor = c
		}
	}
	switch v.Command {
	case cmdTopOn:
		s.fillPixels(32, 48, onColor)
	case cmdTopOff:
		s.fillPixels(32, 48, color.Black)
	case cmdSidesOn:
		s.fillPixels(0, 32, onColor)
	case cmdSidesOff:
		s.fillPixels(0, 32, color.Black)
	default:
		panic("unrecognized command")
	}
	if err := s.leds.Apply(); err != nil {
		panic(err)
	}
	c.Status(http.StatusNoContent)
}
