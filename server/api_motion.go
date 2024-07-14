package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) apiMotionLog(c *gin.Context) {
	if s.motion == nil {
		panic("motion detection not active")
	}
	c.JSON(http.StatusOK, s.motion.Log())
}
