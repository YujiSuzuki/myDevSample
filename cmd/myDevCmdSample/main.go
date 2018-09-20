package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)

	app := cli.NewApp()

	app.Name = "command 名"
	app.Usage = "説明"
	app.Version = "0.0.1"

	app.Action = func(context *cli.Context) error {
		if context.Bool("cat") {
			fmt.Println(context.Args().Get(0) + "！！Cat！！")
		} else {
			log.Println("⭕")
			fmt.Println(context.Args().Get(0))
			fmt.Println(context.Args().Get(1))
			fmt.Println(context.Args().Get(2))
		}
		return nil
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "cat, c",
			Usage: "Echo with cat",
		},
	}

	app.Run(os.Args)
}
