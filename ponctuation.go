package main

// estPonctuation détecte la ponctuation à partir d'une rune
func estPonctuation(r rune) bool {
	if r == '.' || r == ',' || r == '!' || r == '?' || r == ':' || r == ';' {
		return true
	}
	return false
}

// corrigerPonctuation prend un string et lui applique les règles typographiques de ponctuation demandées dans l'énoncé
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
