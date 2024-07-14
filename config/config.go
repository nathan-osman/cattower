package config

import (
	"github.com/nathan-osman/cattower/influxdb"
	"github.com/nathan-osman/cattower/motion"
)

// Config stores configuration data for the application.
type Config struct {
	InfluxDB *influxdb.Config `yaml:"influxdb"`
	Motion   *motion.Config   `yaml:"motion"`
}
