package main

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

type Objects struct {
	nom      string
	quantity int
}

type Character struct {
	Nom    string
	Classe string
	Lvl    int
	HpMax  int
	Hp     int
	inv    []Objects
}

func main() {
	hpPot := Objects{nom: "HP potion", quantity: 1}
	poisonPot := Objects{nom: "Poison potion", quantity: 1}
	swordCom := Objects{nom: "Sword C", quantity: 1}
	swordRare := Objects{nom: "Sword B", quantity: 1}
	swordLegend := Objects{nom: "Sword A", quantity: 1}
	armorCom := Objects{nom: "Armor C", quantity: 1}
	armorRare := Objects{nom: "Armor B", quantity: 1}
	armorLegend := Objects{nom: "Armor A", quantity: 1}
	oneyBag := Objects{nom: "Gold", quantity: 1}

	Simon := initCharacter("Simon", "Elfe", 1, 100, 40, []Objects{{"HP potion", 3}})

	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	fmt.Println("Press I to open inventory, H to drink potion, P to pause, D to display info, Q to quit.")

	for {
		char, _, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		switch char {
		case 'i', 'I':
			Simon.accessInventory2()
		case 'h', 'H':
			Simon.takePot()
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
					Simon.accessInventory2()
					break MenuLoop
				case 'd', 'D':
					Simon.displayInfo()
					break MenuLoop
				case 'q', 'Q':
					fmt.Println("Closing menu...")
					return
				default:
					fmt.Println("Invalid option. Please press I, D, or Q.")
				}

			}
		case 'd', 'D':
			Simon.displayInfo()
		case 'b', 'B':
			PrintMarchand()

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

func initCharacter(nom string, classe string, lvl int, hpmax int, hp int, inv []Objects) Character {
	return Character{Nom: nom, Classe: classe, Lvl: lvl, HpMax: hpmax, Hp: hp, inv: inv}
}

func (c Character) displayInfo() {
	fmt.Println("\n")
	fmt.Println("------------------------------")
	fmt.Println("------------------------------")
	fmt.Println("\tNom : " + c.Nom)
	fmt.Println("\tClasse : " + c.Classe)
	fmt.Println("\tLvl :", c.Lvl)
	fmt.Println("\tHp Max :", c.HpMax)
	fmt.Println("\tHp :", c.Hp)
	fmt.Println("------------------------------")
	fmt.Println("------------------------------")
	fmt.Println("\n")
}

func (c Character) accessInventory() {
	fmt.Println("\n")
	fmt.Println("            		Inventory:")
	fmt.Println("-------------------------------------------------------------")

	slots := 16 // 4x4 cases fixes
	perRow := 4

	for i := 0; i < slots; i++ {
		if i < len(c.inv) {
			obj := c.inv[i]
			// Affiche nom + quantité de l'objet
			text := Espace(10, obj.nom, fmt.Sprintf("x%d", obj.quantity))
			fmt.Printf("[ %s ]", text)
		} else {
			// Case vide
			text := Espace(10, "", "")
			fmt.Printf("[ %s ]", text)
		}
		// retour ligne après 4 cases
		if (i+1)%perRow == 0 {
			fmt.Println()
			fmt.Println("-------------------------------------------------------------")
		}
	}
	fmt.Println("\n")
}

func (c Character) accessInventory2() {
	fmt.Println("\n")
	fmt.Println("            		Inventory:")
	fmt.Println("████████████████████████████████████████████████████████████")

	slots := 16 // 4x4 cases fixes
	perRow := 4

	for i := 0; i < slots; i++ {
		if i < len(c.inv) {
			obj := c.inv[i]
			// Affiche nom + quantité de l'objet
			text := Espace(10, obj.nom, fmt.Sprintf("x%d", obj.quantity))
			fmt.Printf("█ %s █", text)
		} else {
			// Case vide
			text := Espace(10, "", "")
			fmt.Printf("█ %s █", text)
		}
		// retour ligne après 4 cases
		if (i+1)%perRow == 0 {
			fmt.Println()
			fmt.Println("████████████████████████████████████████████████████████████")
		}
	}
	fmt.Println("\n")
}

func (c *Character) takePot() {
	for i := 0; i < len(c.inv); i++ {
		if c.inv[i].nom == "HP potion" && c.inv[i].quantity > 0 {
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
	c.inv = append(c.inv, obj)
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

func (c *Character) Marchand() {
	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	fmt.Println("Chose an item: 1, 2, 3, 5, 6, 7")

	for {
		char, _, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		switch char {
		case '1', '&':
			c.accessInventory2()
		case '2', 'é':
			c.takePot()
		case '3', '"':
			fmt.Println("Goodbye!")
			return
		case '5', '(':
			PrintMenu()
		case '5', '(':
			PrintMenu()
		case '5', '(':
			PrintMenu()
		}
	}
}
