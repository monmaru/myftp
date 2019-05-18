package cmd

import (
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
		Action: func(c *cli.Context) error {
			err := client.Upload(client.Config{
				Address:     c.String("a"),
				Certificate: c.String("cert"),
				SrcDir:      c.String("d"),
				Parallelism: c.Int("p"),
			})
			if err != nil {
				log.Println(err)
			}
			return err
		},
	}
}
