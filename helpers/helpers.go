package helpers

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func LoadTestScanner(filename string) (*bufio.Scanner, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return bufio.NewScanner(f), nil
}

func StringToInts(s string) []int {
	intSlice := []int{}
	for _, vString := range strings.Split(s, " ") {
		vInt, err := strconv.Atoi(vString)
		if err != nil {
			continue
		}
		intSlice = append(intSlice, vInt)
	}
	return intSlice
}