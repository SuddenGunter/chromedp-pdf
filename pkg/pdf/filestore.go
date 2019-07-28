package pdf

import (
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

type FileSystemStoreConfig struct {
	Path        string
	Permissions os.FileMode
}

type FileSystemStore struct {
	config FileSystemStoreConfig
	random *rand.Rand
}

func NewFileSystemStore(config *FileSystemStoreConfig) (*FileSystemStore, error) {
	if err := validateConfig(config); err != nil {
		return nil, err
	}

	return &FileSystemStore{config: *config, random: rand.New(rand.NewSource(time.Now().Unix()))}, nil
}

func validateConfig(config *FileSystemStoreConfig) error {
	if len(config.Path) <= 0 {
		return ErrEmptyPath
	}

	return nil
}

func (fs *FileSystemStore) Write(buffer []byte) (n int, err error) {
	fullName := path.Join(fs.config.Path, fs.nextFileName())
	if err = ioutil.WriteFile(fullName, buffer, fs.config.Permissions); err != nil {
		return 0, errors.Wrap(err, "Unable to save PDF")
	}
	return len(buffer), nil
}

func (fs *FileSystemStore) nextFileName() string {
	return strconv.FormatInt(time.Now().Unix(), 10) + "__" + strconv.Itoa(fs.random.Int()) + ".pdf"
}

var ErrEmptyPath = errors.New("FileSystemStoreConfig.Path must be not empty")
