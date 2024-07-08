package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/urfave/cli/v2"
)

const (
	systemdUnitFileName = "/lib/systemd/system/cattower.service"
	systemdUnitFileData = `[Unit]
Description=cattower
Wants=network-online.target
After=network-online.target

[Service]
ExecStart={{.path}}

[Install]
WantedBy=multi-user.target
`
)

var installCommand = &cli.Command{
	Name:   "install",
	Usage:  "install the application as a local service",
	Action: install,
}

func install(c *cli.Context) error {

	// Determine the full path to the executable
	p, err := os.Executable()
	if err != nil {
		return err
	}

	// Write the unit file
	if err := os.MkdirAll(filepath.Dir(systemdUnitFileName), 0755); err != nil {
		return err
	}
	t, err := template.New("").Parse(systemdUnitFileData)
	if err != nil {
		return err
	}
	f, err := os.Create(systemdUnitFileName)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := t.Execute(
		f,
		map[string]interface{}{
			"path": p,
		},
	); err != nil {
		return err
	}

	fmt.Println("Service installed!")
	fmt.Println("")
	fmt.Println("To enable the service and start it, run:")
	fmt.Println("  systemctl enable cattower")
	fmt.Println("  systemctl start cattower")

	return nil
}
