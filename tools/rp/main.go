package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/urfave/cli"
)

var Version string

func main() {
	if Version == "" {
		Version = "unset/develop"
	}

	app := cli.NewApp()
	app.Name = "Replace"
	app.Version = Version
	app.ArgsUsage = "PATTERN REPLACEMENT TARGET"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "write, w",
			Usage: "write changes in source files",
		},
		cli.BoolFlag{
			Name:  "interactive, i",
			Usage: "prompt before replacing",
		},
		cli.BoolFlag{
			Name:  "recursive, r",
			Usage: "recursively search files",
		},
		cli.StringFlag{
			Name:  "dir, d",
			Usage: "directory to start search",
			Value: ".",
		},
	}

	app.Action = func(c *cli.Context) error {
		args := c.Args()
		pattern := args.Get(0)
		if pattern == "" {
			return fmt.Errorf("Argument PATTERN is required")
		}

		replacement := args.Get(1)
		if replacement == "" {
			return fmt.Errorf("Argument REPLACEMENT is required")
		}

		target := args.Get(2)
		if target == "" {
			return fmt.Errorf("Argument TARGET is required")
		}

		replacer := NewPatternByteReplacer([]byte(pattern), []byte(replacement))
		if c.Bool("interactive") {
			replacer = NewInteractiveReplacer(os.Stdin, os.Stdout, replacer)
		}

		filePaths, err := GlobMatchFiles(c.String("dir"), target, c.Bool("recursive"))
		if err != nil {
			return fmt.Errorf("Failed to match files: %s", err.Error())
		}

		for _, filePath := range filePaths {
			input, err := ioutil.ReadFile(filePath)
			if err != nil {
				return err
			}

			fmt.Println(filePath)
			r := bytes.NewReader(input)

			if !c.Bool("write") {
				if err := Replace(r, os.Stdout, replacer); err != nil {
					return err
				}

				continue
			}

			output := bytes.NewBuffer(nil)
			if err := Replace(r, output, replacer); err != nil {
				return err
			}

			fileInfo, err := os.Stat(filePath)
			if err != nil {
				return err
			}

			file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, fileInfo.Mode())
			if err != nil {
				return err
			}
			defer file.Close()

			if _, err := output.WriteTo(file); err != nil {
				return err
			}
		}

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func Replace(r io.Reader, w io.Writer, replacer ReplacerFN) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		before := append(scanner.Bytes(), []byte("\n")...)
		after, err := replacer(before)
		if err != nil {
			return err
		}

		if _, err := w.Write(after); err != nil {
			return err
		}
	}

	return scanner.Err()
}
