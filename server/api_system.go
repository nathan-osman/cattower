package server

import (
	"bytes"
	"net/http"
	"os/exec"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

// pactl set-sink-volume 0 10%

var (
	getSinkVolumeRe = regexp.MustCompile(`\d+%`)
)

func (s *Server) apiSystemVolume(c *gin.Context) {
	var (
		cmd = exec.Command(
			"pactl",
			"get-sink-volume",
			"0",
		)
		out = bytes.Buffer{}
	)
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	m := getSinkVolumeRe.Find(out.Bytes())
	if m == nil {
		panic("unable to retrieve volume")
	}
	v, err := strconv.Atoi(string(m[:len(m)-1]))
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"volume": v,
	})
}

type apiSystemVolumeSetParams struct {
	Volume int `json:"volume"`
}

func (s *Server) apiSystemVolumeSet(c *gin.Context) {
	v := &apiSystemVolumeSetParams{}
	if err := c.ShouldBindJSON(v); err != nil {
		panic(err)
	}
	if err := exec.Command(
		"pactl",
		"set-sink-volume",
		"0",
		strconv.Itoa(v.Volume)+"%",
	).Run(); err != nil {
		panic(err)
	}
	c.Status(http.StatusNoContent)
}
