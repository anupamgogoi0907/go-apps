package utility

import (
	"errors"
	"log"
	"os"
	"strconv"
	"unicode"
)

const (
	B  = "B"
	KB = "KB"
	MB = "MB"
	GB = "GB"

	b  = 1
	kb = 1024 * b
	mb = 1024 * kb
	gb = 1024 * mb
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

func GetChunkSize(chunkSize string) int {
	value, unit := SplitSize(chunkSize)

	switch unit {
	case B:
		value = value * b
	case KB:
		value = value * kb
	case MB:
		value = value * mb
	case GB:
		value = value * gb
	}
	return value
}

func SplitSize(chunkSize string) (int, string) {
	var unit string
	var value string
	for _, letter := range chunkSize {
		if unicode.IsNumber(letter) {
			value = value + string(letter)
		} else if unicode.IsLetter(letter) {
			unit = unit + string(letter)
		}

	}
	intVar, err := strconv.Atoi(value)
	if err != nil {
		log.Panicln(err)
		return 0, ""
	}
	return intVar, unit
}
