package lenv

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

var paths = map[string]string{
	"env":            ".env",
	"env2":           ".env2",
	"lenv":           ".lenv",
	"testdata_env":   filepath.Join("testdata", ".env"),
	"testdata_a_env": filepath.Join("testdata", "a", ".env"),
	"testdata_b_env": filepath.Join("testdata", "b", ".env"),
}

func TestMain(m *testing.M) {
	code := m.Run()
	clean()
	os.Exit(code)
}

func clean() {
	for _, path := range paths {
		os.Remove(path)
	}
}

func cleanupTempFile(f *os.File) error {
	err := f.Close()
	if err != nil {
		return fmt.Errorf("failed to close temp file: %v", err)
	}
	err = os.Remove(f.Name())
	if err != nil {
		return fmt.Errorf("failed to remove temp file: %v", err)
	}
	return nil
}

func TestGetEnvFilePath(t *testing.T) {
	t.Cleanup(clean)

	tmpEnvFile, err := os.Create(paths["env"])
	if err != nil {
		t.Fatalf("failed to create temp .env file: %v", err)
	}
	defer cleanupTempFile(tmpEnvFile)

	path, err := GetEnvFilePath(paths["env"])
	fmt.Printf("path: %s\n", path)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if path == "" {
		t.Error("expected .env file path to exist")
	}
}

func TestGetEnvFilePath_FileDoesNotExist(t *testing.T) {
	t.Cleanup(clean)

	_, err := os.Stat(paths["env"])
	if err == nil {
		t.Fatal(".env file should not exist before the test")
	}

	path, err := GetEnvFilePath(paths["env"])
	if err == nil {
		t.Error("expected an error when .env file does not exist")
	}
	if path != "" {
		t.Error("expected .env file path to be empty when file does not exist")
	}
}

func TestReadLenvFile(t *testing.T) {
	t.Cleanup(clean)

	tmpLenvFile, err := os.Create(paths["lenv"])
	if err != nil {
		t.Fatalf("failed to create temp .lenv file: %v", err)
	}
	expectedDestinations := []string{"a", "b"}
	tmpLenvFile.WriteString(strings.Join(expectedDestinations, "\n"))
	defer cleanupTempFile(tmpLenvFile)

	destinations, err := ReadLenvFile()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if reflect.DeepEqual(destinations, expectedDestinations) {
		t.Error("expected .lenv file content to match the expected destinations")
	}
}

func TestReadLenvFile_FileDoesNotExist(t *testing.T) {
	t.Cleanup(clean)

	_, err := os.Stat(paths["lenv"])
	if err == nil {
		t.Fatal(".lenv file should not exist before the test")
	}

	destinations, err := ReadLenvFile()
	if err == nil {
		t.Error("expected an error when .lenv file does not exist")
	}
	if !reflect.DeepEqual(destinations, []string{}) {
		t.Error("expected .env file path to be empty when file does not exist")
	}
}

func TestCheck(t *testing.T) {
	t.Cleanup(clean)

	source := paths["env"]
	destinations := []string{"a", "b"}

	tmpEnvFile, err := os.Create(source)
	if err != nil {
		t.Fatalf("failed to create temp .env file: %v", err)
	}
	defer cleanupTempFile(tmpEnvFile)

	for _, dest := range destinations {
		err := os.Symlink(source, dest)
		if err != nil {
			t.Fatalf("failed to create symlink: %v", err)
		}
		defer os.Remove(dest)
	}

	err = Check(source, destinations)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCheck_SymlinkPointsToDifferentSource(t *testing.T) {
	t.Cleanup(clean)

	source := paths["env"]
	destinations := []string{paths["testdata_env"]}

	tmpEnvFile, err := os.Create(source)
	if err != nil {
		t.Fatalf("failed to create temp .env file: %v", err)
	}
	defer cleanupTempFile(tmpEnvFile)

	differentSource := paths["env2"]
	tmpEnvFile2, err := os.Create(differentSource)
	if err != nil {
		t.Fatalf("failed to create temp .env2 file: %v", err)
	}
	defer cleanupTempFile(tmpEnvFile2)

	err = os.Symlink(differentSource, destinations[0])
	if err != nil {
		t.Fatalf("failed to create symlink: %v", err)
	}
	defer os.Remove(destinations[0])

	err = Check(source, destinations)
	if err == nil {
		t.Error("expected an error when symlink points to a different source")
	}
}

func TestCheck_PhysicalFileExists(t *testing.T) {
	t.Cleanup(clean)

	source := paths["env"]
	destinations := []string{paths["testdata_env"]}

	tmpEnvFile, err := os.Create(source)
	if err != nil {
		t.Fatalf("failed to create temp .env file: %v", err)
	}
	defer cleanupTempFile(tmpEnvFile)

	tmpDestFile, err := os.Create(destinations[0])
	if err != nil {
		t.Fatalf("failed to create temp destination file: %v", err)
	}
	defer cleanupTempFile(tmpDestFile)

	err = Check(source, destinations)
	if err == nil {
		t.Error("expected an error when physical file exists at destination")
	}
}

func TestLink(t *testing.T) {
	t.Cleanup(clean)

	source := paths["env"]
	destinations := []string{paths["testdata_a_env"], paths["testdata_b_env"]}

	tmpEnvFile, err := os.Create(source)
	if err != nil {
		t.Fatalf("failed to create temp .env file: %v", err)
	}
	defer cleanupTempFile(tmpEnvFile)

	err = Link(source, destinations)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestLink_PhysicalFileExists(t *testing.T) {
	t.Cleanup(clean)

	source := paths["env"]
	destinations := []string{paths["testdata_a_env"], paths["testdata_b_env"]}

	tmpEnvFile, err := os.Create(source)
	if err != nil {
		t.Fatalf("failed to create temp .env file: %v", err)
	}
	defer cleanupTempFile(tmpEnvFile)

	tmpDestFile, err := os.Create(destinations[0])
	if err != nil {
		t.Fatalf("failed to create temp destination file: %v", err)
	}
	defer cleanupTempFile(tmpDestFile)

	err = Link(source, destinations)
	if err == nil {
		t.Error("expected an error when physical file exists at destination")
	}
}

func TestUnlink(t *testing.T) {
	t.Cleanup(clean)

	source := paths["env"]
	destinations := []string{paths["testdata_a_env"], paths["testdata_b_env"]}

	tmpEnvFile, err := os.Create(source)
	if err != nil {
		t.Fatalf("failed to create temp .env file: %v", err)
	}
	defer cleanupTempFile(tmpEnvFile)

	for _, dest := range destinations {
		err := os.Symlink(source, dest)
		if err != nil {
			t.Fatalf("failed to create symlink: %v", err)
		}
		defer os.Remove(dest)
	}

	err = Unlink(destinations)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestUnlink_NoSymlink(t *testing.T) {
	t.Cleanup(clean)

	destinations := []string{paths["testdata_a_env"], paths["testdata_b_env"]}

	err := Unlink(destinations)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
