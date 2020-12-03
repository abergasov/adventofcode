package utils

import (
	"io/ioutil"
	"os"
	"strings"
)

func GetData(pass string) ([]string, error) {
	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadFile(path + "/2020/" + pass + "/input.txt")
	if err != nil {
		return nil, err
	}
	return strings.Split(string(data), "\n"), nil
}
