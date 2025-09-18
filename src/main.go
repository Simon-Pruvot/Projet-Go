package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/eiannone/keyboard"
	"golang.org/x/term"
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
	fmt.Println("Press I to open inventory, H to drink potion, P to pause, D to display info, Q to quit.")

	baseHpMax := player.HpMax // save class's true base HP
	player.Hp = baseHpMax

	for {
		char, _, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		switch char {
		case 'i', 'I':
			player.accessInventory(baseHpMax)
		case 'h', 'H':
			player.takePot()
		case 'q', 'Q':
			fmt.Println("Goodbye!")
			return
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
		case 'd', 'D':
			player.displayInfo()
		case 'b', 'B':
			player.Marchand()
		case 'f', 'F':
			player.Forgeron(swordLegend, armorLegend, potionLegend)
		case 't', 'T':
			goblin := initGoblin()
			trainingFight(&player, goblin)
		}
	}
}

func Espace(esp int, chaine1 string, chaine2 string) string {
	esp -= (len(chaine1) + len(chaine2))
	for i := 0; i <= esp; i++ {
		chaine1 += " "
	}
	return chaine1 + chaine2 + string(rune(127))
}

func (c *Character) UseArmorSet(setName string) {
	switch setName {
	case "Armor A":
		c.inv = append(c.inv, objects["Chapeau A"], objects["Tunique A"], objects["Bottes A"])
		fmt.Println("🎁 Vous avez reçu Chapeau A, Tunique A et Bottes A !")

	case "Armor B":
		c.inv = append(c.inv, objects["Chapeau B"], objects["Tunique B"], objects["Bottes B"])
		fmt.Println("🎁 Vous avez reçu Chapeau B, Tunique B et Bottes B !")

	case "Armor C":
		c.inv = append(c.inv, objects["Chapeau C"], objects["Tunique C"], objects["Bottes C"])
		fmt.Println("🎁 Vous avez reçu Chapeau C, Tunique C et Bottes C !")
	}
}

func (c Character) displayInfo() {
	lines := []string{
		"╔══════════════════════════════╗",
		"║                              ║",
		fmt.Sprintf("║ Nom    : %-18s  ║", c.Nom),
		fmt.Sprintf("║ Classe : %-18s  ║", c.Classe),
		fmt.Sprintf("║ Lvl    : %-18d  ║", c.Lvl),
		fmt.Sprintf("║ Hp Max : %-18d  ║", c.HpMax),
		fmt.Sprintf("║ Hp     : %-18d  ║", c.Hp),
		"║                              ║",
		"╚══════════════════════════════╝",
	}

	FullScreenDrawCentered(lines)
}

func (c *Character) accessInventory(baseHpMax int) {
	slots := 10
	perRow := 5

BoucleInventaire:
	for {
		// --- Affichage de l'inventaire ---
		var cols [][]string
		for row := 0; row < perRow; row++ {
			cols = append(cols, []string{})
		}

		for i := 0; i < slots; i++ {
			item := "[ vide ]"
			if i < len(c.inv) {
				item = fmt.Sprintf("[%d] %s", i, c.inv[i].nom)
			}
			cols[i%perRow] = append(cols[i%perRow], item)
		}

		lines := []string{"🎒 Inventaire :", strings.Repeat("-", 80)}
		grid := CombineColumnsToLines(cols, 4) // 4 espaces entre colonnes
		lines = append(lines, grid...)
		lines = append(lines, strings.Repeat("-", 80))
		lines = append(lines,
			"Options : (U)tiliser | (E)quiper | (R)etirer | (Q)uitter",
		)

		FullScreenDrawCentered(lines)

		// --- Gestion des entrées ---
		char, _, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		switch char {
		case 'q', 'Q':
			fmt.Println("✅ Inventoire Quite :")
			break BoucleInventaire

		case 'u', 'U':
			index := demanderIndex(len(c.inv))
			if index >= 0 {
				item := c.inv[index]
				if strings.HasPrefix(item.nom, "Armor ") {
					c.UseArmorSet(item.nom)
					// Remove the used set
					c.inv = append(c.inv[:index], c.inv[index+1:]...)
				} else {
					fmt.Println("⚠️ Cet objet ne peut pas être utilisé :", item.nom)
				}
			}

		case 'e', 'E':
			index := demanderIndex(len(c.inv))
			if index >= 0 {
				// TODO: demander le slot exact (chapeau/tunique/bottes)
				c.Equip(c.inv[index], "chapeau", baseHpMax)
				fmt.Println("✅ Équipé :", c.inv[index].nom)
			}

		case 'r', 'R':
			index := demanderIndex(len(c.inv))
			if index >= 0 {
				fmt.Println("🗑️ Retiré :", c.inv[index].nom)
				c.inv = append(c.inv[:index], c.inv[index+1:]...)
			}
		}
	}
}

