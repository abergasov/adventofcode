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

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func UniqueSlice(a []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range a {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
