package main

import (
	"fmt"
	"strings"
)

// Given an absolute path for a file (Unix-style), simplify it.

/*
path = "/home/", => "/home"
path = "/a/./b/../../c/", => "/c"
*/

// https://www.interviewbit.com/courses/programming/

func AbsPath(path string) (string, error) {
	stack := []string{}
	for _, s := range strings.Split(path, "/") {
		switch s {
		case ".", "":
		case "..":
			if len(stack) == 0 {
				return "", fmt.Errorf("Path above root!")
			}

			stack = stack[:len(stack)-1]
		default:
			stack = append(stack, s)
		}
	}

	var abs string
	for _, s := range stack {
		abs += "/" + s
	}

	return abs, nil
}

func main() {
	fmt.Println(AbsPath("/a/./b/../../c/d"))
}
