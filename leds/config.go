package leds

type Config struct {
	Addr  string `yaml:"addr"`
	Count int    `yaml:"count"`
}
