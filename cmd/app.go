package cmd

import (
	"github.com/urfave/cli"
)

const (
	appName  = "Etcd dump tools"
	appUsage = "dump etcd K/V to a file"
	version  = "1.0.0"
)

// NewApp creates a new CLI app.
func NewApp() *cli.App {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Name = appName
	app.Usage = appUsage
	app.Version = version

	app.Commands = []cli.Command{
		dumpCmd(),
		restoreCmd(),
	}

	return app
}
