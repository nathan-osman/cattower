package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/nathan-osman/cattower/config"
	"github.com/nathan-osman/cattower/hardware"
	"github.com/nathan-osman/cattower/influxdb"
	"github.com/nathan-osman/cattower/ui"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Server struct {
	server   http.Server
	logger   zerolog.Logger
	cfg      *config.Config
	hardware *hardware.Hardware
	influxdb *influxdb.InfluxDB
}

func init() {
	// Switch Gin to release mode
	gin.SetMode(gin.ReleaseMode)
}

func New(
	h *hardware.Hardware,
	i *influxdb.InfluxDB,
) (*Server, error) {

	// Initialize the server
	var (
		r = gin.New()
		s = &Server{
			server: http.Server{
				Addr:    ":8000",
				Handler: r,
			},
			logger:   log.With().Str("package", "server").Logger(),
			hardware: h,
			influxdb: i,
		}
	)

	// Serve the static files from /
	r.Use(
		static.Serve(
			"/",
			ui.EmbedFileSystem{
				FileSystem: http.FS(ui.Content),
			},
		),
	)

	groupApi := r.Group("/api")
	{
		// Use the session and our custom user middleware for the API
		groupApi.Use(
			gin.CustomRecoveryWithWriter(nil, panicToJSONError),
		)

		groupApi.POST("/set-colors", s.apiSetColors)
		groupApi.GET("/get-sensors", s.apiGetSensors)
	}

	// Serve the static files on all other paths too
	r.NoRoute(func(c *gin.Context) {
		c.Request.URL.Path = "/"
		r.HandleContext(c)
		c.Abort()
	})

	// Listen for connections in a separate goroutine
	go func() {
		defer s.logger.Info().Msg("server stopped")
		s.logger.Info().Msg("server started")
		if err := s.server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			s.logger.Error().Msg(err.Error())
		}
	}()

	return s, nil
}

func (s *Server) Close() {
	s.server.Shutdown(context.Background())
}
