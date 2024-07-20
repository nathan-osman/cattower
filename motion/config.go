package motion

import "time"

type Config struct {
	DetectPin uint8         `yaml:"detect_pin"`
	AlertPin  uint8         `yaml:"alert_pin"`
	Cooldown  time.Duration `yaml:"cooldown"`
	LogSize   int           `yaml:"log_size"`
}
