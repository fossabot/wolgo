package main

import (
	"errors"
	"log"
	"os"

	"github.com/minami14/wolgo/wol"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "wolgo"
	app.Usage = "./cli [target mac address]"
	app.Action = func(c *cli.Context) error {
		if !c.Args().Present() {
			return errors.New("wolgo requires 1 argument. it is mac address")
		}
		err := wol.WakeOnLan(c.Args().Get(0))
		return err
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
