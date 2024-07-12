package config

// Config stores configuration data for the application.
type Config struct {
	InfluxDB struct {
		Addr     string `yaml:"addr"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	} `yaml:"influxdb"`
}