// Fonction pour demander un index
func demanderIndex(max int) int {
	fmt.Printf("👉 Choisissez l’index de l’objet (0-%d) : ", max-1)
	var index int
	_, err := fmt.Scanf("%d\n", &index)
	if err != nil || index < 0 || index >= max {
		fmt.Println("❌ Index invalide")
		return -1
	}
	return index
}

func (c *Character) takePot() {
	for i := 0; i < len(c.inv); i++ {
		if c.inv[i].nom == "Potion de vie" && c.inv[i].quantity > 0 {
			c.inv[i].quantity--
			c.Hp += 50
			if c.Hp > c.HpMax {
				c.Hp = c.HpMax
			}
			fmt.Println("🍷 You drank a potion! HP:", c.Hp, "/", c.HpMax)
			fmt.Println(``)

			// if no more left → remove it
			if c.inv[i].quantity == 0 {
				c.inv = append(c.inv[:i], c.inv[i+1:]...)
			}
			return
		}
	}
	fmt.Println("⚠️ No potion left!")
}

// Fonction pour ajouter un objet à l'inventaire
func (c *Character) addInventory(obj Objects) {
	if c.canAddItem() {
		c.inv = append(c.inv, obj)
	}
}

func (c *Character) removeInventory(obj Objects) {
	for i := 0; i < len(c.inv); i++ {
		if c.inv[i].nom == obj.nom {
			c.inv[i].quantity--

			// Si plus d’objets, on supprime l'entrée du slice
			if c.inv[i].quantity <= 0 {
				c.inv = append(c.inv[:i], c.inv[i+1:]...)
			}
			return // On sort après avoir trouvé l'objet
		}
	}
}

func (c *Character) canAddItem() bool {
	if len(c.inv) >= c.MaxInventorySlots {
		fmt.Println("⚠️ Inventaire plein ! Impossible d’ajouter plus d’objets.")
		return false
	}
	return true
}

func (c *Character) upgradeInventorySlot() {
	if c.InventoryUpgrades >= 3 {
		fmt.Println("⚠️ Maximum d'augmentations atteint.")
		return
	}
	c.MaxInventorySlots += 10
	c.InventoryUpgrades++
	fmt.Printf("Capacité d'inventaire augmentée : %d (restant %d)\n", c.MaxInventorySlots, 3-c.InventoryUpgrades)
}
func initCharacter(nom, classe string, lvl, hpmax, hp int, inv []Objects, skills []string, equipment Equipment) Character {
	return Character{Nom: nom, Classe: classe, Lvl: lvl, HpMax: hpmax, Hp: hp, inv: inv, Money: 100, Skills: skills, Equipment: equipment, MaxInventorySlots: 10, InventoryUpgrades: 0, Mana: 30, ManaMax: 30, bonusDMG: 0}
}

// Equip allows to equip an object in the right slot
func (c *Character) Equip(item Objects, slot string, baseHpMax int) {
	// Remove from inventory
	for i := 0; i < len(c.inv); i++ {
		if c.inv[i].nom == item.nom {
			c.inv = append(c.inv[:i], c.inv[i+1:]...)
			break
		}
	}

	// Handle replacement
	var replaced *Objects
	switch slot {
	case "chapeau":
		replaced = c.Equipment.Chapeau
		c.Equipment.Chapeau = &item
	case "tunique":
		replaced = c.Equipment.Tunique
		c.Equipment.Tunique = &item
	case "bottes":
		replaced = c.Equipment.Bottes
		c.Equipment.Bottes = &item
	case "Épée C":
		c.bonusDMG += 5
	case "Épée B":
		c.bonusDMG += 10
	case "Épée A":
		c.bonusDMG += 15
	default:
		fmt.Println("⚠️ Slot inconnu :", slot)
		return
	}

	// Return old equipment to inventory if replaced
	if replaced != nil {
		c.inv = append(c.inv, *replaced)
	}

	// Recalculate max HP
	c.updateHpMax(baseHpMax)
}

