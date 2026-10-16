package main

import "unicode"

// estPonctuation détecte la ponctuation à partir d'une rune
func estPonctuation(r rune) bool {
	if r == '.' || r == ',' || r == '!' || r == '?' || r == ':' || r == ';' {
		return true
	}
	return false
}

// estApostropheDeMot nous permet de différencier les ' utilisés pour une citation des ' inclus dans la grammaire
func estApostropheDeMot(runeLigne []rune, i int) bool {
	if i > 0 && i < len(runeLigne)-1 {
		if unicode.IsLetter(runeLigne[i-1]) == true && unicode.IsLetter(runeLigne[i+1]) == true {
			return true
		}
	}
	return false
}

// corrigerPonctuation prend le string, puis lui applique les règles typographiques de ponctuation, puis le renvoie ainsi modifié
func corrigerPonctuation(ligne string) string {
	var sortie []rune
	runeLigne := []rune(ligne)
	for i, lettre := range runeLigne {
		if estPonctuation(lettre) == false {
			sortie = append(sortie, lettre)
		} else {
			for len(sortie) > 0 && sortie[len(sortie)-1] == ' ' {
				sortie = sortie[:len(sortie)-1]
			}
			sortie = append(sortie, lettre)

			// je vais appeler runeLigne[i+1] donc j'applique une protection pour pas planter si on est arrivé au bout
			if i+1 >= len(runeLigne) {
				return string(sortie)
			}

			// je regarde la rune suivante pour voir ce que je dois faire
			if runeLigne[i+1] == ' ' {
				continue
			}
			if estPonctuation(runeLigne[i+1]) == true {
				continue
			}
			sortie = append(sortie, ' ')
		}
	}
	return string(sortie)
}

// corrigerApostrophes reprend le string, cherche un apostrophe, et retire si besoin les espaces en trop d'une citation
func corrigerApostrophes(ligne string) string {
	var sortie []rune
	var citation bool = false
	runeLigne := []rune(ligne)
	for i, lettre := range runeLigne {
		// je détecte mon apostrophe
		if lettre == '\'' {

			//fin de citation
			if citation == true {
				if estApostropheDeMot(runeLigne, i) == false {
					for len(sortie) > 0 && sortie[len(sortie)-1] == ' ' {
						sortie = sortie[:len(sortie)-1]
					}
					sortie = append(sortie, lettre)
					citation = false

					//ce n'est pas la fin de la citation
				} else {
					sortie = append(sortie, lettre)
				}

			} else if estApostropheDeMot(runeLigne, i) == false {
				//début de citation
				sortie = append(sortie, lettre)
				citation = true

			} else {
				// ce n'est pas une citation
				sortie = append(sortie, lettre)
			}
		} else if lettre == ' ' && citation == true && len(sortie) > 0 && sortie[len(sortie)-1] == '\'' {
			continue
		} else {
			sortie = append(sortie, lettre)
		}
	}
	return string(sortie)
}
