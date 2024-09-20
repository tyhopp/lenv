package lenv

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// GetEnvFilePath returns the absolute path to the '.env' file
// in the current directory.
func GetEnvFilePath() (string, error) {
	path, err := filepath.Abs(".env")
	if err != nil {
		return "", fmt.Errorf("lenv: failed to get absolute path to .env file")
	}

	_, err = os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("lenv: failed to find .env file in current directory")
		} else {
			return "", fmt.Errorf("lenv: failed to check .env file")
		}
	}

	return path, nil
}

// ReadLenvFile reads a '.lenv' file in the current directory
// and returns a slice of absolute destination paths.
//
// Destination paths are not guaranteed to exist or be valid.
func ReadLenvFile() ([]string, error) {
	file, err := os.Open(".lenv")
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, fmt.Errorf("lenv: failed to find .lenv file in current directory")
		} else {
			return []string{}, fmt.Errorf("lenv: failed to check .lenv file in current directory")
		}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return []string{}, fmt.Errorf("lenv: failed to scan .lenv file")
	}

	var destinations []string
	for _, line := range lines {
		path, err := filepath.Abs(line)
		if err != nil {
			return []string{}, fmt.Errorf("lenv: failed to get absolute path to %s", line)
		}
		destinations = append(destinations, path)
	}

	return destinations, nil
}

// Check investigates the current status of symlinks between the
// source '.env' file and destinations.
func Check(source string, destinations []string) error {
	for _, destination := range destinations {
		stats, err := os.Lstat(destination)

		if err != nil {
			if os.IsNotExist(err) {
				fmt.Printf("lenv: no symlinked or physical file found at %s\n", destination)
				continue
			} else {
				return fmt.Errorf("lenv: failed to check %s", destination)
			}
		}

		if stats.Mode().Type() == os.ModeSymlink {
			link, err := os.Readlink(destination)
			if err != nil {
				panic(err)
			}
			if link == source {
				fmt.Printf("lenv: symlink links %s to %s\n", source, destination)
			} else {
				return fmt.Errorf("lenv: symlink %s does not link to %s, it should be removed", destination, source)
			}
		}

		if stats.Mode().IsRegular() {
			return fmt.Errorf("lenv: found physical file %s, it should be removed", destination)
		}
	}
	return nil
}

func Link() {
	fmt.Println("todo: implement link")
}

func Unlink() {
	fmt.Println("todo: implement unlink")
}
