package pdfstore

import (
	"errors"
	"io/ioutil"
	"os"
)

type FileStoreConfig struct {
	Path              string
	Permissions       os.FileMode
	FileNameGenerator func() string
}

type FileStore struct {
	config FileStoreConfig
}

func NewFileStore(config *FileStoreConfig) (*FileStore, error) {
	if err := validateConfig(config); err != nil {
		return nil, err
	}

	return &FileStore{config: *config}, nil
}

func validateConfig(config *FileStoreConfig) error {
	if len(config.Path) <= 0 {
		return ErrEmptyPath
	}
	if config.FileNameGenerator == nil {
		return ErrNilFileNameGenerator
	}

	return nil
}

func (fs FileStore) Write(buffer []byte) (n int, err error) {
	if err = ioutil.WriteFile(fs.config.Path+fs.config.FileNameGenerator(), buffer, fs.config.Permissions); err != nil {
		return 0, err
	}
	return len(buffer), nil
}

// DefaultFileNameGenerator provides function that generates file names in next format
// dd-mm-yyyy-hh-mm-{base64 string of 6 chars}.pdf
func DefaultFileNameGenerator() func() string {
	return func() string {
		//todo
		return "NOT_IMPLEMENTED.pdf"
	}
}

var ErrEmptyPath = errors.New("FileStoreConfig.Path must be not empty")
var ErrNilFileNameGenerator = errors.New("FileStoreConfig.FileNameGenerator must be non-nil value")