// updateHpMax recalculates HpMax from baseHpMax + equipment bonuses
func (c *Character) updateHpMax(baseHpMax int) {
	c.HpMax = baseHpMax

	if c.Equipment.Chapeau != nil && c.Equipment.Chapeau.nom == "Chapeau de l’aventurier" {
		c.HpMax += 10
	}
	if c.Equipment.Tunique != nil && c.Equipment.Tunique.nom == "Tunique de l’aventurier" {
		c.HpMax += 25
	}
	if c.Equipment.Bottes != nil && c.Equipment.Bottes.nom == "Bottes de l’aventurier" {
		c.HpMax += 15
	}

	// Adjust current HP if it exceeds new max
	if c.Hp > c.HpMax {
		c.Hp = c.HpMax
	}
}

// DisplayEquipment shows current equipment
func (c Character) DisplayEquipment() {
	fmt.Println("🛡️ Équipement actuel :")
	if c.Equipment.Chapeau != nil {
		fmt.Println("  Chapeau :", c.Equipment.Chapeau.nom)
	} else {
		fmt.Println("  Chapeau : Aucun")
	}
	if c.Equipment.Tunique != nil {
		fmt.Println("  Tunique :", c.Equipment.Tunique.nom)
	} else {
		fmt.Println("  Tunique : Aucune")
	}
	if c.Equipment.Bottes != nil {
		fmt.Println("  Bottes :", c.Equipment.Bottes.nom)
	} else {
		fmt.Println("  Bottes : Aucune")
	}
}

