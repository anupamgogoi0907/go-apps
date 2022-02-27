package utility

import (
	"errors"
	"log"
	"os"
)

func CreateDir(dir string) error {
	_, err := os.Stat(dir)
	if errors.Is(err, os.ErrNotExist) {
		errDir := os.MkdirAll(dir, 0755)
		if errDir != nil {
			log.Fatal(err)
			return errDir
		}
	}
	return nil
}
