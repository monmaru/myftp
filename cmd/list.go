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

		Action: func(c *cli.Context) error {
			err := client.List(client.Config{
				Address:     c.String("a"),
				Certificate: c.String("cert"),
				SrcDir:      c.String("d"),
				Parallelism: 1,
			})

			if err != nil {
				log.Println(err)
				return err
			}
			return nil
		},
	}
}
