package server

import (
	"image/color"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	cmdTopOn    = "top_on"
	cmdTopOff   = "top_off"
	cmdSidesOn  = "sides_on"
	cmdSidesOff = "sides_off"
)

type apiSetColorsParams struct {
	Command string `json:"command"`
}

func (s *Server) fillPixels(start, end int, c color.Color) {
	for i := start; i < end; i++ {
		s.hardware.SetPixel(i, c)
	}
}

func (s *Server) apiSetColors(c *gin.Context) {
	v := &apiSetColorsParams{}
	if err := c.ShouldBindJSON(v); err != nil {
		panic(err)
	}
	switch v.Command {
	case cmdTopOn:
		s.fillPixels(32, 48, color.White)
	case cmdTopOff:
		s.fillPixels(32, 48, color.Black)
	case cmdSidesOn:
		s.fillPixels(0, 32, color.White)
	case cmdSidesOff:
		s.fillPixels(0, 32, color.Black)
	default:
		panic("unrecognized command")
	}
	if err := s.hardware.Write(); err != nil {
		panic(err)
	}
	c.Status(http.StatusNoContent)
}
