package data

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func PreProcessRawData(sourceFolder, targetFolder string) {
	err := os.MkdirAll(targetFolder, 0755)
	if err != nil {
		fmt.Printf("Error, couldn't create target folder %s: %s", targetFolder, err)
		return
	}
	err = preProcessCities(sourceFolder+"/commune2021.csv", targetFolder+"/cities.json")
	if err != nil {
		fmt.Println("Error during cities pre processing")
		fmt.Println(err.Error())
	}
	err = preProcessDepartments(sourceFolder+"/departement2021.csv", targetFolder+"/departments.json")
	if err != nil {
		fmt.Println("Error during departments pre processing")
		fmt.Println(err.Error())
	}
	err = preProcessCountries(sourceFolder+"/pays2021.csv", targetFolder+"/countries.json")
	if err != nil {
		fmt.Println("Error during countries pre processing")
		fmt.Println(err.Error())
	}
}

func preProcessCities(sourceFileName, targetFileName string) error {
	return preProcessSimpleFile(sourceFileName, targetFileName, "COM")
}

func preProcessDepartments(sourceFileName, targetFileName string) error {
	return preProcessSimpleFile(sourceFileName, targetFileName, "DEP")
}

func preProcessSimpleFile(sourceFileName, targetFileName, codeColumn string) error {
	sourceFile, err := os.Open(sourceFileName)
	if err != nil {
		return err
	}
	defer sourceFile.Close()
	reader := csv.NewReader(sourceFile)
	line, err := reader.Read()
	if err != nil {
		return err
	}
	columnsToIndex := map[string]int{codeColumn: -1, "NCCENR": -1}

	for index, header := range line {
		_, present := columnsToIndex[header]
		if present {
			columnsToIndex[header] = index
		}
	}
	for column, index := range columnsToIndex {
		if index == -1 {
			return fmt.Errorf("column %s is missing in file. Found columns: %s", column, line)
		}
	}

	sourceItems := make(map[string]string)

	lines, err := reader.ReadAll()
	if err != nil {
		return err
	}
	codeIndex := columnsToIndex[codeColumn]
	nameIndex := columnsToIndex["NCCENR"]
	for _, line = range lines {
		code := line[codeIndex]
		name := line[nameIndex]
		sourceItems[code] = name
	}

	itemsJson, err := json.Marshal(sourceItems)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(targetFileName, itemsJson, 0644)
	if err != nil {
		return err
	}

	return nil
}

func preProcessCountries(sourceFileName, targetFileName string) error {
	sourceFile, err := os.Open(sourceFileName)
	if err != nil {
		return err
	}
	defer sourceFile.Close()
	reader := csv.NewReader(sourceFile)
	line, err := reader.Read()
	if err != nil {
		return err
	}
	columnsToIndex := map[string]int{"COG": -1, "LIBCOG": -1}

	for index, header := range line {
		_, present := columnsToIndex[header]
		if present {
			columnsToIndex[header] = index
		}
	}
	for column, index := range columnsToIndex {
		if index == -1 {
			return fmt.Errorf("column %s is missing in file. Found columns: %s", column, line)
		}
	}

	sourceCountries := make(map[string][]string)

	lines, err := reader.ReadAll()
	if err != nil {
		return err
	}
	codeIndex := columnsToIndex["COG"]
	nameIndex := columnsToIndex["LIBCOG"]
	for _, line = range lines {
		code := line[codeIndex]
		name := line[nameIndex]
		current, present := sourceCountries[code]
		if !present {
			current = []string{}
		}
		current = append(current, name)
		sourceCountries[code] = current
	}

	citiesJson, err := json.Marshal(sourceCountries)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(targetFileName, citiesJson, 0644)
	if err != nil {
		return err
	}

	return nil
}
