package main

import (
	"strconv"
	"strings"
	"unicode"
)

// transformer prend un string et le remanie en lui appliquant des fonctions de mise en forme ou de mise en page
func transformer(texte string) string {
	lignes := strings.Split(texte, "\n")

	for i, ligne := range lignes {
		ligne = strings.TrimRight(ligne, "\r")
		mots := strings.Fields(ligne)
		mots = appliquerMarqueurs(mots)
		mots = corrigerArticle(mots)
		lignes[i] = strings.Join(mots, " ")
	}

	return strings.Join(lignes, "\n")
}

// capitaliser met une majuscule au début du mot et le reste en mjnuscule
func capitaliser(mot string) string {
	if mot == "" {
		return mot
	}
	lettres := []rune(strings.ToLower(mot))
	lettres[0] = unicode.ToUpper(lettres[0])
	return string(lettres)
}

// appliquerMarqueurs détecte dans le string certains "marqueurs" prédéfinis et modifie le string en fonction d'eux
func appliquerMarqueurs(mots []string) []string {
	var resultat []string

	for i := 0; i < len(mots); i++ {
		mot := mots[i]
		if mot == "(up)" || mot == "(low)" || mot == "(cap)" || mot == "(hex)" || mot == "(bin)" {
			if len(resultat) == 0 {
				continue
			}

			dernier := len(resultat) - 1
			// puis on choisit la transformation selon le marqueur
			switch mot {
			case "(up)":
				resultat[dernier] = strings.ToUpper(resultat[dernier])
			case "(low)":
				resultat[dernier] = strings.ToLower(resultat[dernier])
			case "(cap)":
				resultat[dernier] = capitaliser(resultat[dernier])
			case "(hex)":
				valeur, err := strconv.ParseInt(resultat[dernier], 16, 64)
				if err != nil {
					continue
				}
				texte := strconv.FormatInt(valeur, 10)
				resultat[dernier] = texte

			case "(bin)":
				valeur, err := strconv.ParseInt(resultat[dernier], 2, 64)
				if err != nil {
					continue
				}
				texte := strconv.FormatInt(valeur, 10)
				resultat[dernier] = texte

			}
			// Section marqueur suivie d'un n
		} else if strings.HasPrefix(mot, "(up,") || strings.HasPrefix(mot, "(low,") || strings.HasPrefix(mot, "(cap,") {

			if i+1 >= len(mots) {
				continue
			}
			// je lui coupe la ) à la fin
			stringMarqueur := strings.TrimSuffix(mots[i+1], ")")

			// je convertit le string en int, que j'appelle nMarqueur
			nMarqueur, err := strconv.Atoi(stringMarqueur)
			i++

			// Pour montrer que le marqueur est incorrect, je le bazarde dans le texte tel quel
			if err != nil {
				resultat = append(resultat, mot)
			}

			// si le chiffre est plus gros que le texte, j'applique le marqueur sans planter
			if nMarqueur > len(resultat) {
				nMarqueur = len(resultat)
			}

			// je fais ma sauce en fonction du préfixe
			switch mot {
			case "(up,":

				// mon j parcourt resultat, il reviendra en arrière en fonction de n et ravancera jusqu'à là où on s'était arrêtés
				for j := len(resultat) - nMarqueur; j < len(resultat); j++ {
					resultat[j] = strings.ToUpper(resultat[j])
				}
			case "(low,":
				for j := len(resultat) - nMarqueur; j < len(resultat); j++ {
					resultat[j] = strings.ToLower(resultat[j])
				}
			case "(cap,":
				for j := len(resultat) - nMarqueur; j < len(resultat); j++ {
					resultat[j] = capitaliser(resultat[j])
				}
			}
		} else {
			resultat = append(resultat, mot)
		}
	}
	return resultat
}

func corrigerArticle(mots []string) []string {
	var voyelle bool = false
	for i := 0; i < len(mots)-1; i++ {
		// je prend le mot suivant, je le met en minuscule, je le convertit en chaine de runes
		runeMot := []rune(strings.ToLower(mots[i+1]))

		if runeMot[0] == 'a' || runeMot[0] == 'e' || runeMot[0] == 'i' || runeMot[0] == 'o' || runeMot[0] == 'u' || runeMot[0] == 'h' {
			voyelle = true
		}
		if mots[i] == "a" && voyelle == true {
			mots[i] = "an"
		}
		if mots[i] == "A" && voyelle == true {
			mots[i] = "An"
		}
		voyelle = false
	}
	return mots
}
