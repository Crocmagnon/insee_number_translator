# INSEE number translator

Extract data from INSEE number (France)

## Getting started
### Using a pre-built binary
Go to the [latest release](https://git.augendre.info/gaugendre/insee_number_translator/releases/latest)
and download the binary matching your OS and architecture.

Then, run the tool via the command line:
```shell
./insee 269059913116714 168127982980507 299122A00498723 299129742398791 144089943287340
```

### From sources
#### Run
```shell
go run . 269059913116714 168127982980507 299122A00498723 299129742398791 144089943287340
```

#### Build & install
```shell
go install
insee_number_translator 269059913116714 168127982980507 299122A00498723 299129742398791 144089943287340
```

## Data sources

* https://www.insee.fr/fr/information/2560452, Millésime 2021 : Téléchargement des fichiers, CSV
