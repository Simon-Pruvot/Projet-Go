package main

import (
	"fmt"
)

func main() {
	player := &Characters{
		Nom:     "Héros",
		Hp:      70,
		HpMax:   100,
		Mana:    30,
		ManaMax: 50,
	}

	// Création d’un monstre
	monster := &Monsters{
		Nom:   "Squidward",
		Hp:    120,
		HpMax: 120,
	}
	combat(player, monster)
}

func combat(c *Characters, m *Monsters) {
	man := []string{
		"",
		"",
		"		  .-.",
		"	     (o.o)",
		"	     \\-_-/",
		"	     __|__",
		"	   // .|. \\",
		"	  //  .|.  \\",
		"	  \\  .|.  //",
		"	   \\(_-_)//",
		"	    (:| |:)",
		"	     || ||",
		"	     () ()",
		"	     || ||",
		"	     || ||",
		"	    ==' '==",
	}

	squidward := []string{
		"",
		"",
		"            .'/",
		"           / /",
		"          / /",
		"         / /",
		"        / /",
		"       / /",
		"      / /",
		"     / /",
		"    / /",
		"  __|/",
		" -\\__\\",
		"|f- Y\\|",
		"\\\\7L/",
		" cgD                            __ _",
		" |\\\\                          .'  Y '>,",
		"  \\ \\                        / _   _   \\",
		"   \\\\                       )(_) (_)(|}",
		"    \\\\                      {  4A   } /",
		"     \\\\                      \\uLuJJ/\\l",
		"      \\\\                     |3    p)/",
		"       \\\\___ __________      /nnm_n//",
		"       c7___-__,__-)\\,__)( .  \\_>-<_/D",
		"                  //V     \\_- ._.__G G_c__.-__< / ( \\",
		"                         < -._>__-,G_.___)\\   \\7\\",
		"                        ( -.__.| \\ <.__.- )   \\ \\",
		"                        | -.__ \\  | -.__.- .\\   \\ \\",
		"                        ( -.__  . \\ -.__.- .|    \\_\\",
		"                        \\ -.__ |!| -.__.- .)     \\ \\",
		"                          -.__  \\_| -.__.- ./      \\ l",
		"                           .__  >G>-.__.- >       .--,_",
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
	fmt.Printf("%-70s %s\n",
		fmt.Sprintf("PV: %d / %d | Mana: %d / %d", c.Hp, c.HpMax, c.Mana, c.ManaMax),
		fmt.Sprintf("PV: %d / %d | Mana: %d / %d", m.Hp, m.HpMax),
	)
	fmt.Println(`
_______________________________________________________________________________________________________________________________                                                                                      
	`)
}
