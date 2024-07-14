package motion

import (
	"errors"
	"sync"
	"time"

	"github.com/nathan-osman/cattower/hardware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Event struct {
	Timestamp time.Time `json:"timestamp"`
	Motion    bool      `json:"motion"`
}

type Motion struct {
	mutex      sync.RWMutex
	cfg        *Config
	log        []*Event
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
			newVal := m.hardware.ReadPin(m.cfg.Pin)
			if newVal != lastVal {
				now := time.Now()
				if lastFlip.Add(m.cfg.Cooldown).Before(now) {
					func() {
						m.mutex.Lock()
						defer m.mutex.Unlock()
						m.log = append(m.log, &Event{
							Timestamp: now,
							Motion:    newVal,
						})
						if len(m.log) > m.cfg.LogSize {
							m.log = m.log[1:]
						}
					}()
					lastVal = newVal
					lastFlip = now
				}
			}
		case <-m.closeChan:
			return
		}
	}
}

func New(cfg *Config, h *hardware.Hardware) (*Motion, error) {

	// Sanity check on config
	if cfg.Pin == 0 {
		return nil, errors.New("pin for motion sensor not specified")
	}
	if cfg.Cooldown == 0 {
		cfg.Cooldown = 2 * time.Second
	}
	if cfg.LogSize == 0 {
		cfg.LogSize = 20
	}

	// Initialize the selected pin
	h.InitPin(cfg.Pin, hardware.Input)

	// Create the Motion instance
	m := &Motion{
		cfg:      cfg,
		logger:   log.With().Str("package", "motion").Logger(),
		hardware: h,
	}

	// Start the monitoring goroutine
	go m.run()

	return m, nil
}

func (m *Motion) Log() []*Event {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	ret := make([]*Event, len(m.log))
	copy(ret, m.log)
	return ret
}

func (m *Motion) Close() {
	close(m.closeChan)
	<-m.closedChan
}
