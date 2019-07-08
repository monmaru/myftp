package cmd

import (
	"fmt"
	"log"

	"github.com/monmaru/myftp/client"
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
				Name:  "cert",
				Value: "",
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

			if err := c.Upload(ctx.String("d"), ctx.Int("p")); err != nil {
				log.Println(err)
				return err
			}

			fmt.Println("")
			fmt.Println(" Upload done!")
			return nil
		},
	}
}
