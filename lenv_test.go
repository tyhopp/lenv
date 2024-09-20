package lenv

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	clean()
	os.Exit(code)
}

func clean() {
	os.Remove(".env")
	os.Remove(".env2")
	os.Remove(".lenv")
	os.Remove("testdata/.env")
	os.Remove("testdata/a/.env")
	os.Remove("testdata/b/.env")
}

func TestGetEnvFilePath(t *testing.T) {
	t.Cleanup(clean)

	tmpEnvFile, err := os.Create(".env")
	if err != nil {
		t.Fatalf("failed to create temp .env file: %v", err)
	}
	defer os.Remove(tmpEnvFile.Name())

	path, err := GetEnvFilePath(".env")
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

	_, err := os.Stat(".env")
	if err == nil {
		t.Fatal(".env file should not exist before the test")
	}

	path, err := GetEnvFilePath(".env")
	if err == nil {
		t.Error("expected an error when .env file does not exist")
	}
	if path != "" {
		t.Error("expected .env file path to be empty when file does not exist")
	}
}

func TestReadLenvFile(t *testing.T) {
	t.Cleanup(clean)

	tmpLenvFile, err := os.Create(".lenv")
	if err != nil {
		t.Fatalf("failed to create temp .lenv file: %v", err)
	}
	expectedDestinations := []string{"a", "b"}
	tmpLenvFile.WriteString(strings.Join(expectedDestinations, "\n"))
	defer os.Remove(tmpLenvFile.Name())

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

	_, err := os.Stat(".lenv")
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

	source := ".env"
	destinations := []string{"a", "b"}

	tmpEnvFile, err := os.Create(source)
	if err != nil {
		t.Fatalf("failed to create temp .env file: %v", err)
	}
	defer os.Remove(tmpEnvFile.Name())

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

	source := ".env"
	destinations := []string{"testdata/.env"}

	tmpEnvFile, err := os.Create(source)
	if err != nil {
		t.Fatalf("failed to create temp .env file: %v", err)
	}
	defer os.Remove(tmpEnvFile.Name())

	differentSource := ".env2"
	tmpEnvFile2, err := os.Create(differentSource)
	if err != nil {
		t.Fatalf("failed to create temp .env2 file: %v", err)
	}
	defer os.Remove(tmpEnvFile2.Name())

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

	source := ".env"
	destinations := []string{"testdata/.env"}

	tmpEnvFile, err := os.Create(source)
	if err != nil {
		t.Fatalf("failed to create temp .env file: %v", err)
	}
	defer os.Remove(tmpEnvFile.Name())

	tmpDestFile, err := os.Create(destinations[0])
	if err != nil {
		t.Fatalf("failed to create temp destination file: %v", err)
	}
	defer os.Remove(tmpDestFile.Name())

	err = Check(source, destinations)
	if err == nil {
		t.Error("expected an error when physical file exists at destination")
	}
}

func TestLink(t *testing.T) {
	t.Cleanup(clean)

	source := ".env"
	destinations := []string{"testdata/a/.env", "testdata/b/.env"}

	tmpEnvFile, err := os.Create(source)
	if err != nil {
		t.Fatalf("failed to create temp .env file: %v", err)
	}
	defer os.Remove(tmpEnvFile.Name())

	err = Link(source, destinations)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestLink_PhysicalFileExists(t *testing.T) {
	t.Cleanup(clean)

	source := ".env"
	destinations := []string{"testdata/a/.env", "testdata/b/.env"}

	tmpEnvFile, err := os.Create(source)
	if err != nil {
		t.Fatalf("failed to create temp .env file: %v", err)
	}
	defer os.Remove(tmpEnvFile.Name())

	tmpDestFile, err := os.Create(destinations[0])
	if err != nil {
		t.Fatalf("failed to create temp destination file: %v", err)
	}
	defer os.Remove(tmpDestFile.Name())

	err = Link(source, destinations)
	if err == nil {
		t.Error("expected an error when physical file exists at destination")
	}
}

func TestUnlink(t *testing.T) {
	t.Cleanup(clean)

	source := ".env"
	destinations := []string{"testdata/a/.env", "testdata/b/.env"}

	tmpEnvFile, err := os.Create(source)
	if err != nil {
		t.Fatalf("failed to create temp .env file: %v", err)
	}
	defer os.Remove(tmpEnvFile.Name())

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

	destinations := []string{"testdata/a/.env", "testdata/b/.env"}

	err := Unlink(destinations)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
