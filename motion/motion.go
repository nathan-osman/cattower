package motion

import (
	"time"

	"github.com/nathan-osman/cattower/hardware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	cooldownInterval = 10 * time.Second
)

type Motion struct {
	pin        uint8
	logger     zerolog.Logger
	hardware   *hardware.Hardware
	closeChan  chan any
	closedChan chan any
}

func (m *Motion) run() {
	defer close(m.closedChan)
	defer m.logger.Info().Msg("motion detection stopped")
	m.logger.Info().Msg("motion detection started")
	var (
		lastVal  bool
		lastFlip time.Time
		t        = time.NewTicker(500 * time.Millisecond)
	)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			newVal := m.hardware.ReadPin(m.pin)
			if newVal != lastVal {
				now := time.Now()
				if lastFlip.Add(cooldownInterval).Before(now) {
					// TODO: something when motion is detected
					lastVal = newVal
					lastFlip = now
				}
			}
		case <-m.closeChan:
			return
		}
	}
}

func New(cfg *Config, h *hardware.Hardware) *Motion {
	h.InitPin(cfg.Pin, hardware.Input)
	m := &Motion{
		pin:      cfg.Pin,
		logger:   log.With().Str("package", "motion").Logger(),
		hardware: h,
	}
	go m.run()
	return m
}

func (m *Motion) Close() {
	close(m.closeChan)
	<-m.closedChan
}
