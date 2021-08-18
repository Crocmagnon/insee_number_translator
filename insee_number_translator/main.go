package main

import (
	"strconv"
	"strings"
	"time"
)

type InseeData struct {
	inseeNumber  string
	gender       string
	year         int
	month        int
	department   string
	city         string
	foreign      bool
	country      string
	orderOfBirth int
	controlKey   int
}

func newInseeData(inseeNumber string) (*InseeData, error) {
	num := inseeNumber
	department := num[5:7]
	city := num[7:10]
	if department == "97" || department == "98" {
		department = num[5:8]
		city = num[8:10]
	}
	dep, err := strconv.Atoi(department)
	if err != nil {
		return nil, err
	}
	foreign := false
	country := ""
	if (dep >= 91 && dep <= 96) || dep == 99 {
		foreign = true
		country = city
		city = ""
	}
	order, err := strconv.Atoi(num[10:13])
	if err != nil {
		return nil, err
	}
	controlKey, err := strconv.Atoi(num[13:])
	if err != nil {
		return nil, err
	}
	month, err := strconv.Atoi(num[3:5])
	if err != nil {
		return nil, err
	}
	year, err := strconv.Atoi(num[1:3])
	if err != nil {
		return nil, err
	}
	year += 2000
	now := time.Now()
	birthday := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, nil)
	if birthday.After(now) {
		year -= 100
	}
	return &InseeData{
		inseeNumber:  num,
		gender:       num[0:1],
		year:         year,
		month:        month,
		department:   department,
		city:         city,
		foreign:      foreign,
		country:      country,
		orderOfBirth: order,
		controlKey:   controlKey,
	}, nil
}

func (insee InseeData) isValid() (bool, error) {
	num := strings.Replace(insee.inseeNumber, "2A", "19", 1)
	num = strings.Replace(num, "2B", "18", 1)[:len(num)-2]
	numInt, err := strconv.Atoi(num)
	if err != nil {
		return false, err
	}
	code := 97 - (numInt % 97)
	return code == insee.controlKey, nil
}

func main() {

}
