package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/drhodes/golorem"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Lipsum Name Generator"
	app.Usage = `lipsum-name [args]
Generates a name with between 5-100 letters`
	app.Action = func(c *cli.Context) {
		name := lorem.Word(5, 100)
		fmt.Println(name)
	}

	app.Run(os.Args)

}
