package main

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

func main() {
	//J'appelle la fonction de l'écran de départ
	chosendif := TextBienvenu()
	lorefirst()
	player := CreateClasse()

	//Histoire ou non
	var reponse string
	fmt.Print("Voulez-vous voir l'histoire ? (oui/non) : ")
	fmt.Scanln(&reponse)

	history := false
	if reponse == "oui" || reponse == "OUI" {
		history = true
		if player.Classe == "Elfe" {
			loreelfe()
		} else if player.Classe == "Humain" {
			lorehumain()
		} else if player.Classe == "Nain" {
			lorenain()
		}
	} else {
	}

	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

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
			monst := player.initGoblin("Gobelin")
			trainingFight(&player, monst)
		case 'e', 'E':
			player.DisplayEquipment()
			ShowMap()
		case 'n', 'N':
			if history {
				loresquelette()
			}
			monst := player.initGoblin("skeleton")
			Fight1(&player, monst)
		case 'm', 'M':
			if history {
				loreelfe2()
			}
			monst := player.initGoblin("dead")
			Fight2(&player, monst)
		case 'l', 'L':
			if history {
				loreboss3()
			}
			monst := player.initGoblin("dragon")
			Fight3(&player, monst)
		case 's', 'S':
			player.displaySkills()
		default:
			ShowMap()
		}
	}
}
