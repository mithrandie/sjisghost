package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/urfave/cli"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

var version = "v0.1.0"

func main() {
	app := cli.NewApp()

	app.Name = "SJIS Ghost"
	app.Usage = "Convert garbled text encoded with Shift-JIS"
	app.Version = version

	app.Commands = []cli.Command{
		{
			Name:      "encode",
			Usage:     "Encode UTF-8 text to garbled Shift-JIS text",
			ArgsUsage: "text",
			Action: func(c *cli.Context) error {
				if c.NArg() != 1 {
					return cli.NewExitError("too many arguments", 1)
				}

				input := c.Args().First()
				s := encode(input)
				fmt.Println(s)
				return nil
			},
		},
		{
			Name:      "decode",
			Usage:     "Decode garbled Shift-JIS text to UTF-8",
			ArgsUsage: "text",
			Action: func(c *cli.Context) error {
				if c.NArg() != 1 {
					return cli.NewExitError("too many arguments", 1)
				}

				input := c.Args().First()
				s := decode(input)
				fmt.Println(s)
				return nil
			},
		},
	}

	app.Run(os.Args)
}

func encode(s string) string {
	r := transform.NewReader(strings.NewReader(s), japanese.ShiftJIS.NewDecoder())
	b, _ := ioutil.ReadAll(r)
	return string(b)
}

func decode(s string) string {
	r := transform.NewReader(strings.NewReader(s), japanese.ShiftJIS.NewEncoder())
	b, _ := ioutil.ReadAll(r)
	return string(b)
}
