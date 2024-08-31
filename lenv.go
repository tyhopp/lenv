package lenv

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// ReadLenvFile reads a '.lenv' file in the current directory
// and returns a slice of absolute paths.
//
// Paths are not guaranteed to exist or be valid.
func ReadLenvFile() []string {
	file, err := os.Open(".lenv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	var paths []string
	for _, line := range lines {
		path, err := filepath.Abs(line)
		if err != nil {
			panic(err)
		}
		paths = append(paths, path)
	}

	return paths
}

func Check() {
	fmt.Println("todo: implement check")
}

func Link() {
	fmt.Println("todo: implement link")
}

func Unlink() {
	fmt.Println("todo: implement unlink")
}
