package cmd

import (
	"github.com/urfave/cli"
)

// Upload ..
func Upload() cli.Command {
	return cli.Command{
		Name:  "upload",
		Usage: "Upload files in parallel",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "a",
				Value: "localhost:5000",
				Usage: "server address",
			},
			cli.StringFlag{
				Name:  "d",
				Value: ".",
				Usage: "base directory",
			},
			cli.StringFlag{
				Name:  "certificate",
				Value: "",
				Usage: "directory to the TLS server.crt file",
			},
		},
		Action: func(c *cli.Context) error {
			return nil
		},
	}
}
