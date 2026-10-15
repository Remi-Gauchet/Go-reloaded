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
	for _, lettre := range runeLigne {
		if estPonctuation(lettre) == false {
			sortie = append(sortie, lettre)
		} else {
			for len(sortie) > 0 && sortie[len(sortie)-1] == ' ' {
				sortie = sortie[:len(sortie)-1]
			}
			sortie = append(sortie, lettre)
		}
	}
	return string(sortie)
}
