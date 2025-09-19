package main

import (
	"fmt"
)

func CreateClasse() Character {
	fmt.Println(`
	"",	                                                                                           
 _____ _       _     _                            _                  _                 _ 
|     | |_ ___|_|___|_|___ ___ ___ ___    _ _ ___| |_ ___ ___    ___| |___ ___ ___ ___|_|  
|   --|   | . | |_ -| |_ -|_ -| -_|- _|  | | | . |  _|  _| -_|  |  _| | .'|_ -|_ -| -_|_
|_____|_|_|___|_|___|_|___|___|___|___|  \\_/|___|_| |_| |___|  |___|_|__,|___|___|___|_|  
     

	`)
	elfe := []string{
		"",
		"           .-----.",
		" \\ ' /   _/    )/",
		"- ( ) -('---''--)",
		" / . \\((()\\^_^/)()",
		"  \\_\\ (()_)-((()()",
		"   '- \\ )/\\._./(()",
		"      /\\/( X   ) \\",
		"     (___)|___/   \\",
		"          |.#_|(___)",
		"         /\\    \\ ( (_",
		"         \\/\\/\\/\\) \\",
		"         | / \\ |",
		"         |(   \\|",
		"        _|_)__|_\\_",
		"        )...()...(",
		"         | (   \\ |",
		"      .-'__,)  (  \\",
		"                '\\_-,",
		"",
		"",
	}

	man := []string{
		"",
		"",
		"      ////^\\\\\\\\",
		"      | ^   ^ |",
		"     @ (o) (o) @",
		"      |   <   |",
		"      |  ___  |",
		"       \\_____/",
		"     ____|  |____",
		"    /    \\__/    \\",
		"   /              \\",
		"  /\\_/|        |\\_/\\",
		" / /  |        |  \\ \\",
		"( <   |        |   > )",
		" \\ \\  |        |  / /",
		"  \\ \\ |________| / /",
		"",
		"",
	}

	nain := []string{
		"",
		"",
		"",
		"",
		"",
		"",
		"   ,====,",
		"  c , _,{",
		"  /\\  @ )                 __",
		" /  ^~~^\\          <=.,__/ '}=",
		"(_/ ,, ,,)          \\_ _>_/~",
		" ~\\_(/-\\)'-,_,_,_,-'(_)-(_)",
		"",
		"",
	}

	maxLines := len(elfe)
	if len(man) > maxLines {
		maxLines = len(man)
	}
	if len(nain) > maxLines {
		maxLines = len(nain)
	}

	for i := 0; i < maxLines; i++ {
		col1 := ""
		if i < len(elfe) {
			col1 = elfe[i]
		}

		col2 := ""
		if i < len(man) {
			col2 = man[i]
		}

		col3 := ""
		if i < len(nain) {
			col3 = nain[i]
		}
		fmt.Printf("%-40s %-40s %-40s\n", col1, col2, col3)
	}

	fmt.Println(`


   _ _____ _ ___                          _ _____               _                          _ _____     _     
  / |   __| |  _|___                     / |  |  |_ _ _____ ___|_|___                     / |   | |___|_|___ 
 / /|   __| |  _| -_|                   / /|     | | |     | .'| |   |                   / /| | | | .'| |   |
|_/ |_____|_|_| |___|                  |_/ |__|__|___|_|_|_|__,|_|_|_|                  |_/ |_|___|__,|_|_|_|
                                                                                                            `)
	var class string
	fmt.Scanln(&class)
	if class != "/elfe" && class != "/humain" && class != "/nain" || class != "/ELFE" && class != "/HUMAIN" && class != "/NAIN" {
	} else {
		return CreateClasse()
	}

	fmt.Printf("Hi %s!\n", class)
	return Perso(class)
}

func Perso(classe string) Character {
	Clear()
	fmt.Println(`
	
		 _____                                    _              _____             _
		|     |___ ___ ___ ___ ___ ___    _ _ ___| |_ ___ ___   |   | |___ _____  |_|
		|-   -|   |_ -| -_|  _| -_|- _|  | | | . |  _|  _| -_|  | | | | . |     |  _
		|_____|_|_|___|___|_| |___|___|  \\_/|___|_| |_| |___|  |_|___|___|_|_|_| |_|
	`)
	var Nom string
	fmt.Scanln(&Nom)
	fmt.Printf("Hi %s!\n", Nom)

	switch classe {
	case "/Elfe", "/elfe", "ELFE":
		return initCharacter(Nom, "Elfe", 1, 80, 80, []Objects{{"Potion de vie", 3}}, []string{}, Equipment{})
	case "/Humain", "/humain", "/HUMAIN":
		return initCharacter(Nom, "Humain", 1, 100, 100, []Objects{{"Potion de vie", 3}}, []string{}, Equipment{})
	case "/Nain", "/nain", "/NAIN":
		return initCharacter(Nom, "Nain", 1, 120, 120, []Objects{{"Potion de vie", 3}}, []string{}, Equipment{})
	default:
		return CreateClasse()
	}
}
