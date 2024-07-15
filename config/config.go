package config

import (
	"github.com/nathan-osman/cattower/influxdb"
	"github.com/nathan-osman/cattower/leds"
	"github.com/nathan-osman/cattower/motion"
	"github.com/nathan-osman/cattower/server"
)

// Config stores configuration data for the application.
type Config struct {
	InfluxDB influxdb.Config `yaml:"influxdb"`
	Leds     leds.Config     `yaml:"leds"`
	Motion   motion.Config   `yaml:"motion"`
	Server   server.Config   `yaml:"server"`
}
