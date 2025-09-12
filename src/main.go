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
	Simon := initCharacter("Simon", "Elfe", 1, 100, 40, []Objects{{"HP potion", 3}})

	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	fmt.Println("Press I to open inventory, P to drink potion, Q to quit.")

	for {
		char, _, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		switch char {
		case 'i', 'I':
			Simon.accessInventory2() // or accessInventory()
		case 'p', 'P':
			Simon.takePot()
		case 'q', 'Q':
			fmt.Println("Goodbye!")
			return
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
			fmt.Println(`
        ▄▄▄▀▀▀▀▀▀▄▄
        █▌        █▌
        █▀        ▀█▌
        █   ███ ██▌ █▌
        ▀▄        ▄█
         ▀▌▄     ▄▀
            ▀▄▓▓▄▀▀
            █▐█ ▓
          ▄▀▀▐█  ▓▌
          █  ▐█    ▀▓
        ▄▀   ▐█      ▀▀▄
      ▄▄▀    ▐▌         ▀
             ▐▌
             ▐▌
             ▓▌
            █▓▌
          █  █▌
        ▓▀   ▓▌
      ▓▌     █▌
    ▄█       █▌
    █        █▌
    █        ▓▌
    █        ▓▌
    ▀▀▀▀▀▀▀  ▀▀▀▀▀▀
`)

			// if no more left → remove it
			if c.inv[i].quantity == 0 {
				c.inv = append(c.inv[:i], c.inv[i+1:]...)
			}
			return
		}
	}
	fmt.Println("⚠️ No potion left!")
}
