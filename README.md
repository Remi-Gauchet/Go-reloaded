# Go-reloaded
Programme servant à appliquer à un texte des modifications d'ordre typographique, en se servant de marqueurs incorporés au texte.

## Installation
Teléchargez le dossier.  
Invite de commande à la racine du projet.  
Dans votre éditeur de code, donnez en argument le nom du fichier à modifier, suivi du nom du fichier à créer.  

## Exemple d'utilisation

```bash
go run . entree.txt sortie .txt
```
**entree.txt** est le nom du fichier texte à modifier  
**sortie.txt** et le nom du fichier modifié

## Règles d'utilisation
|    Marqueur    |              Effet           |  Exemple   | Résultat |
|----------------|------------------------------|------------|----------|
|     `(up)`     |       tout en majuscule      | chat (up)  |   CHAT   |
|     `(low)`    |       tout en minuscule      | CHAT (low) |   chat   |
|     `(cap)`    | première lettre en majuscule | chat (cap) |   Chat   |
|    `(up, 'n')` | 'up' appliqué aux 'n' mots précédents | le rat chat (up, 2) | le RAT CHAT |
|   `(low, 'n')` | 'low' appliqué aux 'n' mots précédents | LE RAT CHAT (low, 2) | LE rat chat |
| `(cap, 'n')`   | 'cap' appliqué aux 'n' mots précédents | le rat chat (cap, 2) | le Rat Chat |

## Les messages d'erreur
Impossible de lire le fichier  
Impossible d'écrire le fichier  
Veuillez entrer le fichier d'entrée puis le fichier de sortie.  
Veillez à ce que le nom du fichier d'entrée et de sortie soient différents.  


## Les tests
Lancez l'input suivant :  
go run . test/ferre.txt sortie.txt  
Le fichier sortie.txt doit avoir le même contenu que le fichier ferre.resultat.txt  

Vous pouvez reproduire le test avec un autre fichier en suivant le même schéma.  


## Organisation du code
fichiers.go lit le fichier donné, et en crée un nouveau  
regles.go définit et applique les différentes modifications  
main.go lance le programme et appelle les autres fonctions  

## Limites
Votre fichier à modifier doit être sous le format .txt  
Les marqueurs doivent impérativement être séparés d'un espace avant et après votre texte.  

## Auteur
Rémi GAUCHET