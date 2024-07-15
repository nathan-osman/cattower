package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/nathan-osman/cattower/config"
	"github.com/nathan-osman/cattower/hardware"
	"github.com/nathan-osman/cattower/influxdb"
	"github.com/nathan-osman/cattower/leds"
	"github.com/nathan-osman/cattower/motion"
	"github.com/nathan-osman/cattower/server"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

func main() {
	app := &cli.App{
		Name:  "cattower",
		Usage: "Web application for controlling the cat tower",
		Flags: []cli.Flag{
			configFlag,
		},
		Commands: []*cli.Command{
			installCommand,
		},
		Action: func(c *cli.Context) error {

			// Load the configuration file
			f, err := os.Open(c.String("config"))
			if err != nil {
				return err
			}
			defer f.Close()

			// Parse the configuration file
			cfg := &config.Config{}
			if err := yaml.NewDecoder(f).Decode(cfg); err != nil {
				return err
			}

			// Init the hardware
			h, err := hardware.New()
			if err != nil {
				return err
			}
			defer h.Close()

			// Initialize the connection to the LEDs
			l := leds.New(&cfg.Leds)

			// Init InfluxDB
			i, err := influxdb.New(&cfg.InfluxDB)
			if err != nil {
				return err
			}
			defer i.Close()

			// Init motion detection
			m, err := motion.New(&cfg.Motion, h)
			if err != nil {
				fmt.Fprintf(os.Stderr, "warning: %s\n", err.Error())
			} else {
				defer m.Close()
			}

			// Create the server
			s, err := server.New(&cfg.Server, h, i, l, m)
			if err != nil {
				return err
			}
			defer s.Close()

			// Wait for SIGINT or SIGTERM
			sigChan := make(chan os.Signal, 1)
			signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
			<-sigChan

			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "fatal: %s\n", err.Error())
	}
}
