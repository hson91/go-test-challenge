package libs

import (
	"io/ioutil"
	"os"

	"github.com/go-test-challenge/errors"
)

// ReadFile : read data from file json
// Params : path
// Result : []byte
func ReadFile(path string) ([]byte, error) {
	fi, err := os.Open(path)
	if err != nil {
		return nil, errors.ErrorWithMessage(errors.OpenFileHasError, err.Error())
	}

	defer fi.Close()

	valueBytes, err := ioutil.ReadAll(fi)

	if err != nil {
		return nil, errors.ErrorWithMessage(errors.ReadFileHasError, err.Error())
	}
	return valueBytes, nil
}
