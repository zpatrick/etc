package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func tempDir(t *testing.T, dir string) (string, func()) {
	name, err := ioutil.TempDir(dir, "")
	if err != nil {
		t.Fatal(err)
	}

	remove := func() {
		if err := os.RemoveAll(name); err != nil {
			t.Fatal(err)
		}
	}

	return name, remove
}

func tempFile(t *testing.T, dir, prefix string) *os.File {
	file, err := ioutil.TempFile(dir, prefix)
	if err != nil {
		t.Fatal(err)
	}

	return file
}

func TestGlobMatchFiles(t *testing.T) {
	dir, remove := tempDir(t, os.TempDir())
	defer remove()

	files := []*os.File{
		tempFile(t, dir, "foo.tmp"),
		tempFile(t, dir, "bar.tmp"),
		tempFile(t, dir, "baz.tmp"),
	}

	cases := map[string][]string{
		"*": []string{
			files[0].Name(),
			files[1].Name(),
			files[2].Name(),
		},
		"foo*": []string{
			files[0].Name(),
		},
		"bar.tmp*": []string{
			files[1].Name(),
		},
		"ba*": []string{
			files[1].Name(),
			files[2].Name(),
		},
		"b*.tmp*": []string{
			files[1].Name(),
			files[2].Name(),
		},
		"*.tmp*": []string{
			files[0].Name(),
			files[1].Name(),
			files[2].Name(),
		},
		"x": []string{},
	}

	for pattern, expected := range cases {
		t.Run(pattern, func(t *testing.T) {
			result, err := GlobMatchFiles(dir, pattern, false)
			if err != nil {
				t.Fatal(err)
			}

			assert.ElementsMatch(t, expected, result)
		})
	}
}

func TestGlobMatchFilesRecursive(t *testing.T) {
	parentDir, removeParentDir := tempDir(t, os.TempDir())
	defer removeParentDir()

	childDir, removeChildDir := tempDir(t, parentDir)
	defer removeChildDir()

	files := []*os.File{
		tempFile(t, parentDir, "foo.tmp"),
		tempFile(t, parentDir, "bar.tmp"),
		tempFile(t, parentDir, "baz.tmp"),
		tempFile(t, childDir, "foo.tmp"),
		tempFile(t, childDir, "bar.tmp"),
		tempFile(t, childDir, "baz.tmp"),
	}

	cases := map[string][]string{
		"*": []string{
			files[0].Name(),
			files[1].Name(),
			files[2].Name(),
			files[3].Name(),
			files[4].Name(),
			files[5].Name(),
		},
		"foo*": []string{
			files[0].Name(),
			files[3].Name(),
		},
		"bar.tmp*": []string{
			files[1].Name(),
			files[4].Name(),
		},
		"ba*": []string{
			files[1].Name(),
			files[2].Name(),
			files[4].Name(),
			files[5].Name(),
		},
		"b*.tmp*": []string{
			files[1].Name(),
			files[2].Name(),
			files[4].Name(),
			files[5].Name(),
		},
		"*.tmp*": []string{
			files[0].Name(),
			files[1].Name(),
			files[2].Name(),
			files[3].Name(),
			files[4].Name(),
			files[5].Name(),
		},
		"x": []string{},
	}

	for pattern, expected := range cases {
		t.Run(pattern, func(t *testing.T) {
			result, err := GlobMatchFiles(parentDir, pattern, true)
			if err != nil {
				t.Fatal(err)
			}

			assert.ElementsMatch(t, expected, result)
		})
	}
}
