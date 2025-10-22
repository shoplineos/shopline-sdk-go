package test

import (
	"fmt"
	"os"
)

func LoadTestData(filename string) []byte {
	return LoadTestDataV2("../../test/", filename)
}

func LoadTestDataV2(prefix string, filename string) []byte {
	f, err := os.ReadFile(prefix + filename)
	if err != nil {
		panic(fmt.Sprintf("Cannot load test data %v", filename))
	}
	return f
}
