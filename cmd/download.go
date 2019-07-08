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
				Name:  "a",
				Usage: "Server address",
				Value: "localhost:5000",
			},
			cli.StringFlag{
				Name:  "d",
				Usage: "Destination directory",
				Value: ".",
			},
			&cli.StringFlag{
				Name:  "cert",
				Usage: "path to the TLS *.crt file",
			},
			cli.IntFlag{
				Name:  "p",
				Value: 5,
				Usage: "num of goroutines",
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

			if err := c.Download(ctx.String("d"), ctx.Int("p")); err != nil {
				log.Println(err)
				return err
			}

			return nil
		},
	}
}
