package util

import (
	"os"
	"strconv"
)

func RequireAtoI(a string) int {
	result, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}

	return result
}

func RequireFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	return file
}