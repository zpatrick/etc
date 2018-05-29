package main

import (
	"os"

	"github.com/urfave/cli"
)

func Before(c *cli.Context) error {
	dir := c.GlobalString("dir")
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	return nil
}
