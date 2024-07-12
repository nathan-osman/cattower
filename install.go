package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"text/template"

	"github.com/urfave/cli/v2"
)

const (
	systemdUnitFile = `[Unit]
Description=cattower
Wants=network-online.target
After=network-online.target

[Service]
ExecStart={{.path}} --config {{.config_path}}

[Install]
WantedBy=multi-user.target
`
	configFile = `# TODO: use this file to configure the application

influxdb:
  addr:
  username:
  password:
  database:
`
)

var (
	configFlag = &cli.StringFlag{
		Name:    "config",
		Value:   "/etc/cattower/config.json",
		EnvVars: []string{"CONFIG"},
		Usage:   "filename of configuration file",
	}
	installCommand = &cli.Command{
		Name:   "install",
		Usage:  "install the application as a local service",
		Flags:  []cli.Flag{configFlag},
		Action: install,
	}
)

func writeTemplate(
	filename string,
	content string,
	perm fs.FileMode,
	data any,
) error {
	if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
		return err
	}
	t, err := template.New("").Parse(content)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	defer f.Close()
	return t.Execute(f, data)
}

func install(c *cli.Context) error {

	// Determine the full path to the executable
	p, err := os.Executable()
	if err != nil {
		return err
	}

	// Write the unit file
	if err := writeTemplate(
		"/lib/systemd/system/cattower.service",
		systemdUnitFile,
		0666,
		map[string]interface{}{
			"path":        p,
			"config_path": c.String("config"),
		},
	); err != nil {
		return err
	}

	// Write the configuration file
	if err := writeTemplate(
		c.String("config"),
		configFile,
		0600,
		nil,
	); err != nil {
		return err
	}

	fmt.Println("Service installed!")
	fmt.Println("")
	fmt.Println("Be sure to edit the configuration file:")
	fmt.Println(c.String("config"))
	fmt.Println("")
	fmt.Println("To enable the service and start it, run:")
	fmt.Println("  systemctl enable cattower")
	fmt.Println("  systemctl start cattower")

	return nil
}
