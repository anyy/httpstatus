package main

import (
	"encoding/json"
	"fmt"
	"github.com/gazelle0130/httpstatus/status"
	"github.com/urfave/cli"
	"golang.org/x/xerrors"
	"log"
	"os"
	"strconv"
)

type StatusCode struct {
	Code        int    `json:"code"`
	Explanation string `json:"message"`
}

var statuses []StatusCode

func main() {
	app := cli.NewApp()
	app.Name = "httpstatus"
	app.Usage = "show http status code"
	app.Version = "1.0.0"

	app.Before = func(c *cli.Context) error {
		bytes := ([]byte)(status.JsonStr)
		if err := json.Unmarshal(bytes, &statuses); err != nil {
			return err
		}
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "all",
			Aliases: []string{"l"},
			Usage:   "print all of http status code",
			Action: func(c *cli.Context) error {
				return all()
			},
		},
		{
			Name:    "find",
			Aliases: []string{"s"},
			Usage:   "find by code",
			Action: func(c *cli.Context) error {
				return find(c.Args().First())
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func find(code string) (err error) {
	c, _ := strconv.Atoi(code)
	for _, s := range statuses {
		if c == s.Code {
			fmt.Printf("%d : %s\n", s.Code, s.Explanation)
			return
		}
	}
	return xerrors.New("input code not found")
}

func all() (err error) {
	for _, s := range statuses {
		fmt.Printf("%d : %s\n", s.Code, s.Explanation)
	}
	return
}
