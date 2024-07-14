package influxdb

import (
	"github.com/influxdata/influxdb/client/v2"
)

type InfluxDB struct {
	client   client.Client
	database string
}

func New(cfg *Config) (*InfluxDB, error) {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     cfg.Addr,
		Username: cfg.Username,
		Password: cfg.Password,
	})
	if err != nil {
		return nil, err
	}
	return &InfluxDB{
		client:   c,
		database: cfg.Database,
	}, nil
}

func (i *InfluxDB) Query(q string) (*client.Response, error) {
	return i.client.Query(client.Query{
		Command:  q,
		Database: i.database,
	})
}

func (i *InfluxDB) Close() {
	i.client.Close()
}
