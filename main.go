package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Wack"
	app.Usage = "Easily creates WordPress actions on the front and back ends"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Im in a cli thing")
		return nil
	}

	error := app.Run(os.Args)
	if error != nil {
		log.Fatal("i died")
	}

}

func fileExists(name string) bool {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false
	}

	return true
}
