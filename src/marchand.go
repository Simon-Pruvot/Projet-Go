package main

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

func (c *Character) Marchand() {
	building := []string{
		"     .     .",
		"                               !!!!!!!",
		"                       .       [[[|]]]    .",
		"                       !!!!!!!!|--_--|!!!!!",
		"                       [[[[[[[[\\_(X)_/]]]]]",
		"               .-.     /-_--__-/_--_-\\-_--\\",
		"               |=|    /-_---__/__-__-\\__-_-\\",
		"           . . |=| ._/-__-__\\===========/-__\\_",
		"           !!!!!!!!!\\========[ /]]|[[\\ ]=====/",
		"          /-_--_-_-_[[[[[[[[[||==  == ||]]]]]]",
		"         /-_--_--_--_|=  === ||=/^|^\\ ||== =|",
		"        /-_-/^|^\\-_--| /^|^\\=|| | | | ||^\\= |",
		"       /_-_-| | |-_--|=| | | ||=|_|_|=||\"|==|",
		"      /-__--|_|_|_-_-| |_|_|=||______=||_| =|",
		"     /_-__--_-__-___-|_=__=_.`---------'._=_|__",
		"    /-----------------------\\===========/-----/",
		"   ^^^\\^^^^^^^^^^^^^^^^^^^^^^[[|]]|[[|]]=====/",
		"       |.' ..==::'\"'::==.. '.[ /~~~~~\\ ]]]]]]]",
		"       | .'=[[[|]]|[[|]]]=`._||==  =  || =\\ ]",
		"       ||== =|/ _____ \\|== = ||=/^|^\\=||^\\ ||",
		"       || == `||-----||' = ==|| | | |=|| |=||",
		"       ||= == ||:^s^:|| = == ||=| | | || |=||",
		"       || = = ||:___:||= == =|| |_|_| ||_|=||",
		"      _||_ = =||o---.|| = ==_||_= == =||==_||_",
		"      \\__/= = ||:   :||= == \\__/[][][][][]\\__/",
		"      [||]= ==||:___:|| = = [||]\\\\//\\\\//\\\\[||]",
		"      }  {---'\"'-----'\"'- --}  {//\\\\//\\\\//}  {",
		"    __[==]__________________[==]\\\\//\\\\//\\\\[==]_",
		"   |`|~~~~|================|~~~~|~~~~~~~~|~~~~||",
		"jgs|^| ^  |================|^   | ^ ^^ ^ |  ^ ||",
		"  \\|//\\/^|/==============\\|/^\\\\\\^/^\\.\\^///\\\\//|///",
		" \\\\///\\\\\\//===============\\\\//\\\\///\\\\\\\\////\\\\\\/////",
		" \"\"'\"\"'\"\"\".'..'. ' '. ''..'.'\"\"'\"\"'\"\"'\"\"''\"''\"''\"\"",
	}

	//objects
	hpPot := Objects{nom: "Potion de vie", quantity: 1}
	poisonPot := Objects{nom: "Potion de poison", quantity: 1}
	swordCom := Objects{nom: "Épée C", quantity: 1}
	swordRare := Objects{nom: "Épée B", quantity: 1}

	armorCom := Objects{nom: "Armor C", quantity: 1}
	armorRare := Objects{nom: "Armor B", quantity: 1}

	LivFeu := Objects{nom: "Livre de Sort : Boule de feu", quantity: 1}
	LivCoup := Objects{nom: "Livre : Coup de poing", quantity: 1}

	//resurces
	rock := Objects{nom: "Rock", quantity: 1}
	wood := Objects{nom: "Wood", quantity: 1}
	scrap := Objects{nom: "Scrap", quantity: 1}
	FOL := Objects{nom: "Fourrure de Loup", quantity: 1}
	PDT := Objects{nom: "Peau de Troll", quantity: 1}
	CDS := Objects{nom: "Cuir de Sanglier", quantity: 1}
	PDC := Objects{nom: "Plume de Corbeau", quantity: 1}

	// Liste des objets à acheter
	achats := []struct {
		key   string
		nom   string
		price int
	}{
		{"1", "Potion de vie", 3},
		{"2", "Potion de poison", 6},
		{"3", "Épée C", 10},
		{"5", "Armure C", 10},
		{"6", "Épée B", 20},
		{"7", "Armure B", 20},
		{"8", "Livre de Sort : Boule de feu", 25},
		{"9", "Augmentation d’inventaire", 30},
	}

	// Liste des objets à vendre
	ventes := []struct {
		key   string
		nom   string
		price int
	}{
		{"a/A", "Rock", 1},
		{"b/B", "Wood", 1},
		{"c/C", "Scrap", 5},
		{"d/D", "Fourrure de Loup", 4},
		{"e/E", "Peau de Troll", 7},
		{"f/F", "Cuir de Sanglier", 3},
		{"g/G", "Plume de Corbeau", 1},
		{"1", "Potion de vie", 1},
		{"2", "Potion de poison", 1},
		{"3", "Épée C", 5},
		{"5", "Armure C", 5},
		{"7", "Armure B", 20},
	}

MarchandLoop:

	for {
		merchantmenu := []string{
			"===================================",
			"Bienvenue chez le marchand !       ",
			"                                   ",
			"   B) Acheter                      ",
			"   S) Vendre                       ",
			"   Q) Quitter le marchand          ",
			"                                   ",
			"===================================",
		}
		lines := CombineColumnsToLines([][]string{building, merchantmenu}, 4) // returns []string
		FullScreenDrawCentered(lines)

		char, _, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		switch char {
		case 'b', 'B':
			// build content lines for achats (keep same information/format as before)
			content := []string{
				"",
				"Objets disponibles à l'achat :",
				"    ________________________________________________",
				"   / \\                                               \\.",
				"  |   |                   Marchand                   |.",
				"   \\_ |                                              |.",
			}

			for _, it := range achats {
				content = append(content, fmt.Sprintf("      |   %-4s %-28s %3d Gold |", it.key, it.nom, it.price))
			}

			content = append(content,
				"      |                                              |.",
				"      |   ___________________________________________|",
				"      |  /                                           /.",
				"      \\_/__________________________________________/.",
				"",
				"Chose an item: 1, 2, 3, 5, 6, 7, 8",
				"Q to leave the marchant",
			)

			lines := CombineColumnsToLines([][]string{building, content}, 4) // returns []string
			FullScreenDrawCentered(lines)

		AchatLoop:
			for {
				char, _, err := keyboard.GetKey()
				if err != nil {

				}

				switch char {
				case '1', '&': // Acheter potion de vie

					if c.Money >= 3 {
						c.Money -= 3
						added := false
						for i := range c.inv {
							if c.inv[i].nom == "Potion de vie" {
								c.inv[i].quantity++
								fmt.Println("✅")
								added = true
								break
							}
						}
						if !added {
							c.addInventory(hpPot)
							fmt.Println("✅")
						}
					} else {
						fmt.Println("Need More Gold")
					}

				case '2', 'é': // Acheter potion de poison
					if c.Money >= 6 {
						c.Money -= 6
						added := false
						for i := range c.inv {
							if c.inv[i].nom == "Potion de poison" {
								c.inv[i].quantity++
								fmt.Println("✅")
								added = true
								break
							}
						}
						if !added {
							c.addInventory(poisonPot)
							fmt.Println("✅")
						}
					} else {
						fmt.Println("Need More Gold")
					}

				case '3', '"': // Acheter Sword C
					if c.Money >= 10 {
						c.Money -= 10
						added := false
						for i := range c.inv {
							if c.inv[i].nom == "Épée C" {
								c.inv[i].quantity++
								fmt.Println("✅")
								added = true
								break
							}
						}
						if !added {
							c.addInventory(swordCom)
							fmt.Println("✅")
						}
					} else {
						fmt.Println("Need More Gold")
					}

				case '5', '(': // Acheter Armor C
					if c.Money >= 10 {
						c.Money -= 10
						added := false
						for i := range c.inv {
							if c.inv[i].nom == "Armor C" {
								c.inv[i].quantity++
								fmt.Println("✅")
								added = true
								break
							}
						}
						if !added {
							c.addInventory(armorCom)
							fmt.Println("✅")
						}
					} else {
						fmt.Println("Need More Gold")
					}

				case '6', '-': // Acheter Sword B
					if c.Money >= 20 {
						c.Money -= 20
						added := false
						for i := range c.inv {
							if c.inv[i].nom == "Épée B" {
								c.inv[i].quantity++
								fmt.Println("✅")
								added = true
								break
							}
						}
						if !added {
							c.addInventory(swordRare)
							fmt.Println("✅")
						}
					} else {
						fmt.Println("Need More Gold")
					}

				case '7', 'è': // Acheter Armor B
					if c.Money >= 20 {
						c.Money -= 20
						added := false
						for i := range c.inv {
							if c.inv[i].nom == "Armor B" {
								c.inv[i].quantity++
								fmt.Println("✅")
								added = true
								break
							}
						}
						if !added {
							c.addInventory(armorRare)
							fmt.Println("✅")
						}
					} else {
						fmt.Println("Need More Gold")
					}

				case '8', '_': // Acheter Livre de Sort
					if c.Money >= 25 {
						c.Money -= 25
						added := false
						for i := range c.inv {
							if c.inv[i].nom == "Livre de Sort : Boule de feu" {
								c.inv[i].quantity++
								fmt.Println("✅")
								added = true
								break
							}
						}
						if !added {
							c.addInventory(LivFeu)
							fmt.Println("✅")
						}
					} else {
						fmt.Println("Need More Gold")
					}
				case '9', 'ç': // Augmentation d’inventaire
					if c.Money >= 30 {
						c.Money -= 30
						c.upgradeInventorySlot()
						fmt.Println("✅")
					} else {
						fmt.Println("💰 Pas assez d’or pour l’augmentation !")
					}
				case 'q', 'Q':
					fmt.Println("Au revoir !")
					break AchatLoop

				case ')', '°': // Acheter Livre de Sort
					if c.Money >= 25 {
						c.Money -= 25
						added := false
						for i := range c.inv {
							if c.inv[i].nom == "Livre : Coup de poing" {
								c.inv[i].quantity++
								fmt.Println("✅")
								added = true
								break
							}
						}
						if !added {
							c.addInventory(LivCoup)
							fmt.Println("✅")
						}
					} else {
						fmt.Println("Need More Gold")
					}

				}
			}

		case 's', 'S':
			// build content lines for ventes
			contentV := []string{
				"",
				"Objets que vous pouvez vendre :",
				"    _________________________________________",
				"   / \\                                       \\.",
				"  |   |               Revente                |.",
				"   \\_ |                                      |.",
			}

			for _, it := range ventes {
				contentV = append(contentV, fmt.Sprintf("      |  %-4s %-22s %2d Gold |", it.key, it.nom, it.price))
			}

			contentV = append(contentV,
				"      |                                      |.",
				"      |   ___________________________________|",
				"      |  /                                   /.",
				"      \\_/___________________________________/.",
				"",
				"Appuyez sur la touche correspondant à l’objet pour le vendre, ou Q pour revenir.",
				"Chose an item: 1, 2, 3, 5, 6, 7",
				"Q to leave the marchant",
			)

			// print building left and vente list right
			lines := CombineColumnsToLines([][]string{building, contentV}, 4) // returns []string
			FullScreenDrawCentered(lines)

		VenteLoop:
			for {
				char, _, err := keyboard.GetKey()
				if err != nil {
					log.Fatal(err)
				}

				switch char {
				case '1', '&':
					for i := 0; i < len(c.inv); i++ {
						if c.inv[i].nom == "Potion de vie" && c.inv[i].quantity > 0 {
							c.inv[i].quantity--
							fmt.Println("✅")
							c.Money++
						}
						if c.inv[i].nom == "Potion de vie" && c.inv[i].quantity == 0 {
							c.removeInventory(c.inv[i])
							fmt.Println("✅")
						}
					}
				case '2', 'é': // Acheter une potion de poison pour 3 Gold
					for i := 0; i < len(c.inv); i++ {
						if c.inv[i].nom == "Potion de poison" && c.inv[i].quantity > 0 {
							c.inv[i].quantity--
							fmt.Println("✅")
							c.Money++
						}
						if c.inv[i].nom == "Potion de poison" && c.inv[i].quantity == 0 {
							c.removeInventory(c.inv[i])
							fmt.Println("✅")
						}
					}

				case '3', '"': // Acheter Sword C pour 10 Gold
					for i := 0; i < len(c.inv); i++ {
						if c.inv[i].nom == "Épée C" && c.inv[i].quantity > 0 {
							c.inv[i].quantity--
							fmt.Println("✅")
							c.Money++
						}
						if c.inv[i].nom == "Épée C" && c.inv[i].quantity == 0 {
							c.removeInventory(c.inv[i])
							fmt.Println("✅")
						}
					}

				case '5', '(': // Acheter Armor C pour 10 Gold
					for i := 0; i < len(c.inv); i++ {
						if c.inv[i].nom == "Armor C" && c.inv[i].quantity > 0 {
							c.inv[i].quantity--
							fmt.Println("✅")
							c.Money++
						}
						if c.inv[i].nom == "Armor C" && c.inv[i].quantity == 0 {
							c.removeInventory(c.inv[i])
							fmt.Println("✅")
						}
					}

				case '6', '-': // Acheter Sword B pour 20 Gold
					for i := 0; i < len(c.inv); i++ {
						if c.inv[i].nom == "Épée B" && c.inv[i].quantity > 0 {
							c.inv[i].quantity--
							fmt.Println("✅")
							c.Money++
						}
						if c.inv[i].nom == "Épée B" && c.inv[i].quantity == 0 {
							c.removeInventory(c.inv[i])
							fmt.Println("✅")
						}
					}

				case '7', 'è': // Acheter Armor B pour 20 Gold
					for i := 0; i < len(c.inv); i++ {
						if c.inv[i].nom == "Armor B" && c.inv[i].quantity > 0 {
							c.inv[i].quantity--
							fmt.Println("✅")
							c.Money++
						}
						if c.inv[i].nom == "Armor B" && c.inv[i].quantity == 0 {
							c.removeInventory(c.inv[i])
							fmt.Println("✅")
						}
					}

					// Ressources
				case 'a', 'A':
					// Vendre Rock pour 1 Gold
					for i := 0; i < len(c.inv); i++ {
						if c.inv[i].nom == rock.nom && c.inv[i].quantity > 0 {
							c.inv[i].quantity--
							fmt.Println("✅")
							c.Money++
						}
						if c.inv[i].nom == rock.nom && c.inv[i].quantity == 0 {
							c.removeInventory(c.inv[i])
							fmt.Println("✅")
						}
					}

				case 'b', 'B':
					// Vendre Wood pour 1 Gold
					for i := 0; i < len(c.inv); i++ {
						if c.inv[i].nom == wood.nom && c.inv[i].quantity > 0 {
							c.inv[i].quantity--
							fmt.Println("✅")
							c.Money++
						}
						if c.inv[i].nom == wood.nom && c.inv[i].quantity == 0 {
							c.removeInventory(c.inv[i])
							fmt.Println("✅")
						}
					}

				case 'c', 'C':
					// Vendre Scrap pour 1 Gold
					for i := 0; i < len(c.inv); i++ {
						if c.inv[i].nom == scrap.nom && c.inv[i].quantity > 0 {
							c.inv[i].quantity--
							fmt.Println("✅")
							c.Money += 5
						}
						if c.inv[i].nom == scrap.nom && c.inv[i].quantity == 0 {
							c.removeInventory(c.inv[i])
							fmt.Println("✅")
						}
					}

				case 'd', 'D':
					// Vendre Fourrure de Loup pour 1 Gold
					for i := 0; i < len(c.inv); i++ {
						if c.inv[i].nom == FOL.nom && c.inv[i].quantity > 0 {
							c.inv[i].quantity--
							fmt.Println("✅")
							c.Money += 4
						}
						if c.inv[i].nom == FOL.nom && c.inv[i].quantity == 0 {
							c.removeInventory(c.inv[i])
							fmt.Println("✅")
						}
					}

				case 'e', 'E':
					// Vendre Peau de Troll pour 1 Gold
					for i := 0; i < len(c.inv); i++ {
						if c.inv[i].nom == PDT.nom && c.inv[i].quantity > 0 {
							c.inv[i].quantity--
							fmt.Println("✅")
							c.Money += 7
						}
						if c.inv[i].nom == PDT.nom && c.inv[i].quantity == 0 {
							c.removeInventory(c.inv[i])
							fmt.Println("✅")
						}
					}

				case 'f', 'F':
					// Vendre Cuir de Sanglier pour 1 Gold
					for i := 0; i < len(c.inv); i++ {
						if c.inv[i].nom == CDS.nom && c.inv[i].quantity > 0 {
							c.inv[i].quantity--
							fmt.Println("✅")
							c.Money += 3
						}
						if c.inv[i].nom == CDS.nom && c.inv[i].quantity == 0 {
							c.removeInventory(c.inv[i])
							fmt.Println("✅")
						}
					}

				case 'g', 'G':
					// Vendre Plume de Corbeau pour 1 Gold
					for i := 0; i < len(c.inv); i++ {
						if c.inv[i].nom == PDC.nom && c.inv[i].quantity > 0 {
							c.inv[i].quantity--
							fmt.Println("✅")
							c.Money++
						}
						if c.inv[i].nom == PDC.nom && c.inv[i].quantity == 0 {
							c.removeInventory(c.inv[i])
							fmt.Println("✅")
						}
					}

				case 'q', 'Q':
					fmt.Println("Au revoir !")
					break VenteLoop
				}
			}
		case 'q', 'Q':
			break MarchandLoop

		}
	}
}
