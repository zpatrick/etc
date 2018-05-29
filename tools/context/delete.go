package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
)

func Delete(c *cli.Context) error {
	dir := c.GlobalString("dir")

	for _, context := range c.Args() {
		path := fmt.Sprintf("%s/%s", dir, context)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			return fmt.Errorf("Context '%s' does not exist", context)
		}

		if !c.Bool("force") {
			ok, err := prompt(context)
			if err != nil {
				return err
			}

			if !ok {
				return nil
			}
		}

		if err := os.Remove(path); err != nil {
			return err
		}
	}

	return nil
}

func prompt(context string) (bool, error) {
	fmt.Printf("remove context '%s'?: ", context)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}

	text = strings.ToLower(text)
	return text == "y\n" || text == "yes\n", nil
}
