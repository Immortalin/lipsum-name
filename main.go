package main

import (
	"fmt"
	// "github.com/atotto/clipboard"
	"github.com/Immortalin/clipboard"
	// "gopkg.in/urfave/cli.v1"
	"github.com/Immortalin/cli"
	// "github.com/drhodes/golorem"
	"github.com/Immortalin/golorem"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Lorem Ipsum Name Generator"
	app.Version = "0.0.1"
	app.Usage = `lipsum-name - Generates a name with between 5-100 letters`
	app.Action = func(c *cli.Context) error {
		name := lorem.Word(5, 100)
		fmt.Println(name)
		clipboard.WriteAll(name)
		return nil
	}

	app.Run(os.Args)

}
