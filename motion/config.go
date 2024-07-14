package motion

import "time"

type Config struct {
	Pin      uint8         `yaml:"pin"`
	Cooldown time.Duration `yaml:"cooldown"`
	LogSize  int           `yaml:"log_size"`
}
