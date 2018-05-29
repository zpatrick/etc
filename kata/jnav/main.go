package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/urfave/cli"
	"github.com/zpatrick/etc/kata/jnav/nav"
)

func main() {
	cli.OsExiter = func(i int) {}

	if len(os.Args) == 1 {
		fmt.Println("Argument PATH is required")
		os.Exit(1)
	}

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	navigator := nav.NewNavigator()
	if err := navigator.Load(data); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		app := cli.NewApp()
		app.Commands = []cli.Command{
			ChangeNode(navigator),
			Concatenate(navigator),
			Exit(),
			List(navigator),
		}

		args := append([]string{"main"}, strings.Split(scanner.Text(), " ")...)
		app.Run(args)

		prompt := fmt.Sprintf("[jnav]: ")
		if name := navigator.CurrentNode().Name; name != "root" {
			prompt = fmt.Sprintf("[jnav %s]: ", name)
		}

		fmt.Printf(prompt)
	}
}

func List(n *nav.Navigator) cli.Command {
	return cli.Command{
		Name: "ls",
		Action: func(c *cli.Context) error {
			for _, child := range n.List() {
				fmt.Printf("%s ", child.Name)
			}

			fmt.Println()
			return nil
		},
	}
}

func ChangeNode(n *nav.Navigator) cli.Command {
	return cli.Command{
		Name: "cd",
		Action: func(c *cli.Context) error {
			if c.NArg() == 0 {
				return nil
			}

			return n.ChangeNode(c.Args().Get(0))
		},
	}
}

func Concatenate(n *nav.Navigator) cli.Command {
	return cli.Command{
		Name: "cat",
		Action: func(c *cli.Context) error {
			node := n.CurrentNode()
			if c.NArg() == 0 {
				fmt.Printf("%v\n", node.Value)
				return nil
			}

			name := c.Args().Get(0)
			child := node.Child(name)
			if child == nil {
				return fmt.Errorf("%s: no such node", name)
			}

			fmt.Printf("%v\n", child.Value)
			return nil
		},
	}
}

func Exit() cli.Command {
	return cli.Command{
		Name: "exit",
		Action: func(c *cli.Context) error {
			os.Exit(0)
			return nil
		},
	}
}
