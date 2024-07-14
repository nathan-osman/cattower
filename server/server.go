package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/nathan-osman/cattower/hardware"
	"github.com/nathan-osman/cattower/influxdb"
	"github.com/nathan-osman/cattower/motion"
	"github.com/nathan-osman/cattower/ui"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Server struct {
	cfg      *Config
	server   http.Server
	logger   zerolog.Logger
	hardware *hardware.Hardware
	influxdb *influxdb.InfluxDB
	motion   *motion.Motion
}

func init() {
	// Switch Gin to release mode
	gin.SetMode(gin.ReleaseMode)
}

func New(
	cfg *Config,
	h *hardware.Hardware,
	i *influxdb.InfluxDB,
	m *motion.Motion,
) (*Server, error) {

	// Initialize the server
	var (
		r = gin.New()
		s = &Server{
			cfg: cfg,
			server: http.Server{
				Addr:    ":8000",
				Handler: r,
			},
			logger:   log.With().Str("package", "server").Logger(),
			hardware: h,
			influxdb: i,
			motion:   m,
		}
	)

	// Serve the video directory from /fs/videos (if provided)
	if cfg.VideoDir != "" {
		r.Use(
			static.Serve(
				"/fs/videos/",
				static.LocalFile(cfg.VideoDir, false),
			),
		)
	}

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

		groupApi.POST("/leds/set-colors", s.apiLedsSetColors)
		groupApi.GET("/motion/log", s.apiMotionLog)
		groupApi.GET("/sensors/overview", s.apiSensorsOverview)
		groupApi.GET("/videos", s.apiVideosIndex)
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
