# Go-reloaded
Programme servant à appliquer à un texte des modifications d'ordre typographique, en se servant de marqueurs incorporés au texte.

## Installation
Teléchargez le dossier.  
Invite de commande à la racine du projet.  
Dans votre éditeur de code, donnez en argument le nom du fichier à modifier, suivi du nom du fichier à créer.  


## Exemple d'utilisation

```bash
go run . entree.txt sortie.txt
```
**entree.txt** est le nom du fichier texte à modifier  
**sortie.txt** et le nom du fichier modifié


## Règles d'utilisation
|    Marqueur    |              Effet           |  Exemple   | Résultat |
|----------------|------------------------------|------------|----------|
|     `(up)`     | met le mot précédent en majuscule | chat (up)  |   CHAT   |
|     `(low)`    | met le mort précédent en minuscule | CHAT (low) |   chat   |
|     `(cap)`    | première lettre du mot précédent en majuscule | chat (cap) |   Chat   |
|    `(up, n)` | *up* appliqué aux *n* mots précédents (avec n = nombre entier) | le chien loup (up, 2) | le CHIEN LOUP |
|   `(low, n)` | *low* appliqué aux *n* mots précédents (avec n = nombre entier) | LE CHIEN LOUP (low, 2) | LE chien loup |
| `(cap, n)`   | *cap* appliqué aux *n* mots précédents (avec n = nombre entier) | le chien loup (cap, 2) | le Chien Loup |
| `(bin)` | remplace le mot précédent par sa valeur en base 10 | 10 (bin) ans | 2 ans |
| `(hex)`| remplace le mot précédent par sa valeur en base 10 | 1E (hex) fichiers | 30 fichiers |


## Les messages d'erreur
Impossible de lire le fichier  
Impossible d'écrire le fichier  
Veuillez entrer le fichier d'entrée puis le fichier de sortie.  
Veuillez à ce que le nom du fichier d'entrée et de sortie soient différents.  


## Les tests
Lancez l'input suivant :  
go run . test/ferre.txt sortie.txt  
Le fichier sortie.txt doit avoir le même contenu que le fichier ferre.resultat.txt  

Vous pouvez reproduire le test avec un autre fichier en suivant le même schéma.  


## Organisation du code
fichiers.go lit le fichier donné, et en crée un nouveau  
regles.go définit différentes règles et les applique au texte pour le modifier  
ponctuation.go rajoute à regles.go des règles portant sur la ponctuation  
main.go lance le programme, appelle les autres fonctions, ou affiche un message d'erreur  


## Limites
Votre fichier à modifier doit être sous le format .txt  
Les marqueurs doivent impérativement être séparés d'un espace avant et après votre texte.  
Un marqueur à l'entrée du texte est supprimé du texte, sans ne rien faire.  
Les fautes de frappe dans les marqueurs ne sont pas corrigées.  


## Auteur
Rémi GAUCHET