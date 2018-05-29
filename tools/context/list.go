package main

import (
	"fmt"
	"io/ioutil"

	"github.com/urfave/cli"
)

func List(c *cli.Context) error {
	dir := c.GlobalString("dir")
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() {
			fmt.Println(file.Name())
		}
	}

	return nil
}
