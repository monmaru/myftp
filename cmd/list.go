package cmd

import (
	"log"

	"github.com/monmaru/myftp/client"
	"github.com/urfave/cli"
)

// List ...
func List() cli.Command {
	return cli.Command{
		Name:  "list",
		Usage: "List files",
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
				Name:  "cert",
				Value: "",
				Usage: "path to the TLS *.crt file",
			},
		},

		Action: func(ctx *cli.Context) error {
			c, err := client.New(
				ctx.String("a"),
				ctx.String("cert"),
			)

			if err != nil {
				return err
			}

			if err := c.List(); err != nil {
				log.Println(err)
				return err
			}

			return nil
		},
	}
}
