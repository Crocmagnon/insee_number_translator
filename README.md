# INSEE number translator

Extract data from INSEE number (France)

## Getting started

```shell
pyenv virtualenv 3.9.6 insee
pyenv local insee
poetry install
python insee_number_translator/main.py
python insee_number_translator/main.py 123456789123456
```

## Data sources

- cities : https://public.opendatasoft.com/explore/dataset/correspondance-code-insee-code-postal/export/
- countries : https://www.insee.fr/fr/information/2028273
- departments : https://www.data.gouv.fr/fr/datasets/regions-departements-villes-et-villages-de-france-et-doutre-mer/
