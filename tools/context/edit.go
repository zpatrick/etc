package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

func Edit(c *cli.Context) error {
	context := c.Args().Get(0)
	if context == "" {
		return fmt.Errorf("Argument CONTEXT is required")
	}

	dir := c.GlobalString("dir")
	path := fmt.Sprintf("%s/%s", dir, context)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		file.Close()
	}

	editor := c.String("editor")
	cmd := exec.Command(editor, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
