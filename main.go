package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/nathan-osman/cattower/hardware"
	"github.com/nathan-osman/cattower/server"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "quicksend",
		Usage: "Simple web-based application for sending test emails",
		Commands: []*cli.Command{
			installCommand,
		},
		Action: func(c *cli.Context) error {

			// Init the hardware
			h, err := hardware.New()
			if err != nil {
				return err
			}
			defer h.Close()

			// Create the server
			s, err := server.New(h)
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
