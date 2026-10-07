package main

import (
	"fmt"
	"os"

	"github.com/Remi-Gauchet/Go-reloaded.git/capitalize"
)

func main() {
	sample, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	resultat := capitalize.Capitalize(string(sample))
	err = os.WriteFile("result.txt", []byte(resultat), 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
