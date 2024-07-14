package server

import (
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func (s *Server) apiVideosIndex(c *gin.Context) {
	entries, err := os.ReadDir(s.cfg.VideoDir)
	if err != nil {
		panic(err)
	}
	fileUris := []string{}
	for _, e := range entries {
		if e.Type().IsRegular() {
			absPath, err := filepath.Abs(
				path.Join(s.cfg.VideoDir, e.Name()),
			)
			if err != nil {
				panic(err)
			}
			absPath = filepath.ToSlash(absPath)
			if !strings.HasPrefix(absPath, "/") {
				absPath = "/" + absPath
			}
			fileUris = append(
				fileUris,
				(&url.URL{
					Scheme: "file",
					Path:   absPath,
				}).String(),
			)
		}
	}
	c.JSON(http.StatusOK, fileUris)
}
