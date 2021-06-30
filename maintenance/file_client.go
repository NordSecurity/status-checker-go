package maintenance

import (
	"os"

	"github.com/pkg/errors"
)

type fileClient struct {
	filename string
}

func NewFileClient(filename string) Client {
	return &fileClient{filename: filename}
}

func (fc *fileClient) IsMaintenanceEnabled() (bool, error) {
	info, err := os.Stat(fc.filename)
	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, errors.Wrap(err, "stat file")
	}

	return !info.IsDir(), nil
}
