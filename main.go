package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/codegangsta/cli"
	"github.com/drhodes/golorem"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Lorem Ipsum Name Generator"
	app.Version = "0.0.1"
	app.Usage = `lipsum-name - Generates a name with between 5-100 letters`
	app.Action = func(c *cli.Context) {
		name := lorem.Word(5, 100)
		fmt.Println(name)
		clipboard.WriteAll(name)
	}

	app.Run(os.Args)

}
