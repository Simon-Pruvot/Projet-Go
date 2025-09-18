package main

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

func main() {
	// Objets globaux utilisables dans Marchand + Forgeron
	swordLegend := Objects{nom: "Épée A", quantity: 1}
	armorLegend := Objects{nom: "Armor A", quantity: 1}
	potionLegend := Objects{nom: "Potion Rage", quantity: 1}

	//J'appelle la fonction de l'écran de départ
	chosendif := TextBienvenu()

	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	player := CreateClasse()

	if chosendif == "/start" {
	} else if chosendif == "/hard" {
		player.HpMax = player.HpMax / 2
		player.Hp = player.Hp / 2
	} else if chosendif == "/easy" {
	}

	baseHpMax := player.HpMax // save class's true base HP
	player.Hp = baseHpMax

	for {
		char, _, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}
		switch char {
		case 'p', 'P':
			PrintMenu()
		MenuLoop:
			for {
				menuKey, _, err := keyboard.GetKey()
				if err != nil {
					log.Fatal(err)
				}

				switch menuKey {
				case 'i', 'I':
					player.accessInventory(baseHpMax)
					break MenuLoop
				case 'd', 'D':
					player.displayInfo()
					break MenuLoop
				case 'q', 'Q':
					fmt.Println("Closing menu...")
					return
				default:
					fmt.Println("Invalid option. Please press I, D, or Q.")
				}

			}
		case 'b', 'B':
			player.Marchand()
			ShowMap()
		case 'f', 'F':
			player.Forgeron(swordLegend, armorLegend, potionLegend)
			ShowMap()
		case 't', 'T':
			monst := initGoblin("Gobelin")
			trainingFight(&player, monst)
		case 'e', 'E':
			player.DisplayEquipment()
			ShowMap()
		case 'n', 'N':
			monst := initGoblin("dead")
			Fight1(&player, monst)
		case 'm', 'M':
			monst := initGoblin("skeleton")
			Fight2(&player, monst)
		case 'l', 'L':
			monst := initGoblin("dragon")
			Fight3(&player, monst)
		case 's', 'S':
			player.displaySkills()
		default:
			ShowMap()
		}
	}
}
