package lenv

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Logger returns a new logger that writes to stdout with the prefix "lenv: ".
func Logger() *log.Logger {
	lenvLog := log.New(os.Stdout, "lenv: ", 0)
	return lenvLog
}

// GetEnvFilePath returns the absolute path to the '.env' file
// in the current directory.
func GetEnvFilePath(logger *log.Logger) string {
	path, err := filepath.Abs(".env")
	if err != nil {
		logger.Fatalf("failed to get absolute path to .env file")
	}

	_, err = os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			logger.Fatalf("failed to find source .env file in current directory")
		} else {
			logger.Fatalf("failed to check source .env file")
		}
	}

	return path
}

// ReadLenvFile reads a '.lenv' file in the current directory
// and returns a slice of absolute destination paths.
//
// Destination paths are not guaranteed to exist or be valid.
func ReadLenvFile(logger *log.Logger) []string {
	file, err := os.Open(".lenv")
	if err != nil {
		if os.IsNotExist(err) {
			logger.Fatalf("failed to find .lenv file in current directory")
		} else {
			logger.Fatalf("failed to read .lenv file in current directory")
		}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		logger.Fatalf("failed to scan .lenv file")
	}

	var destinations []string
	for _, line := range lines {
		path, err := filepath.Abs(line)
		if err != nil {
			logger.Fatalf("failed to get absolute path to %s", line)
		}
		destinations = append(destinations, path)
	}

	return destinations
}

// Check investigates the current status of symlinks between the
// source '.env' file and destinations.
func Check(logger *log.Logger, source string, destinations []string) {
	for _, destination := range destinations {
		stats, err := os.Lstat(destination)

		if err != nil {
			if os.IsNotExist(err) {
				logger.Printf("no symlinked or physical file found at %s", destination)
				continue
			} else {
				logger.Fatalf("failed to check %s", destination)
			}
		}

		if stats.Mode().Type() == os.ModeSymlink {
			link, err := os.Readlink(destination)
			if err != nil {
				panic(err)
			}
			if link == source {
				logger.Printf("symlink links %s to %s", source, destination)
			} else {
				logger.Fatalf("symlink %s does not link to %s, it should be removed", destination, source)
			}
		} else {
			logger.Fatalf("found physical file %s, it should be removed", destination)
		}
	}
}

func Link() {
	fmt.Println("todo: implement link")
}

func Unlink() {
	fmt.Println("todo: implement unlink")
}
