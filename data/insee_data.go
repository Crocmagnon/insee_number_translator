package data

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type InseeData struct {
	InseeNumber  string
	Gender       string
	Year         int
	Month        time.Month
	Department   string
	City         string
	Foreign      bool
	Country      string
	OrderOfBirth int
	ControlKey   int
}

func NewInseeData(inseeNumber string) (*InseeData, error) {
	if len(inseeNumber) != 15 {
		return nil, fmt.Errorf("provided insee number must contain 15 characters")
	}
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
	country := "FR"
	if (dep >= 91 && dep <= 96) || dep == 99 {
		foreign = true
		country = city
		city = "unknown"
		department = "unknown"
	}
	order, err := strconv.Atoi(num[10:13])
	if err != nil {
		return nil, err
	}
	controlKey, err := strconv.Atoi(num[13:])
	if err != nil {
		return nil, err
	}
	monthInt, err := strconv.Atoi(num[3:5])
	if err != nil {
		return nil, err
	}
	month := time.Month(monthInt)
	year, err := strconv.Atoi(num[1:3])
	if err != nil {
		return nil, err
	}
	year += 2000
	now := time.Now()
	birthday := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	if birthday.After(now) {
		year -= 100
	}
	gender := num[0:1]
	if gender == "1" {
		gender = "Male"
	} else {
		gender = "Female"
	}
	return &InseeData{
		InseeNumber:  num,
		Gender:       gender,
		Year:         year,
		Month:        month,
		Department:   department,
		City:         city,
		Foreign:      foreign,
		Country:      country,
		OrderOfBirth: order,
		ControlKey:   controlKey,
	}, nil
}

func (insee InseeData) IsValid() (bool, error) {
	num := strings.Replace(insee.InseeNumber, "2A", "19", 1)
	num = strings.Replace(num, "2B", "18", 1)[:len(num)-2]
	numInt, err := strconv.Atoi(num)
	if err != nil {
		return false, err
	}
	code := 97 - (numInt % 97)
	return code == insee.ControlKey, nil
}

func (insee InseeData) String() string {
	var message []string
	valid, err := insee.IsValid()
	if err != nil {
		return err.Error()
	}
	if valid {
		message = append(message, "The number is valid.")
	} else {
		message = append(message, "The number is invalid.")
	}
	msg := fmt.Sprintf("You're a %s, born in %s, probably in %d.", insee.Gender, insee.Month, insee.Year)
	message = append(message, msg)

	var zoneType string
	var line string
	if insee.Foreign {
		zoneType = "country"
		line = "You were born outside France, "
		country, present := countries["99"+insee.Country]
		continent := continents[insee.Country[0:1]]
		continent = fmt.Sprintf(" (%s)", continent)
		if present && len(country) > 1 {
			line += "probably in either one of these countries/territories: "
			line += strings.Join(country, ", ")
			line += continent
		} else if present && len(country) == 1 {
			line += fmt.Sprintf("in %s%s", country[0], continent)
		} else {
			line += "in an unknown country numbered " + insee.Country
		}
	} else {
		zoneType = "city"
		city, cityPresent := cities[insee.Department+insee.City]
		department, depPresent := departments[insee.Department]
		if !depPresent {
			department = "unknown department"
		}

		if cityPresent {
			line = fmt.Sprintf("You were born in %s, %s (%s), France.", city.Name, city.ZipCode, department)
		} else {
			line = fmt.Sprintf("You were born in an unknown city numbered %s, which is located in %s, in France", insee.City, department)
		}
	}
	message = append(message, line)
	line = fmt.Sprintf("You were the %dth to be born in this %s on this month.", insee.OrderOfBirth, zoneType)
	message = append(message, line)

	return strings.Join(message, "\n")
}
