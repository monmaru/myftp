package cmd

import (
	"log"

	"github.com/monmaru/myftp/client"
	"github.com/urfave/cli"
)

// Download ...
func Download() cli.Command {
	return cli.Command{
		Name:  "download",
		Usage: "download a file",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "name",
				Usage: "Target file name",
			},
			cli.StringFlag{
				Name:  "d",
				Usage: "Destination directory",
				Value: ".",
			},
			&cli.StringFlag{
				Name:  "a",
				Usage: "Address to listen",
				Value: "localhost:5000",
			},
			&cli.StringFlag{
				Name:  "key",
				Usage: "TLS certificate key",
			},
			&cli.StringFlag{
				Name:  "cert",
				Usage: "path to the TLS *.crt file",
			},
		},
		Action: func(c *cli.Context) error {
			err := client.Download(
				client.Config{
					Address:     c.String("a"),
					Certificate: c.String("cert"),
					SrcDir:      c.String("d"),
					Parallelism: 1,
				}, c.String("name"))

			if err != nil {
				log.Println(err)
				return err
			}
			return nil
		},
	}
}
