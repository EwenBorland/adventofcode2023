package helpers

import (
	"bufio"
	"os"
)

func LoadTestScanner(filename string) (*bufio.Scanner, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return bufio.NewScanner(f), nil
}