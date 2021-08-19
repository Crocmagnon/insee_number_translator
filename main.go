package main

import (
	"flag"
	"fmt"
	"insee_number_translator/data"
)

func main() {
	var preProcess bool
	flag.BoolVar(&preProcess, "pre-process", false, "Pre-process data files.")
	flag.Parse()

	if preProcess {
		data.PreProcessRawData("data/raw_data", "data/curated_data")
		return
	}

	numbers := flag.Args()
	for _, number := range numbers {
		insee, err := data.NewInseeData(number)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(insee)
		}
		fmt.Println()
	}
}
