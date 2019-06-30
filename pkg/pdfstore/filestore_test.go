package pdfstore

import (
	"fmt"
	"testing"
)

func TestNewFileStore_WithCorrectConfig_ReturnsNewFileStore(t *testing.T) {
	config := &FileStoreConfig{
		Path:              "/store",
		Permissions:       0666,
		FileNameGenerator: DefaultFileNameGenerator(),
	}

	fs, err := NewFileStore(config)
	if err != nil {
		t.Error(fmt.Errorf("unable to create default file store: %v", err))
	}

	if fs == nil {
		t.Error("unable to create default file store")
	}
}

func TestNewFileStore_WithEmptyPath_ReturnsError(t *testing.T) {
	config := &FileStoreConfig{
		Path:              "",
		Permissions:       0666,
		FileNameGenerator: DefaultFileNameGenerator(),
	}

	_, err := NewFileStore(config)
	if err != ErrEmptyPath {
		t.Error("new file store config validation must check that path is not empty")
	}
}

func TestNewFileStore_WithNilNameGenerator_ReturnsError(t *testing.T) {
	config := &FileStoreConfig{
		Path:        "/store",
		Permissions: 0666,
	}

	_, err := NewFileStore(config)
	if err != ErrNilFileNameGenerator {
		t.Error("new file store config validation must check that fileNameGenerator() is not nil")
	}
}
