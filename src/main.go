package main

import (
	"fmt"
	"log"
	"os"

	"strings"
	"unicode/utf8"

	"github.com/eiannone/keyboard"
	"golang.org/x/term"
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
	Money  int
	Skills []string
}

func main() {

	Simon := initCharacter("Simon", "Elfe", 1, 100, 40, []Objects{{"Potion de vie", 3}, {"Potion de poison", 5}}, 100, []string{"Coup de poing"})

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
			Simon.accessInventory()
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
					Simon.accessInventory()
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
			Simon.Marchand()
			fmt.Println("Press I to open inventory, H to drink potion, P to pause, D to display info, Q to quit.")
		case '9', 'ç':
			Simon.UsePoison()

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

func initCharacter(nom string, classe string, lvl int, hpmax int, hp int, inv []Objects, money int, skills []string) Character {
	return Character{Nom: nom, Classe: classe, Lvl: lvl, HpMax: hpmax, Hp: hp, inv: inv, Money: money, Skills: skills}
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
	fmt.Println("            		  		    Inventory:")
	fmt.Println("----------------------------------------------------------------------------------------------------")

	slots := 16 // 4x4 cases fixes
	perRow := 4

	for i := 0; i < slots; i++ {
		if i < len(c.inv) {
			obj := c.inv[i]
			// Affiche nom + quantité de l'objet — bigger space now (20)
			text := Espace(20, obj.nom, fmt.Sprintf("x%d", obj.quantity))
			fmt.Printf("[ %s ]", text)
		} else {
			// Case vide
			text := Espace(20, "", "")
			fmt.Printf("[ %s ]", text)
		}
		// retour ligne après 4 cases
		if (i+1)%perRow == 0 {
			fmt.Println()
			fmt.Println("----------------------------------------------------------------------------------------------------")
		}
	}
	fmt.Println("\n")
}

func (c Character) accessInventory2() {
	fmt.Println("\n")
	fmt.Println("            				    Inventory:")
	fmt.Println("████████████████████████████████████████████████████████████████████████████████████████████████")

	slots := 16 // 4x4 cases fixes
	perRow := 4

	for i := 0; i < slots; i++ {
		if i < len(c.inv) {
			obj := c.inv[i]
			text := Espace(20, obj.nom, fmt.Sprintf("x%d", obj.quantity))
			fmt.Printf("█ %s █", text)
		} else {
			text := Espace(20, "", "")
			fmt.Printf("█ %s █", text)
		}
		if (i+1)%perRow == 0 {
			fmt.Println()
			fmt.Println("████████████████████████████████████████████████████████████████████████████████████████████████")
		}
	}
	fmt.Println("\n")
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

			if c.inv[i].quantity <= 0 {
				c.inv = append(c.inv[:i], c.inv[i+1:]...)
			}
			return
		}
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
		"    |   /-_      /-_     /-_      /      ",
		"    |_-/   \\-_-/   \\-_-/   \\-_-/|     ",
		"    |          _______             |     ",
		"    |         /       \\            |    ",
		"    |         | .   . |            |     ",
		"    |         |  -_-  |            |     ",
		"    |         \\-------/            |    ",
		"    |              |               |     ",
		"    |              |               |     ",
		"  ------------------------------------   ",
		"   | |============================| |    ",
		"   | |             |              | |    ",
		"   | |            / \\             | |   ",
		"   | |           /   \\            | |   ",
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
		{"6", "Épée B", 10},
		{"7", "Armure B", 10},
	}

MarchandLoop:

	for {
		fmt.Println("\n===================================")
		fmt.Println("Bienvenue chez le marchand !")
		fmt.Println("B) Acheter")
		fmt.Println("S) Vendre")
		fmt.Println("Q) Quitter le marchand")
		fmt.Println("===================================")

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
					log.Fatal(err)
				}

				switch char {
				case '1', '&':
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

				case '2', 'é':
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

				case '3', '"':
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

				case '5', '(':
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

				case '6', '-':
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

				case '7', 'è':
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

				case '8', '_':
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
			lines := CombineColumnsToLines([][]string{building, contentV}, 4)
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
				case '2', 'é':
					for i := 0; i < len(c.inv); i++ {
						if c.inv[i].nom == "Potion de poison" && c.inv[i].quantity > 0 {
							c.inv[i].quantity--
							c.Money++
						}
						if c.inv[i].nom == "Potion de poison" && c.inv[i].quantity == 0 {
							c.removeInventory(c.inv[i])
						}
					}

				case '3', '"':
					for i := 0; i < len(c.inv); i++ {
						if c.inv[i].nom == "Épée C" && c.inv[i].quantity > 0 {
							c.inv[i].quantity--
							c.Money++
						}
						if c.inv[i].nom == "Épée C" && c.inv[i].quantity == 0 {
							c.removeInventory(c.inv[i])
						}
					}

				case '5', '(':
					for i := 0; i < len(c.inv); i++ {
						if c.inv[i].nom == "Armor C" && c.inv[i].quantity > 0 {
							c.inv[i].quantity--
							c.Money++
						}
						if c.inv[i].nom == "Armor C" && c.inv[i].quantity == 0 {
							c.removeInventory(c.inv[i])
						}
					}

				case '6', '-':
					for i := 0; i < len(c.inv); i++ {
						if c.inv[i].nom == "Épée B" && c.inv[i].quantity > 0 {
							c.inv[i].quantity--
							c.Money++
						}
						if c.inv[i].nom == "Épée B" && c.inv[i].quantity == 0 {
							c.removeInventory(c.inv[i])
						}
					}

				case '7', 'è':
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

	maxLines := 0
	for _, col := range cols {
		if len(col) > maxLines {
			maxLines = len(col)
		}
	}

	sep := strings.Repeat(" ", distance)

	for r := 0; r < maxLines; r++ {
		for ci, col := range cols {
			cell := ""
			if r < len(col) {
				cell = col[r]
			}
			fmt.Printf("%-*s", widths[ci], cell)
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

func CombineColumnsToLines(cols [][]string, distance int) []string {
	widths := make([]int, len(cols))
	for ci, col := range cols {
		for _, line := range col {
			if l := utf8.RuneCountInString(line); l > widths[ci] {
				widths[ci] = l
			}
		}
	}
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
			parts = append(parts, fmt.Sprintf("%-*s", widths[ci], cell))
		}
		out = append(out, strings.Join(parts, sep))
	}
	return out
}

// FullScreenDrawCentered clears terminal and prints lines centered horizontally & vertically.
func FullScreenDrawCentered(lines []string) {
	fmt.Print("\033[?25l")
	fmt.Print("\033[2J\033[H")

	// get terminal size
	w, h, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		for _, ln := range lines {
			fmt.Println(ln)
		}
		fmt.Print("\033[?25h")
		return
	}

	for i, ln := range lines {
		if utf8.RuneCountInString(ln) > w {
			lines[i] = truncateRunes(ln, w)
		}
	}

	if len(lines) > h {
		lines = lines[:h]
	}
	topPad := (h - len(lines)) / 2

	for i := 0; i < topPad; i++ {
		fmt.Println()
	}

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

	printed := topPad + len(lines)
	for printed < h {
		fmt.Println()
		printed++
	}

	fmt.Print("\033[?25h")
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
