package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	res, err := getData()
	if err != nil {
		log.Fatal("error load file data", err)
	}
	a, b := stepOne(res, 2020)
	println(a, b, a*b)

	a, b, c := stepTwo(res, 2020)
	println(a, b, c, a*b*c)
}

func stepOne(res map[int]struct{}, sum int) (int, int) {
	for i := range res {
		sec := sum - i
		if sec < 0 {
			continue
		}
		_, ok := res[sec]
		if !ok {
			continue
		}
		return i, sec
	}
	return 0, 0
}

func stepTwo(res map[int]struct{}, year int) (int, int, int) {
	for i := range res {
		sec := year - i
		n, m := stepOne(res, sec)
		if n == 0 || m == 0 {
			continue
		}
		return i, n, m
	}
	return 0, 0, 0
}

func getData() (map[int]struct{}, error) {
	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadFile(path + "/2020/1/input.txt")
	if err != nil {
		return nil, err
	}
	res := make(map[int]struct{})
	for _, d := range strings.Split(string(data), "\n") {
		i, _ := strconv.Atoi(d)
		res[i] = struct{}{}
	}
	return res, nil
}
