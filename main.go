package main

import (
	"flag"
	"fmt"
	"insee_number_translator/data"
)

func main() {
	flag.Parse()
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
