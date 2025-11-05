package test

import (
	"fmt"
	"os"
)

func LoadTestDataFromCurrentDir(filename string) []byte {
	return LoadTestDataV2("", filename)
}

func LoadTestData(filename string) []byte {
	return LoadTestDataV2("../../test/", filename)
}

func LoadTestDataV2(dir string, filename string) []byte {
	f, err := os.ReadFile(dir + filename)
	if err != nil {
		panic(fmt.Sprintf("Cannot load test data %v", filename))
	}
	return f
}
