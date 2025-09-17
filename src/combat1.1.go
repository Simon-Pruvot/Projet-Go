package main

import (
	"fmt"
)

func main() {
	Menu()
}

func Menu() {
	man := []string{
		"",
		"",
		"            ↑↑↑↑↑↑↑",
		"            ↑     ↑",
		"            ↑  ↑↑↑↑",
		"          ↑↑↑      ↑↑",
		"         ↑↑↑   ↑↑↑↑",
		"         ↑↑         ↑",
		"         ↑     ↑↑↑↑",
		"        ↑↑    ↑↑↑↑↑↑",
		"         ↑    ↑↑",
		"         ↑↑    ↑↑",
		"         ↑↑     ↑↑",
		"        ↑↑  ↑↑↑  ↑↑",
		"       ↑↑↑  ↑↑↑↑  ↑↑",
		"      ↑↑↑  ↑↑  ↑↑  ↑",
		"     ↑↑↑ ↑↑↑   ↑↑  ↑↑",
		"     ↑  ↑↑↑    ↑↑  ↑",
		"     ↑↑↑↑       ↑↑↑↑",
	}

	squidward := []string{

		"           .,--.",
		"         .' __  \\",
		"         | .._  |",
		"         |{)(} .'",
		"         / /|  |.",
		"        (_/ /____)",
		"          |_||",
		"            /'",
		"            //",
		"          .'''\\",
		"         /\\:::/\\",
		"        ( /|::|\\\\",
		"        _\\:|;;|{/_",
		"        '.;|**|\\;,/",
		"           \\_ /",
		"           | ||",
		"           | ||",
		"           | ||",
		"           | ||",
		"         ._| ||_.",
		"        ;,_.-._,;",
	}

	maxLines := len(man)
	if len(squidward) > maxLines {
		maxLines = len(squidward)
	}

	for i := 0; i < maxLines; i++ {
		left := ""
		if i < len(man) {
			left = man[i]
		}

		right := ""
		if i < len(squidward) {
			right = squidward[i]
		}

		// Pad left column to fixed width
		fmt.Printf("%-70s %s\n", left, right)
	}
	fmt.Println(`
_______________________________________________________________________________________________________________________________                                                                                      
	`)
}
