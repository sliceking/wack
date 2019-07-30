package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	var phpFile string
	var jsFile string

	app := cli.NewApp()
	app.Name = "Wack"
	app.Usage = "Easily creates WordPress actions on the front and back ends"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "php, p",
			Usage:       "Append php to `FILE`",
			Destination: &phpFile,
		},
		cli.StringFlag{
			Name:        "js, j",
			Usage:       "Append js to `FILE`",
			Destination: &jsFile,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "php",
			Usage:  "adds php boilerplate to functions.php",
			Action: php,
		},
		{
			Name:   "js",
			Usage:  "adds js boilerplate to a specified js file",
			Action: js,
		},
	}

	error := app.Run(os.Args)
	if error != nil {
		log.Fatal("i died")
	}
}

func php(c *cli.Context) error {
	file, err := os.OpenFile("functions.php", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal("something went wrong opening functions.php")
	}

	defer file.Close()

	boilerplate := fmt.Sprintf(`
		// nonce needs to be localized to the js file
		$ajax_nonce = wp_create_nonce( 'nonce' )

		add_action( 'wp_ajax_%[1]s', '%[1]s' );
		function %[1]s(){
			check_ajax_referer( 'nonce', 'nonce' );
			// do some work
			wp_send_json_success();
		}

		`, c.Args().Get(0))

	file.WriteString(boilerplate)

	return nil
}

func js(c *cli.Context) error {
	jsFile := c.Args().Get(0)

	file, err := os.OpenFile(jsFile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal("something went wrong opening js file")
	}

	defer file.Close()

	boilerplate := fmt.Sprintf(`

	var data = {
		'nonce': nonce,
		'action': '%[1]s',
	};

	$.ajax({
		method: 'POST',
		dataType: 'json',
		url: ajax_url,
		data: data,
		success: function (response) {
			// do some work
		},
		error: function () {
			// something went wrong
		}
	}).always(function () {
		// catch all behavior
	});

	`, c.Args().Get(1))

	file.WriteString(boilerplate)

	return nil
}
