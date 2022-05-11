package main

import (
	"flag"
	"fmt"
	"os"

	"git.augendre.info/gaugendre/insee_number_translator/data"
)

func main() {
	flag.Usage = func() {
		out := flag.CommandLine.Output()
		fmt.Fprintf(out, "Usage: %s [flags] [numero_insee...]\n", os.Args[0])
		fmt.Fprintf(out, "\nCe programme décode les informations contenues dans votre numéro INSEE (numéro de sécurité sociale français) ")
		fmt.Fprintf(out, "et vous les affiche d'une manière lisible et claire.\n")
		flag.PrintDefaults()
		fmt.Fprintf(out, "\nLes arguments numero_insee doivent comporter 15 caractères. Il est possible d'en spécifier plusieurs séparés par un espace.\n")
	}

	flag.Parse()

	numbers := flag.Args()
	if len(numbers) == 0 {
		flag.Usage()
		return
	}

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
