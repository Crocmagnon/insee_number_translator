package data

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var Male = "homme"
var Female = "femme"
var Unknown = "inconnu(e)"

//go:embed curated_data/countries.json
var rawCountries []byte

//go:embed curated_data/departments.json
var rawDepartments []byte

//go:embed curated_data/cities.json
var rawCities []byte

type InseeData struct {
	InseeNumber  string     `json:"insee_number"`
	Gender       string     `json:"gender"`
	Year         int        `json:"year"`
	Month        time.Month `json:"month"`
	Department   string     `json:"department"`
	City         string     `json:"city"`
	CityCode     string     `json:"city_code"`
	Foreign      bool       `json:"foreign"`
	Countries    []string   `json:"countries"`
	CountryCode  string     `json:"country_code"`
	Continent    string     `json:"continent"`
	OrderOfBirth int        `json:"order_of_birth"`
	ControlKey   int        `json:"control_key"`
}

func NewInseeData(inseeNumber string) (*InseeData, error) {
	if len(inseeNumber) != 15 {
		return nil, fmt.Errorf("provided insee number must contain 15 characters")
	}
	num := inseeNumber
	departmentCode := num[5:7]
	cityCode := num[7:10]
	if departmentCode == "97" || departmentCode == "98" {
		departmentCode = num[5:8]
		cityCode = num[8:10]
	}
	dep, err := strconv.Atoi(departmentCode)
	if err != nil && departmentCode != "2A" && departmentCode != "2B" {
		return nil, err
	}
	var city string
	var department string
	countries_ := []string{"FRANCE"}
	countryCode := ""
	continent := "Europe"
	foreign := (dep >= 91 && dep <= 96) || dep == 99
	if foreign {
		foreign = true
		countryCode = cityCode
		cityCode = ""
		countries_ = getCountry("99" + countryCode)
		continent = continents[countryCode[0:1]]
		department = Unknown
	} else {
		city = getCity(departmentCode + cityCode)
		department = getDepartment(departmentCode)
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
	birthday := time.Date(year, month, 1, 0, 0, 0, 0, now.Location())
	if birthday.After(now) {
		year -= 100
	}
	gender := num[0:1]
	if gender == "1" {
		gender = Male
	} else {
		gender = Female
	}
	return &InseeData{
		InseeNumber:  num,
		Gender:       gender,
		Year:         year,
		Month:        month,
		Department:   department,
		City:         city,
		CityCode:     cityCode,
		Foreign:      foreign,
		Countries:    countries_,
		CountryCode:  countryCode,
		Continent:    continent,
		OrderOfBirth: order,
		ControlKey:   controlKey,
	}, nil
}

func (insee InseeData) IsValid() (bool, error) {
	r := strings.NewReplacer(
		"2A", "19",
		"2B", "18",
	)
	num := r.Replace(insee.InseeNumber[:len(insee.InseeNumber)-2])
	numInt, err := strconv.Atoi(num)
	if err != nil {
		return false, err
	}
	code := 97 - (numInt % 97)
	return code == insee.ControlKey, nil
}

func (insee InseeData) String() string {
	var result []string
	result = append(result, insee.InseeNumber)
	var line string
	valid, err := insee.IsValid()
	if err != nil {
		line = "Erreur lors de la vérification de la validité : " + err.Error()
	} else if valid {
		line = "Le numéro est valide."
	} else {
		line = "Le numéro est invalide."
	}
	result = append(result, line)
	var one, born string
	if insee.Gender == Male {
		one = "un"
		born = "né"
	} else {
		one = "une"
		born = "née"
	}
	line = fmt.Sprintf("Vous êtes %s %s, %s en %s probablement en %d.", one, insee.Gender, born, frenchMonth(insee.Month), insee.Year)
	result = append(result, line)
	line = ""
	var zoneType string
	if insee.Foreign {
		zoneType = "ce pays"
		if len(insee.Countries) > 1 {
			line = fmt.Sprintf("Vous êtes %s dans l'un de ces pays/territoires: %s (%s)", born, strings.Join(insee.Countries, ", "), insee.Continent)
		} else if len(insee.Countries) == 1 {
			line = fmt.Sprintf("Vous êtes %s en %s (%s).", born, insee.Countries[0], insee.Continent)
		} else {
			line = fmt.Sprintf("Vous êtes %s dans un pays inconnu, portant l'identifiant %s.", born, insee.CountryCode)
		}
	} else {
		zoneType = "cette ville"
		if insee.City != Unknown {
			line = fmt.Sprintf("Vous êtes %s à %s (%s, France)", born, insee.City, insee.Department)
		} else {
			line = fmt.Sprintf("Vous êtes %s dans une ville inconnue portant l'identifiant %s dans le département \"%s\"", born, insee.CityCode, insee.Department)
		}
	}
	result = append(result, line)
	line = fmt.Sprintf("Vous êtes la %de personne née dans %s le mois de votre naissance.", insee.OrderOfBirth, zoneType)
	result = append(result, line)
	lineB, err := json.Marshal(insee)
	if err != nil {
		line = "Erreur lors de la conversion en JSON: " + err.Error()
	} else {
		line = string(lineB)
	}
	result = append(result, line)
	return strings.Join(result, "\n")
}

func getCity(cityCode string) string {
	return getString(cityCode, rawCities)
}

func getDepartment(departmentCode string) string {
	return getString(departmentCode, rawDepartments)
}

func getString(code string, rawData []byte) string {
	cities := map[string]string{}
	err := json.Unmarshal(rawData, &cities)
	if err != nil {
		return Unknown
	}
	item, present := cities[code]
	if present {
		return item
	} else {
		return Unknown
	}
}

func getCountry(code string) []string {
	cities := map[string][]string{}
	err := json.Unmarshal(rawCountries, &cities)
	if err != nil {
		return []string{}
	}
	city, present := cities[code]
	if present {
		return city
	} else {
		return []string{}
	}
}
