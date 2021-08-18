package data

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewInseeData_ValidForeign(t *testing.T) {
	assert := assert.New(t)

	number := "269059913116714"
	insee, err := NewInseeData(number)
	if err != nil {
		t.Error(err)
	}
	valid, err := insee.IsValid()
	if err != nil {
		t.Error(err)
	}
	assert.True(valid)
	assert.Equal(number, insee.InseeNumber)
	assert.Equal(Female, insee.Gender)
	assert.Equal(1969, insee.Year)
	assert.Equal(time.Month(5), insee.Month)
	assert.Equal(Unknown, insee.Department)
	assert.Nil(insee.City)
	assert.True(insee.Foreign)
	assert.Equal([]string{"BELGIQUE"}, insee.Countries)
	assert.Equal(167, insee.OrderOfBirth)
	assert.Equal(14, insee.ControlKey)
	assert.Equal("131", insee.CountryCode)
	assert.Equal("", insee.CityCode)
	assert.Equal("Europe", insee.Continent)
}

func TestNewInseeData_InvalidFrench(t *testing.T) {
	assert := assert.New(t)

	number := "168127982980507"
	insee, err := NewInseeData(number)
	if err != nil {
		t.Error(err)
	}
	valid, err := insee.IsValid()
	if err != nil {
		t.Error(err)
	}
	assert.False(valid)
	assert.Equal(number, insee.InseeNumber)
	assert.Equal(Male, insee.Gender)
	assert.Equal(1968, insee.Year)
	assert.Equal(time.Month(12), insee.Month)
	assert.Equal("Deux-Sèvres", insee.Department)
	assert.Nil(insee.City)
	assert.False(insee.Foreign)
	assert.Equal([]string{"FRANCE"}, insee.Countries)
	assert.Equal(805, insee.OrderOfBirth)
	assert.Equal(7, insee.ControlKey)
	assert.Equal("", insee.CountryCode)
	assert.Equal("829", insee.CityCode)
	assert.Equal("Europe", insee.Continent)
}

func TestNewInseeData_ValidFrench(t *testing.T) {
	assert := assert.New(t)

	number := "201120100512334"
	insee, err := NewInseeData(number)
	if err != nil {
		t.Error(err)
	}
	valid, err := insee.IsValid()
	if err != nil {
		t.Error(err)
	}
	assert.True(valid)
	assert.Equal(number, insee.InseeNumber)
	assert.Equal(Female, insee.Gender)
	assert.Equal(2001, insee.Year)
	assert.Equal(time.Month(12), insee.Month)
	assert.Equal("Ain", insee.Department)
	assert.NotNil(insee.City)
	assert.Equal("AMBERIEUX-EN-DOMBES", insee.City.Name)
	assert.Equal("01330", insee.City.ZipCode)
	assert.False(insee.Foreign)
	assert.Equal([]string{"FRANCE"}, insee.Countries)
	assert.Equal(123, insee.OrderOfBirth)
	assert.Equal(34, insee.ControlKey)
	assert.Equal(insee.CountryCode, "")
	assert.Equal(insee.CityCode, "005")
	assert.Equal("Europe", insee.Continent)
}

func TestNewInseeData_ValidMultiCountry(t *testing.T) {
	assert := assert.New(t)

	number := "144089943287340"
	insee, err := NewInseeData(number)
	if err != nil {
		t.Error(err)
	}
	valid, err := insee.IsValid()
	if err != nil {
		t.Error(err)
	}
	assert.True(valid)
	assert.Equal(number, insee.InseeNumber)
	assert.Equal(Male, insee.Gender)
	assert.Equal(1944, insee.Year)
	assert.Equal(time.Month(8), insee.Month)
	assert.Equal(Unknown, insee.Department)
	assert.Nil(insee.City)
	assert.True(insee.Foreign)
	expectedCountry := []string{
		"PORTO RICO",
		"VIERGES DES ETATS-UNIS (ILES)",
		"TERR. DES ETATS-UNIS D'AMERIQUE EN AMERIQUE",
	}
	assert.Equal(expectedCountry, insee.Countries)
	assert.Equal("Amérique", insee.Continent)
	assert.Equal(873, insee.OrderOfBirth)
	assert.Equal(40, insee.ControlKey)
	assert.Equal(insee.CountryCode, "432")
	assert.Equal(insee.CityCode, "")
}

func TestNewInseeData_ValidFrenchDROM(t *testing.T) {
	assert := assert.New(t)

	number := "299129742398791"
	insee, err := NewInseeData(number)
	if err != nil {
		t.Error(err)
	}
	valid, err := insee.IsValid()
	if err != nil {
		t.Error(err)
	}
	assert.True(valid)
	assert.Equal(number, insee.InseeNumber)
	assert.Equal(Female, insee.Gender)
	assert.Equal(1999, insee.Year)
	assert.Equal(time.Month(12), insee.Month)
	assert.Equal("La Réunion", insee.Department)
	assert.NotNil(insee.City)
	assert.Equal("LES TROIS-BASSINS", insee.City.Name)
	assert.Equal("97426", insee.City.ZipCode)
	assert.False(insee.Foreign)
	assert.Equal([]string{"FRANCE"}, insee.Countries)
	assert.Equal(987, insee.OrderOfBirth)
	assert.Equal(91, insee.ControlKey)
	assert.Equal(insee.CountryCode, "")
	assert.Equal(insee.CityCode, "23")
	assert.Equal("Europe", insee.Continent)
}

func TestNewInseeData_ValidFrenchCorsica(t *testing.T) {
	assert := assert.New(t)

	number := "299122A00498723"
	insee, err := NewInseeData(number)
	if err != nil {
		t.Error(err)
	}
	valid, err := insee.IsValid()
	if err != nil {
		t.Error(err)
	}
	assert.True(valid)
	assert.Equal(number, insee.InseeNumber)
	assert.Equal("Corse-du-Sud", insee.Department)
	assert.NotNil(insee.City)
	assert.Equal("AJACCIO", insee.City.Name)
	assert.Equal("20000", insee.City.ZipCode)
	assert.False(insee.Foreign)
	assert.Equal([]string{"FRANCE"}, insee.Countries)
	assert.Equal(23, insee.ControlKey)
}
