package libs

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/go-test-challenge/errors"
)

// ReadFile : read data from file json
// Params : path
// Result : []byte
func ReadFile(path string) ([]byte, error) {
	dir, _ := filepath.Split(os.Args[0])
	filename := filepath.Join(dir, path)

	fi, err := os.Open(filename)
	if err != nil {
		fi, err = os.Open(path)
		if err != nil {
			return nil, errors.ErrorWithMessage(errors.OpenFileHasError, err.Error())
		}
	}

	defer fi.Close()

	valueBytes, err := ioutil.ReadAll(fi)

	if err != nil {
		return nil, errors.ErrorWithMessage(errors.ReadFileHasError, err.Error())
	}
	return valueBytes, nil
}
