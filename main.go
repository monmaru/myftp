package main

import (
	"os"

	"github.com/monmaru/myftp/cmd"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "myftp"
	app.Commands = []cli.Command{
		cmd.Upload(),
		cmd.Serve(),
		cmd.Download(),
		cmd.List(),
	}
	app.Run(os.Args)
}
