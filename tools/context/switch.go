package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/urfave/cli"
)

func Switch(c *cli.Context) error {
	context := c.Args().Get(0)
	if context == "" {
		return fmt.Errorf("Argument CONTEXT is required")
	}

	dir := c.GlobalString("dir")
	path := fmt.Sprintf("%s/%s", dir, context)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("Context '%s' does not exist", context)
		}

		return err
	}

	fmt.Printf(string(content))
	return nil
}
