package test

import (
	"fmt"
	"os"
)

func LoadTestData(filename string) []byte {
	f, err := os.ReadFile("../test/" + filename)
	if err != nil {
		panic(fmt.Sprintf("Cannot load test data %v", filename))
	}
	return f
}
