package config

// Config stores configuration data for the application.
type Config struct {
	InfluxDB struct {
		URL      string `yaml:"url"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	} `yaml:"influxdb"`
}
