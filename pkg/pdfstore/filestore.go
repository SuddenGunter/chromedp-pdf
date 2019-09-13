package pdfstore

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"
)

type FileStoreConfig struct {
	// Path is OS-specific path where to store generated PDF files, e.g. /home/user/pdf
	Path string
	// GetFileName is a function that returns filename for PDF. Optional, for default value see GetDefaultFileNameGenerator
	GetFileName func() string
}

type FileStore struct {
	config FileStoreConfig
}

func NewFileStore(config *FileStoreConfig) (FileStore, error) {
	if err := validateConfig(config); err != nil {
		return FileStore{}, err
	}
	return FileStore{config: *config}, nil
}

func validateConfig(config *FileStoreConfig) error {
	if len(config.Path) <= 0 {
		return ErrEmptyPath
	}
	if config.GetFileName == nil {
		config.GetFileName = GetDefaultFileNameGenerator()
	}

	return nil
}

func (fs FileStore) Write(buffer []byte) (n int, err error) {
	filePath := path.Join(fs.config.Path, fs.config.GetFileName())
	if err = ioutil.WriteFile(filePath, buffer, os.ModePerm); err != nil {
		return 0, err
	}
	return len(buffer), nil
}

// GetDefaultFileNameGenerator provides function that generates file names in next format
// {time.Now().UTC().UnixNano()}.pdf
func GetDefaultFileNameGenerator() func() string {
	return func() string {
		return fmt.Sprintf("&v.pdf", time.Now().UTC().UnixNano())
	}
}

var ErrEmptyPath = errors.New("FileStoreConfig.Path must be not empty")
