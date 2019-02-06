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

var version = "v0.1.1"

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
				return convert(c, getEncodeReader)
			},
		},
		{
			Name:      "decode",
			Usage:     "Decode garbled Shift-JIS text to UTF-8",
			ArgsUsage: "text",
			Action: func(c *cli.Context) error {
				return convert(c, getDecodeReader)
			},
		},
	}

	app.Run(os.Args)
}

func convert(c *cli.Context, getReader func(string) *transform.Reader) error {
	if c.NArg() != 1 {
		return cli.NewExitError("too many arguments", 1)
	}

	input := c.Args().First()
	r := getReader(input)
	b, e := ioutil.ReadAll(r)
	if 0 < len(b) {
		fmt.Println(string(b))
	}
	if e != nil {
		return cli.NewExitError(e.Error(), 1)
	}
	return nil
}

func getEncodeReader(s string) *transform.Reader {
	return transform.NewReader(strings.NewReader(s), japanese.ShiftJIS.NewDecoder())
}

func getDecodeReader(s string) *transform.Reader {
	return transform.NewReader(strings.NewReader(s), japanese.ShiftJIS.NewEncoder())
}
