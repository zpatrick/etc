package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	glob "github.com/ryanuber/go-glob"
)

func GlobMatchFiles(dir, pattern string, recursive bool) ([]string, error) {
	if recursive {
		return recursiveGlobMatch(dir, pattern)
	}

	return globMatch(dir, pattern)
}

func recursiveGlobMatch(dir, pattern string) ([]string, error) {
	matches := []string{}
	walkFN := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			dirMatches, err := globMatch(path, pattern)
			if err != nil {
				return err
			}

			matches = append(matches, dirMatches...)
		}

		return nil
	}

	if err := filepath.Walk(dir, walkFN); err != nil {
		return nil, err
	}

	return matches, nil
}

func globMatch(dir, pattern string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	matches := make([]string, 0, len(files))
	for _, file := range files {
		if !file.IsDir() && glob.Glob(pattern, file.Name()) {
			matches = append(matches, filepath.Join(dir, file.Name()))
		}
	}

	return matches, nil
}
