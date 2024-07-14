package server

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (s *Server) apiVideosIndex(c *gin.Context) {
	entries, err := os.ReadDir(s.cfg.VideoDir)
	if err != nil {
		panic(err)
	}
	files := []string{}
	for _, e := range entries {
		if e.Type().IsRegular() {
			files = append(files, e.Name())
		}
	}
	c.JSON(http.StatusOK, files)
}