func (c *Character) Marchand() {
	building := []string{
		"=========================================",
		"=========================================",
		"     _______                             ",
		"     ||     |   |      |   \\     /      ",
		"     ||    _|   |      |    \\   /       ",
		"     ||---|_    |      |     \\ /        ",
		"     ||     |   |      |       |         ",
		"     ||-----|   |______|       |         ",
		"                                         ",
		"                                         ",
		"                                         ",
		"        __________________________       ",
		"       /                         \\      ",
		"      /                           \\     ",
		"     /_____________________________\\    ",
		"    |   /-_      /-_     /-_      /|     ",
		"    |_-/   \\-_-/   \\-_-/   \\-_-/|     ",
		"    |          _______             |     ",
		"    |         /       \\           |     ",
		"    |         | .   . |            |     ",
		"    |         |  -_-  |            |     ",
		"    |         \\-------/           |     ",
		"    |              |               |     ",
		"    |              |               |     ",
		"  ------------------------------------   ",
		"   | |============================| |    ",
		"   | |             |              | |    ",
		"   | |            / \\            | |    ",
		"   | |           /   \\           | |    ",
		"=========================================",
		"=========================================",
	}

	//objects
	hpPot := Objects{nom: "Potion de vie", quantity: 1}
	poisonPot := Objects{nom: "Potion de poison", quantity: 1}
	swordCom := Objects{nom: "Épée C", quantity: 1}
	swordRare := Objects{nom: "Épée B", quantity: 1}

	armorCom := Objects{nom: "Armor C", quantity: 1}
	armorRare := Objects{nom: "Armor B", quantity: 1}

	LivFeu := Objects{nom: "Livre de Sort : Boule de feu", quantity: 1}

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
								added = true
								break
							}
						}
						if !added {
							c.addInventory(hpPot)
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
								added = true
								break
							}
						}
						if !added {
							c.addInventory(poisonPot)
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
								added = true
								break
							}
						}
						if !added {
							c.addInventory(swordCom)
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
								added = true
								break
							}
						}
						if !added {
							c.addInventory(armorCom)
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
								added = true
								break
							}
						}
						if !added {
							c.addInventory(swordRare)
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
								added = true
								break
							}
						}
						if !added {
							c.addInventory(armorRare)
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
								added = true
								break
							}
						}
						if !added {
							c.addInventory(LivFeu)
						}
					} else {
						fmt.Println("Need More Gold")
					}
				case '9': // Augmentation d’inventaire
					if c.Money >= 30 {
						c.Money -= 30
						c.upgradeInventorySlot()
					} else {
						fmt.Println("💰 Pas assez d’or pour l’augmentation !")
					}
				case 'q', 'Q':
					fmt.Println("Au revoir !")
					break AchatLoop
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
							c.Money++
						}
						if c.inv[i].nom == "Potion de vie" && c.inv[i].quantity == 0 {
							c.removeInventory(c.inv[i])
						}
					}
				case '2', 'é': // Acheter une potion de poison pour 3 Gold
					for i := 0; i < len(c.inv); i++ {
						if c.inv[i].nom == "Potion de poison" && c.inv[i].quantity > 0 {
							c.inv[i].quantity--
							c.Money++
						}
						if c.inv[i].nom == "Potion de poison" && c.inv[i].quantity == 0 {
							c.removeInventory(c.inv[i])
						}
					}

				case '3', '"': // Acheter Sword C pour 10 Gold
					for i := 0; i < len(c.inv); i++ {
						if c.inv[i].nom == "Épée C" && c.inv[i].quantity > 0 {
							c.inv[i].quantity--
							c.Money++
						}
						if c.inv[i].nom == "Épée C" && c.inv[i].quantity == 0 {
							c.removeInventory(c.inv[i])
						}
					}

				case '5', '(': // Acheter Armor C pour 10 Gold
					for i := 0; i < len(c.inv); i++ {
						if c.inv[i].nom == "Armor C" && c.inv[i].quantity > 0 {
							c.inv[i].quantity--
							c.Money++
						}
						if c.inv[i].nom == "Armor C" && c.inv[i].quantity == 0 {
							c.removeInventory(c.inv[i])
						}
					}

				case '6', '-': // Acheter Sword B pour 20 Gold
					for i := 0; i < len(c.inv); i++ {
						if c.inv[i].nom == "Épée B" && c.inv[i].quantity > 0 {
							c.inv[i].quantity--
							c.Money++
						}
						if c.inv[i].nom == "Épée B" && c.inv[i].quantity == 0 {
							c.removeInventory(c.inv[i])
						}
					}

				case '7', 'è': // Acheter Armor B pour 20 Gold
					for i := 0; i < len(c.inv); i++ {
						if c.inv[i].nom == "Armor B" && c.inv[i].quantity > 0 {
							c.inv[i].quantity--
							c.Money++
						}
						if c.inv[i].nom == "Armor B" && c.inv[i].quantity == 0 {
							c.removeInventory(c.inv[i])
						}
					}

					// Ressources
				case 'a', 'A':
					// Vendre Rock pour 1 Gold
					for i := 0; i < len(c.inv); i++ {
						if c.inv[i].nom == rock.nom && c.inv[i].quantity > 0 {
							c.inv[i].quantity--
							c.Money++
						}
						if c.inv[i].nom == rock.nom && c.inv[i].quantity == 0 {
							c.removeInventory(c.inv[i])
						}
					}

				case 'b', 'B':
					// Vendre Wood pour 1 Gold
					for i := 0; i < len(c.inv); i++ {
						if c.inv[i].nom == wood.nom && c.inv[i].quantity > 0 {
							c.inv[i].quantity--
							c.Money++
						}
						if c.inv[i].nom == wood.nom && c.inv[i].quantity == 0 {
							c.removeInventory(c.inv[i])
						}
					}

				case 'c', 'C':
					// Vendre Scrap pour 1 Gold
					for i := 0; i < len(c.inv); i++ {
						if c.inv[i].nom == scrap.nom && c.inv[i].quantity > 0 {
							c.inv[i].quantity--
							c.Money += 5
						}
						if c.inv[i].nom == scrap.nom && c.inv[i].quantity == 0 {
							c.removeInventory(c.inv[i])
						}
					}

				case 'd', 'D':
					// Vendre Fourrure de Loup pour 1 Gold
					for i := 0; i < len(c.inv); i++ {
						if c.inv[i].nom == FOL.nom && c.inv[i].quantity > 0 {
							c.inv[i].quantity--
							c.Money += 4
						}
						if c.inv[i].nom == FOL.nom && c.inv[i].quantity == 0 {
							c.removeInventory(c.inv[i])
						}
					}

				case 'e', 'E':
					// Vendre Peau de Troll pour 1 Gold
					for i := 0; i < len(c.inv); i++ {
						if c.inv[i].nom == PDT.nom && c.inv[i].quantity > 0 {
							c.inv[i].quantity--
							c.Money += 7
						}
						if c.inv[i].nom == PDT.nom && c.inv[i].quantity == 0 {
							c.removeInventory(c.inv[i])
						}
					}

				case 'f', 'F':
					// Vendre Cuir de Sanglier pour 1 Gold
					for i := 0; i < len(c.inv); i++ {
						if c.inv[i].nom == CDS.nom && c.inv[i].quantity > 0 {
							c.inv[i].quantity--
							c.Money += 3
						}
						if c.inv[i].nom == CDS.nom && c.inv[i].quantity == 0 {
							c.removeInventory(c.inv[i])
						}
					}

				case 'g', 'G':
					// Vendre Plume de Corbeau pour 1 Gold
					for i := 0; i < len(c.inv); i++ {
						if c.inv[i].nom == PDC.nom && c.inv[i].quantity > 0 {
							c.inv[i].quantity--
							c.Money++
						}
						if c.inv[i].nom == PDC.nom && c.inv[i].quantity == 0 {
							c.removeInventory(c.inv[i])
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

func (c *Character) isDead() {
	if c.Hp <= 0 {
		fmt.Println("U DEAD")
		c.Hp = c.HpMax / 2
	}
}

func (c *Character) spellBook(spell string) {
	for _, s := range c.Skills {
		if s == spell {
			fmt.Println("⚠️ Vous connaissez déjà le sort :", spell)
			return
		}
	}
	c.Skills = append(c.Skills, spell)
	fmt.Println("✨ Nouveau sort appris :", spell)
}

// --- Utilisation des objets dans l'inventaire ---
func (c *Character) useItem(objName string) {
	for i := 0; i < len(c.inv); i++ {
		if c.inv[i].nom == objName && c.inv[i].quantity > 0 {
			if objName == "Potion de vie" {
				c.takePot()
				return
			}
			if objName == "Livre de Sort : Boule de feu" {
				c.spellBook("Boule de feu")
				c.inv[i].quantity--
				if c.inv[i].quantity == 0 {
					c.inv = append(c.inv[:i], c.inv[i+1:]...)
				}
				return
			}
		}
	}
	fmt.Println("⚠️ Objet introuvable :", objName)
}

func (c Character) displaySkills() {
	fmt.Println("📖 Liste de vos sorts :")
	for _, s := range c.Skills {
		fmt.Println(" -", s)
	}
	fmt.Println()
}

// PrintColumns prints columns side-by-side.
// cols: slice of columns (each column is a []string).
// distance: number of spaces between columns (last parameter).
func PrintColumns(cols [][]string, distance int) {
	if distance < 0 {
		distance = 0
	}
	// compute width for each column (rune-aware)
	widths := make([]int, len(cols))
	for ci, col := range cols {
		for _, line := range col {
			if l := utf8.RuneCountInString(line); l > widths[ci] {
				widths[ci] = l
			}
		}
	}

	// compute max number of lines
	maxLines := 0
	for _, col := range cols {
		if len(col) > maxLines {
			maxLines = len(col)
		}
	}

	sep := strings.Repeat(" ", distance)

	// print row by row
	for r := 0; r < maxLines; r++ {
		for ci, col := range cols {
			cell := ""
			if r < len(col) {
				cell = col[r]
			}
			// left-align to column width
			fmt.Printf("%-*s", widths[ci], cell)
			// print separator except after last column
			if ci < len(cols)-1 {
				fmt.Print(sep)
			}
		}
		fmt.Println()
	}
}

// truncateRunes returns the first n runes of s (safe with utf8), adds "..." if truncated
func truncateRunes(s string, n int) string {
	if n <= 0 {
		return ""
	}
	if utf8.RuneCountInString(s) <= n {
		return s
	}
	r := []rune(s)
	if n <= 3 {
		return string(r[:n])
	}
	return string(r[:n-3]) + "..."
}

// CombineColumnsToLines returns lines composed of columns side-by-side.
// cols: slice of columns (each column is []string). distance: spaces between columns.
func CombineColumnsToLines(cols [][]string, distance int) []string {
	// compute width for each column (rune-aware)
	widths := make([]int, len(cols))
	for ci, col := range cols {
		for _, line := range col {
			if l := utf8.RuneCountInString(line); l > widths[ci] {
				widths[ci] = l
			}
		}
	}
	// max lines
	maxLines := 0
	for _, col := range cols {
		if len(col) > maxLines {
			maxLines = len(col)
		}
	}
	sep := strings.Repeat(" ", distance)
	out := make([]string, 0, maxLines)
	for r := 0; r < maxLines; r++ {
		var parts []string
		for ci, col := range cols {
			cell := ""
			if r < len(col) {
				cell = col[r]
			}
			// left-align to column width
			parts = append(parts, fmt.Sprintf("%-*s", widths[ci], cell))
		}
		out = append(out, strings.Join(parts, sep))
	}
	return out
}

// FullScreenDrawCentered clears terminal and prints lines centered horizontally & vertically.
// It pads the screen to terminal height so nothing else is visible.
func FullScreenDrawCentered(lines []string) {
	// hide cursor
	fmt.Print("\033[?25l")
	// clear screen + move cursor to top-left
	fmt.Print("\033[2J\033[H")

	// get terminal size
	w, h, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		// fallback: print lines normally
		for _, ln := range lines {
			fmt.Println(ln)
		}
		fmt.Print("\033[?25h")
		return
	}

	// ensure lines fit; truncate any line longer than width
	for i, ln := range lines {
		if utf8.RuneCountInString(ln) > w {
			lines[i] = truncateRunes(ln, w)
		}
	}

	// vertical centering: compute top padding
	if len(lines) > h {
		// too many lines → just print the first h lines (can't center)
		lines = lines[:h]
	}
	topPad := (h - len(lines)) / 2

	// print top padding
	for i := 0; i < topPad; i++ {
		fmt.Println()
	}

	// print each line centered horizontally
	for _, ln := range lines {
		rlen := utf8.RuneCountInString(ln)
		leftPad := 0
		if w > rlen {
			leftPad = (w - rlen) / 2
		}
		if leftPad > 0 {
			fmt.Print(strings.Repeat(" ", leftPad))
		}
		fmt.Println(ln)
	}

	// bottom padding to fill screen
	printed := topPad + len(lines)
	for printed < h {
		fmt.Println()
		printed++
	}

	// show cursor again
	fmt.Print("\033[?25h")
}

func (c *Character) Forgeron(swordLegend, armorLegend, potionLegend Objects) {
	// ressources requises
	reqs := map[string]map[string]int{
		"Épée A": {
			"Minerai de fer": 3,
			"Bois":           1,
		},
		"Armor A": {
			"Peau de Troll":    2,
			"Fourrure de Loup": 3,
		},
		"Potion Légendaire": {
			"Herbe magique":   2,
			"Champignon rare": 1,
		},
		"Potion de vie": {
			"Cuir de Sanglier": 1,
		},
	}

	content := []string{
		"",
		"Bienvenue chez le Forgeron ⚒️",
		"Objets que vous pouvez fabriquer (5 Gold chacun) :",
		"  1 - Épée A",
		"  2 - Armor A",
		"  3 - Potion Légendaire",
		"  4 - Potion de vie",
		"",
		"Appuyez sur 1,2,3,4 pour fabriquer, ou Q pour quitter.",
	}

	FullScreenDrawCentered(content)

ForgeronLoop:
	for {
		char, _, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		switch char {
		case '1':
			c.craftItem(swordLegend, reqs[swordLegend.nom])
		case '2':
			c.craftItem(armorLegend, reqs[armorLegend.nom])
		case '3':
			c.craftItem(potionLegend, reqs[potionLegend.nom])
		case 'q', 'Q':
			fmt.Println("👋 Le forgeron vous salue !")
			break ForgeronLoop
		}
	}
}

// Fonction utilitaire pour fabriquer un item
func (c *Character) craftItem(item Objects, req map[string]int) {
	if c.Money < 5 {
		fmt.Println("💰 Vous n’avez pas assez d’or pour fabriquer cet objet !")
		return
	}

	// Vérification des ressources
	for res, qte := range req {
		found := false
		for _, inv := range c.inv {
			if inv.nom == res && inv.quantity >= qte {
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("⚠️ Ressource manquante : %s (besoin de %d)\n", res, qte)
			return
		}
	}

	// Retrait des ressources
	for res, qte := range req {
		for i := 0; i < len(c.inv); i++ {
			if c.inv[i].nom == res {
				c.inv[i].quantity -= qte
				if c.inv[i].quantity <= 0 {
					c.inv = append(c.inv[:i], c.inv[i+1:]...)
				}
				break
			}
		}
	}

	// Retrait de l’or
	c.Money -= 5

	// Ajout de l’objet
	c.addInventory(item)
	CreaftingAnim()
	fmt.Printf("✅ Vous avez fabriqué : %s\n", item.nom)
}

func initGoblin() Monster {
	return Monster{
		Nom:   "Gobelin d'entraînement",
		HpMax: 40,
		Hp:    40, // égal à HpMax
		Atk:   5,
	}
}

// goblinPattern applique l'attaque du gobelin sur le joueur.
func goblinPattern(g *Monster, c *Character, turn int) {
	damage := g.Atk
	if turn%3 == 0 { // tous les 3 tours -> 200%
		damage = g.Atk * 2
	}
	c.Hp -= damage
	if c.Hp < 0 {
		c.Hp = 0
	}
	fmt.Printf("%s inflige à %s %d de dégâts\n", g.Nom, c.Nom, damage)
	fmt.Printf("%s PV : %d / %d\n\n", c.Nom, c.Hp, c.HpMax)
}

func characterTurn(c *Character, m *Monster) bool {
	// Recharge 15 mana par tour (max ManaMax)
	c.Mana += 15
	if c.Mana > c.ManaMax {
		c.Mana = c.ManaMax
	}

	fmt.Println("=== Votre tour ===")
	fmt.Printf("PV: %d / %d | Mana: %d / %d\n", c.Hp, c.HpMax, c.Mana, c.ManaMax)
	fmt.Println("Options : (A)ttaquer | (S)péciale | (M)agie | (I)nventaire | (K) Skip")
	fmt.Print("Choix : ")

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Entrée invalide.")
		return false
	}

	switch strings.ToLower(input) {
	case "a":
		// Attaque basique : 5 dégâts, coût 10 mana
		if c.Mana < 10 {
			fmt.Println("Pas assez de mana pour attaque basique.")
			return false
		}
		c.Mana -= 10
		m.Hp -= 5
		if m.Hp < 0 {
			m.Hp = 0
		}
		fmt.Printf("%s utilise Attaque basique et inflige %d dégâts à %s\n", c.Nom, 5, m.Nom)

	case "s":
		// Spéciale : 12 dégâts, coût 20 mana
		if c.Mana < 20 {
			fmt.Println("Pas assez de mana pour attaque spéciale.")
			return false
		}
		c.Mana -= 20
		m.Hp -= 12
		if m.Hp < 0 {
			m.Hp = 0
		}
		fmt.Printf("%s utilise Attaque spéciale et inflige %d dégâts à %s\n", c.Nom, 12, m.Nom)

	case "m":
		// Magie : uniquement si on a appris au moins un sort
		if len(c.Skills) == 0 {
			fmt.Println("Vous n'avez appris aucun sort.")
			return false
		}
		if c.Mana < 30 {
			fmt.Println("Pas assez de mana pour lancer un sort.")
			return false
		}
		// Pour l'instant : utiliser le premier sort appris
		spell := c.Skills[0]
		c.Mana -= 30
		m.Hp -= 20
		if m.Hp < 0 {
			m.Hp = 0
		}
		fmt.Printf("%s lance %s et inflige %d dégâts à %s\n", c.Nom, spell, 20, m.Nom)

	case "i":
		// Inventaire
		if len(c.inv) == 0 {
			fmt.Println("Inventaire vide.")
			return false
		}
		fmt.Println("Inventaire :")
		for i, it := range c.inv {
			fmt.Printf("[%d] %s x%d\n", i, it.nom, it.quantity)
		}
		fmt.Print("Choisissez index à utiliser (ou -1 pour annuler) : ")
		var idx int
		_, err := fmt.Scanf("%d\n", &idx)
		if err != nil || idx < 0 || idx >= len(c.inv) {
			fmt.Println("Annulation ou index invalide.")
			return false
		}
		item := c.inv[idx]

		// Exemple : livres et potions
		switch item.nom {
		case "Potion de vie":
			c.Hp += 20
			if c.Hp > c.HpMax {
				c.Hp = c.HpMax
			}
		case "Potion de poison":
			c.UsePoison(m)
			fmt.Printf("Vous utilisez %s. PV : %d / %d\n", item.nom, c.Hp, c.HpMax)
		case "Livre de sort":
			// Apprendre un nouveau sort
			c.Skills = append(c.Skills, "Fireball") // tu peux changer le nom du sort selon le livre
			item.quantity--
			fmt.Println("Vous apprenez un nouveau sort : Fireball !")
		default:
			fmt.Printf("Vous utilisez %s (effet non implémenté).\n", item.nom)
			item.quantity--
		}

		if item.quantity <= 0 {
			c.inv = append(c.inv[:idx], c.inv[idx+1:]...)
		} else {
			c.inv[idx].quantity = item.quantity
		}

	case "k":
		// Skip : ne fait rien, regen mana déjà appliquée
		fmt.Printf("%s choisit de concentrer son énergie et passe son tour.\n", c.Nom)

	default:
		fmt.Println("Option inconnue, tour passé.")
	}

	return m.Hp <= 0
}

func trainingFight(player *Character, goblin Monster) {
	turn := 1
	fmt.Printf("=== Début de l'entraînement contre %s ===\n", goblin.Nom)

	// Seed pour la RNG (évite d’avoir toujours le même résultat)
	rand.Seed(time.Now().UnixNano())

	// Définition du loot possible
	lootPool := []Objects{
		{"Fourrure de Loup", 1},
		{"Peau de Troll", 1},
		{"Cuir de Sanglier", 1},
		{"Plume de Corbeau", 1},
	}

	for player.Hp > 0 && goblin.Hp > 0 {
		fmt.Printf("\n---- Tour %d ----\n", turn)
		characterTurn(player, &goblin) // joueur joue

		if goblin.Hp <= 0 {
			fmt.Printf("%s est vaincu !\n", goblin.Nom)

			// Récompenses fixes
			GainXP(player, 50)
			player.Money += 20

			// Loot aléatoire
			randomLoot := lootPool[rand.Intn(len(lootPool))]
			player.inv = append(player.inv, Objects{"Cuir de Sanglier", 1}) // loot 100%
			player.inv = append(player.inv, randomLoot)

			fmt.Printf("🎁 Récompenses : 50 XP, 20 or, +1 %s\n", randomLoot.nom)
			break
		}

		goblinPattern(&goblin, player, turn) // gobelin joue
		turn++
	}

	if player.Hp <= 0 {
		fmt.Println("Vous avez été vaincu... retour au menu.")
	}
}

// Add XP gain
func GainXP(c *Character, amount int) {
	c.XP += amount
	fmt.Printf("%s gagne %d XP (total: %d/%d)\n", c.Nom, amount, c.XP, XPRequired(c.Lvl))
	if c.XP >= XPRequired(c.Lvl) {
		LevelUp(c)
	}
}

// How much XP needed for next level
func XPRequired(level int) int {
	return 100 * level // ex: lvl1 -> 100 XP, lvl2 -> 200, etc.
}

// Handle level up
func LevelUp(c *Character) {
	c.XP -= XPRequired(c.Lvl)
	c.Lvl++
	c.HpMax += 20
	c.ManaMax += 5
	c.Hp = c.HpMax
	c.Mana = c.ManaMax
	fmt.Printf("🎉 %s passe au niveau %d ! HP: %d | Mana: %d\n", c.Nom, c.Lvl, c.HpMax, c.ManaMax)
}
