package data

import "testing"

func TestNewInseeDataValid(t *testing.T) {
	number := "269059913116714"
	insee, err := NewInseeData(number)
	if err != nil {
		t.Error(err)
	}
	if insee.InseeNumber != number {
		t.Errorf("InseeNumber KO. Got %s, expected %s", insee.InseeNumber, number)
	}
	foreign := true
	if insee.Foreign != foreign {
		t.Errorf("Foreign KO. Got %t, expected %t", insee.Foreign, foreign)
	}
	gender := "Female"
	if insee.Gender != gender {
		t.Errorf("Gender KO. Got %s, expected %s", insee.Gender, gender)
	}
}
