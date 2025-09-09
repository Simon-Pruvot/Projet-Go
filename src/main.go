package main

import "fmt"

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
	Simon := initCharacter("Simon", "Elfe", 1, 100, 40, []Objects{{"potion", 3}})

	Simon.displayInfo()

	Simon.accessInventory()
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
