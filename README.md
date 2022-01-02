# INSEE number translator (Français)
(English below)
Extrait les données des numéros INSEE (numéro de sécurité sociale française).

Le numéro de sécurité sociale français est *signifiant*. Cela veut dire que ce n'est pas un numéro aléatoire,
il porte des informations.

* https://www.insee.fr/fr/metadonnees/definition/c1409
* https://fr.wikipedia.org/wiki/Num%C3%A9ro_de_s%C3%A9curit%C3%A9_sociale_en_France

## Commencer
### Utiliser un programme prêt à l'emploi
Rendez-vous sur la [dernière version](https://git.augendre.info/gaugendre/insee_number_translator/releases/latest)
et téléchargez le fichier correspondant à votre système d'exploitation et à votre architecture.  
Si vous ne savez pas de quoi il s'agit, essayez d'abord `amd64` puis `386` si le premier ne fonctionne pas.  
Pour macOS, utilisez le fichier `darwin`.

Ensuite, utilisez la ligne de commande pour exécuter le fichier :
```shell
./insee 269059913116714 168127982980507 299122A00498723 299129742398791 144089943287340
```

Rassurez-vous, ce programme ne transmet aucune information sur le réseau. Toutes les données
nécessaires à son fonctionnement sont comprises dans le fichier que vous venez de télécharger.  
Votre numéro de sécurité sociale ne quitte pas votre ordinateur.

### Depuis les sources
Requiert les outils `go` (https://golang.org/).

#### Exécution directe
```shell
go run . 269059913116714 168127982980507 299122A00498723 299129742398791 144089943287340
```

#### Compilation et installation
```shell
go install
insee_number_translator 269059913116714 168127982980507 299122A00498723 299129742398791 144089943287340
```

# INSEE number translator (English)
Extract data from INSEE number (France).

The french social security number is significant. That means it's not a random number, it bears information.

* (fr) https://www.insee.fr/fr/metadonnees/definition/c1409
* (fr) https://fr.wikipedia.org/wiki/Num%C3%A9ro_de_s%C3%A9curit%C3%A9_sociale_en_France

## Getting started
### Using a pre-built binary
Go to the [latest release](https://git.augendre.info/gaugendre/insee_number_translator/releases/latest)
and download the binary matching your OS and architecture.  
If you don't know what that means, try `amd64` first, then `386` if it doesn't work.  
For macOS, choose `darwin`.

Then, run the tool via the command line:
```shell
./insee 269059913116714 168127982980507 299122A00498723 299129742398791 144089943287340
```
Rest assured, this program doesn't transmit any information on the network. All the data
necessary for its operation are included in the file you just downloaded.  
Your social security number doesn't leave your computer.

### From sources
Requires the `go` toolchain (https://golang.org/).

#### Run
```shell
go run . 269059913116714 168127982980507 299122A00498723 299129742398791 144089943287340
```

#### Build & install
```shell
go install
insee_number_translator 269059913116714 168127982980507 299122A00498723 299129742398791 144089943287340
```

# Tinker
## Update data
You can easily update the data by downloading the CSV files available using the link in `Data sources` below.
Then, unzip the file in `data/raw_data` and run the following command:

```shell
invoke pre-process
```

This requires [`invoke`](https://www.pyinvoke.org/) on your machine (run `pip install -r requirements.txt`).

## Release
Run tests, create a tag and build binaries:
```
inv release <version_name>
```

# Data sources

* https://www.insee.fr/fr/information/2560452, Millésime 2021 : Téléchargement des fichiers, CSV

# Reuse
If you do reuse my work, please consider linking back to this repository 🙂