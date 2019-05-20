package cmd

import (
	"log"

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
				Name:  "cert",
				Usage: "path to the TLS *.crt file",
			},
			&cli.StringFlag{
				Name:  "d",
				Usage: "Destrination directory",
				Value: "/tmp",
			},
			&cli.StringFlag{
				Name:  "log",
				Usage: "Log directory",
				Value: ".",
			},
		},
		Action: func(c *cli.Context) error {
			cfg := server.Config{
				Address:     c.String("a"),
				Certificate: c.String("cert"),
				Key:         c.String("key"),
				DestDir:     c.String("d"),
				LogDir:      c.String("log"),
			}

			close, err := server.Listen(cfg)
			if err != nil {
				log.Println(err)
				return err
			}

			defer close()
			return nil
		},
	}
}
