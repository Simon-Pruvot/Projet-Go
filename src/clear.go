package main

import "fmt"

func Clear() {
	fmt.Print("\033[H\033[2J")              // clear avec ANSI
	fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n") // ajoute de l’espace supplémentaire
}
