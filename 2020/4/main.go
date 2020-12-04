package main

import (
	"adventofcode/utils"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type passStruct struct {
	Byr string
	Iyr string
	Eyr string
	Hgt string
	Hcl string
	Ecl string
	Pid string
	Cid string
}

var re = regexp.MustCompile("[^0-9]")
var ree = regexp.MustCompile("#[0-9,a-f]{6}")
var reee = regexp.MustCompile("(?m)(in|cm)")

func main() {
	passList, err := utils.GetData("4")
	if err != nil {
		log.Fatal("error load file data", err)
	}
	pass := make([]passStruct, 0, 300)
	sngPass := passStruct{}
	for _, i := range passList {
		if len(i) == 0 {
			pass = append(pass, sngPass)
			sngPass = passStruct{}
		}
		data := strings.Split(i, " ")
		for _, j := range data {
			params := strings.Split(j, ":")
			switch params[0] {
			case "byr":
				sngPass.Byr = params[1]
			case "iyr":
				sngPass.Iyr = params[1]
			case "eyr":
				sngPass.Eyr = params[1]
			case "hgt":
				sngPass.Hgt = params[1]
			case "hcl":
				sngPass.Hcl = params[1]
			case "ecl":
				sngPass.Ecl = params[1]
			case "pid":
				sngPass.Pid = params[1]
			case "cid":
				sngPass.Cid = params[1]
			}
		}
	}
	validPassCount := len(pass)
	validPassDataCount := len(pass)
	for _, k := range pass {
		s := reflect.ValueOf(&k).Elem()
		typeOfT := s.Type()

		if !validPass(&k) {
			validPassDataCount -= 1
		}

		for i := 0; i < s.NumField(); i++ {
			f := s.Field(i)
			vl := f.Interface()
			if len(vl.(string)) == 0 && typeOfT.Field(i).Name != "Cid" {
				validPassCount -= 1
				break
			}
		}

	}
	println(validPassCount)
	println(validPassDataCount)
}

func validPass(k *passStruct) (res bool) {
	if !validInt(k.Byr, 1920, 2002) {
		return
	}

	if !validInt(k.Iyr, 2010, 2020) {
		return false
	}

	if !validInt(k.Eyr, 2020, 2030) {
		return
	}

	if !utils.StringInSlice(k.Ecl, []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}) {
		return
	}

	pidValid := len(k.Pid) == 9 && !re.MatchString(k.Pid)
	if !pidValid {
		return
	}

	if !ree.MatchString(k.Hcl) {
		return
	}
	data := reee.FindAllString(k.Hgt, -1)
	if len(data) != 1 {
		return
	}
	if !(data[0] == "cm" || data[0] == "in") {
		return
	} else {
		if data[0] == "cm" && !validInt(strings.ReplaceAll(k.Hgt, "cm", ""), 150, 193) {
			return
		} else if data[0] == "in" && !validInt(strings.ReplaceAll(k.Hgt, "in", ""), 59, 76) {
			return
		}
	}
	return true
}

func validInt(value string, min, max int) bool {
	i, err := strconv.Atoi(value)
	if err != nil {
		return false
	}
	return i >= min && i <= max
}
