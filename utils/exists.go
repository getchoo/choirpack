package utils

import (
	"errors"
	"io/fs"
	"log"
	"os"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		log.Fatal(err)
	}

	return err == nil
}
