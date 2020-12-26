package intergal

import (
	"io/ioutil"
	"os"
)

func readAll(filePath string) ([]byte, error) {
	f, err := os.Open(filePath)
	ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}
