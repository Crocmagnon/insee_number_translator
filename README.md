# INSEE number translator

Extract data from INSEE number (France)

## Getting started
### From sources
#### Run
```shell
go run . 269059913116714 168127982980507
```

#### Build & install
```shell
go install
insee_number_translator 269059913116714 168127982980507
```

## Data sources

* cities : https://public.opendatasoft.com/explore/dataset/correspondance-code-insee-code-postal/export/
* countries : https://www.insee.fr/fr/information/2028273
* departments : https://www.data.gouv.fr/fr/datasets/regions-departements-villes-et-villages-de-france-et-doutre-mer/
