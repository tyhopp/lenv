package lenv

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// GetEnvFilePath returns the absolute path to the env file
// in the current directory.
func GetEnvFilePath(env string) (string, error) {
	path, err := filepath.Abs(env)
	if err != nil {
		return "", fmt.Errorf("lenv: failed to get absolute path to %s file", env)
	}

	_, err = os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("lenv: failed to find %s file in current directory", env)
		} else {
			return "", fmt.Errorf("lenv: failed to check %s file", env)
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
// source env file and destinations.
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
			relSource, err := filepath.Rel(filepath.Dir(destination), source)
			if err != nil {
				return fmt.Errorf("lenv: failed to get relative path from %s to %s", destination, source)
			}
			if link == relSource {
				fmt.Printf("lenv: symlink links %s to %s\n", source, destination)
			} else {
				return fmt.Errorf("lenv: symlink %s does not link to %s, it should be removed", destination, source)
			}
		}

		if stats.Mode().IsRegular() {
			return fmt.Errorf("lenv: found physical file %s, it should be removed", destination)
		}

		if stats.Mode().IsDir() {
			return fmt.Errorf("lenv: found directory %s, destination paths should be files", destination)
		}
	}
	return nil
}

// Link creates symlinks between the source env file and the destinations.
func Link(source string, destinations []string) error {
	for _, destination := range destinations {
		stats, err := os.Lstat(destination)
		if err != nil {
			if !os.IsNotExist(err) {
				return fmt.Errorf("lenv: failed to check %s before symlinking", destination)
			}
		}
		if stats != nil {
			if stats.Mode().Type() == os.ModeSymlink {
				fmt.Printf("lenv: symlink already exists at %s, skipping\n", destination)
				continue
			}
			if stats.Mode().IsRegular() {
				return fmt.Errorf("lenv: physical file at %s should be removed first", destination)
			}
		}
		relSource, err := filepath.Rel(filepath.Dir(destination), source)
		if err != nil {
			return fmt.Errorf("lenv: failed to get relative path from %s to %s", destination, source)
		}
		err = os.Symlink(relSource, destination)
		if err != nil {
			return fmt.Errorf("lenv: failed to symlink %s to %s", source, destination)
		}
		fmt.Printf("lenv: symlinked %s to %s\n", source, destination)
	}
	return nil
}

// Unlink removes symlinks at the destinations.
func Unlink(destinations []string) error {
	for _, destination := range destinations {
		stats, err := os.Lstat(destination)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Printf("lenv: no symlink found at %s, skipping\n", destination)
				continue
			}
			return fmt.Errorf("lenv: failed to check %s before removing symlink", destination)
		}

		if stats.Mode().Type() != os.ModeSymlink {
			fmt.Printf("lenv: no symlink found at %s, skipping\n", destination)
			continue
		}

		err = os.Remove(destination)
		if err != nil {
			return fmt.Errorf("lenv: failed to remove symlink at %s", destination)
		}
		fmt.Printf("lenv: removed symlink at %s\n", destination)
	}
	return nil
}
