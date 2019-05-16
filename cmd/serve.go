package cmd

import (
	"fmt"

	"github.com/monmaru/myftp/server"
	"github.com/urfave/cli"
)

// Serve ...
func Serve() cli.Command {
	return cli.Command{
		Name:  "serve",
		Usage: "initiates a gRPC server",

		Flags: []cli.Flag{
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
				Name:  "certificate",
				Usage: "path to TLS certificate",
			},
			&cli.StringFlag{
				Name:  "d",
				Usage: "Destrination directory",
				Value: "/tmp",
			},
		},
		Action: func(c *cli.Context) error {
			cfg := server.Config{
				Address:     c.String("a"),
				Certificate: c.String("certificate"),
				Key:         c.String("key"),
				DestDir:     c.String("d"),
			}

			close, err := server.Start(cfg)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}

			defer close()
			return nil
		},
	}
}
