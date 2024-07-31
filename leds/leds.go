package leds

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/color"
	"net/http"
	"net/url"
	"strconv"
	"sync"
)

type lampctlChange struct {
	GroupID string `json:"group_id"`
	LampID  string `json:"lamp_id"`
	Color   string `json:"color"`
}

type Leds struct {
	mutex   sync.Mutex
	addr    string
	count   int
	changes []*lampctlChange
}

func (l *Leds) post(path string, v any) error {

	// Encode the data
	b := bytes.Buffer{}
	if err := json.NewEncoder(&b).Encode(v); err != nil {
		return err
	}

	// Send the request
	r, err := http.Post(
		(&url.URL{
			Scheme: "http",
			Host:   l.addr,
			Path:   path,
		}).String(),
		"application/json",
		&b,
	)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return nil
}

func New(cfg *Config) *Leds {
	var (
		l = &Leds{
			addr:  cfg.Addr,
			count: cfg.Count,
		}
	)
	for i := 0; i < cfg.Count; i++ {
		l.changes = append(l.changes, &lampctlChange{
			GroupID: "ws2811",
			LampID:  strconv.Itoa(i),
		})
	}
	return l
}

func (l *Leds) SetPixel(pixel int, c color.Color) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if pixel >= 0 && pixel < l.count {
		r, g, b, _ := c.RGBA()
		l.changes[pixel].Color = fmt.Sprintf(
			"%02s%02s%02s",
			strconv.FormatInt(int64(uint8(r/256)), 16),
			strconv.FormatInt(int64(uint8(g/256)), 16),
			strconv.FormatInt(int64(uint8(b/256)), 16),
		)
	}
}

func (l *Leds) Apply() error {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	return l.post(
		"/api/providers/ws2811/apply",
		l.changes,
	)
}
