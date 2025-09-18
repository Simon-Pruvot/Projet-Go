package main

import (
	"fmt"
)

func combat0(c *Character, m *Monster) {
	Clear()
	man := []string{
		"              .-.",
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
		"                     ,      ,",
		"            /(.-\"\"-.)\\",
		"        |\\  \\/      \\/  /|",
		"        | \\ / =.  .= \\ / |",
		"        \\( \\   o\\/o   / )/",
		"         \\_, '-/  \\-' ,_/",
		"           /   \\__/   \\",
		"           \\ \\__/\\__/ /",
		"         ___\\ \\|--|/ /___",
		"       /`    \\      /    `\\",
		"      /       '----'       \\",
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
		fmt.Sprintf("PV: %d / %d", m.Hp, m.HpMax),
	)
	fmt.Println(`
_______________________________________________________________________________________________________________________________                                                                                      
	`)
}
