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

func main() {
	//J'appelle la fonction de l'écran de départ
	TextBienvenu()

	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	fmt.Println("Press I to open inventory, H to drink potion, P to pause, D to display info, Q to quit.")

	player := classe()

	for {
		char, _, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		switch char {
		case 'i', 'I':
			player.accessInventory2()
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
					player.accessInventory2()
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
		case '9', 'ç':
			player.UsePoison()

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

func (c Character) displayInfo() {
	fmt.Println("")
	fmt.Println("------------------------------")
	fmt.Println("------------------------------")
	fmt.Println("\tNom : " + c.Nom)
	fmt.Println("\tClasse : " + c.Classe)
	fmt.Println("\tLvl :", c.Lvl)
	fmt.Println("\tHp Max :", c.HpMax)
	fmt.Println("\tHp :", c.Hp)
	fmt.Println("------------------------------")
	fmt.Println("------------------------------")
	fmt.Println("")
}

func (c Character) accessInventory() {
	fmt.Println("")
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
	fmt.Println("")
}

func (c Character) accessInventory2() {
	fmt.Println("")
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
	fmt.Println("")
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
	//objects
	hpPot := Objects{nom: "Potion de vie", quantity: 1}
	poisonPot := Objects{nom: "Potion de poison", quantity: 1}
	swordCom := Objects{nom: "Épée C", quantity: 1}
	swordRare := Objects{nom: "Épée B", quantity: 1}
	//swordLegend := Objects{nom: "Épée A", quantity: 1}
	armorCom := Objects{nom: "Armor C", quantity: 1}
	armorRare := Objects{nom: "Armor B", quantity: 1}
	//armorLegend := Objects{nom: "Armor A", quantity: 1}
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
		{"a", "Rock", 1},
		{"b", "Wood", 1},
		{"c", "Scrap", 5},
		{"d", "Fourrure de Loup", 4},
		{"e", "Peau de Troll", 7},
		{"f", "Cuir de Sanglier", 3},
		{"G", "Plume de Corbeau", 1},
		{"1", "Potion de vie", 1},
		{"2", "Potion de poison", 1},
		{"3", "Épée C", 5},
		{"5", "Armure C", 5},
		{"6", "Épée B", 10},
		{"7", "Armure B", 10},
	}

	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	for {
		char, _, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("\n===================================")
		fmt.Println("Bienvenue chez le marchand !")
		fmt.Println("A) Acheter")
		fmt.Println("V) Vendre")
		fmt.Println("Q) Quitter le marchand")
		fmt.Println("===================================")

		switch char {
		case 'b', 'B':
			fmt.Println("Chose an item: 1, 2, 3, 5, 6, 7, 8")
			fmt.Println("Q to leave the marchant")

			for {
				char, _, err := keyboard.GetKey()
				if err != nil {
					log.Fatal(err)
				}

				if char == 'a' || char == 'A' {
					fmt.Println("\nObjets disponibles à l'achat :")
					fmt.Println("    ______________________________")
					fmt.Println("   / \\                             \\.")
					fmt.Println("  |   |         Marchand           |.")
					fmt.Println("   \\_ |                            |.")
					for _, it := range achats {
						fmt.Printf("     |   %s) %-25s %2d Gold |\n", it.key, it.nom, it.price)
					}
					fmt.Println("     |                            |.")
					fmt.Println("     |   _________________________|___")
					fmt.Println("     |  /                            /.")
					fmt.Println("     \\_/dc__________________________/.")

					fmt.Println("\nAppuyez sur la touche correspondant à l’objet pour l’acheter, ou Q pour revenir.")
				}

				switch char {
				case '1', '&': // Acheter une potion de soin pour 2 Gold
					if len(c.inv) < 1 {
						c.addInventory(hpPot)
					}
					for i := 0; i < len(c.inv); i++ {
						if c.inv[i].nom == "Potion de vie" && c.inv[i].quantity > 0 {
						} else {
							c.addInventory(hpPot)
						}
					}
					if c.Money >= 3 {
						c.Money -= 3
						for i := 0; i < len(c.inv); i++ {
							if c.inv[i].nom == "Potion de vie" {
								c.inv[i].quantity++
							} else {
								c.addInventory(hpPot)
							}
						}
					} else {
						fmt.Println("Need More Gold")
					}

				case '2', 'é': // Acheter une potion de poison pour 3 Gold
					if c.Money >= 6 {
						c.Money -= 6
						for i := 0; i < len(c.inv); i++ {
							if c.inv[i].nom == "Potion de poison" {
								c.inv[i].quantity++
							} else {
								c.addInventory(poisonPot)
							}
						}
					} else {
						fmt.Println("Need More Gold")
					}

				case '3', '"': // Acheter Sword C pour 10 Gold
					if c.Money >= 10 {
						c.Money -= 10
						for i := 0; i < len(c.inv); i++ {
							if c.inv[i].nom == "Épée C" {
								c.inv[i].quantity++
							} else {
								c.addInventory(swordCom)
							}
						}
					} else {
						fmt.Println("Need More Gold")
					}

				case '5', '(': // Acheter Armor C pour 10 Gold
					if c.Money >= 10 {
						c.Money -= 10
						for i := 0; i < len(c.inv); i++ {
							if c.inv[i].nom == "Armor C" {
								c.inv[i].quantity++
							} else {
								c.addInventory(armorCom)
							}
						}
					} else {
						fmt.Println("Need More Gold")
					}

				case '6', '-': // Acheter Sword B pour 20 Gold
					if c.Money >= 20 {
						c.Money -= 20
						for i := 0; i < len(c.inv); i++ {
							if c.inv[i].nom == "Épée B" {
								c.inv[i].quantity++
							} else {
								c.addInventory(swordRare)
							}
						}
					} else {
						fmt.Println("Need More Gold")
					}

				case '7', 'è': // Acheter Armor B pour 20 Gold
					if c.Money >= 20 {
						c.Money -= 20
						for i := 0; i < len(c.inv); i++ {
							if c.inv[i].nom == "Armor B" {
								c.inv[i].quantity++
							} else {
								c.addInventory(armorRare)
							}
						}
					} else {
						fmt.Println("Need More Gold")
					}
				case '8', '_': // Acheter Armor B pour 20 Gold
					if c.Money >= 25 {
						c.Money -= 25
						for i := 0; i < len(c.inv); i++ {
							if c.inv[i].nom == "Livre de Sort : Boule de feu" {
								c.inv[i].quantity++
							} else {
								c.addInventory(LivFeu)
							}
						}
					} else {
						fmt.Println("Need More Gold")
					}

				case 'q', 'Q':
					return
				}
			}

		case 's', 'S':
			fmt.Println("Chose an item: 1, 2, 3, 5, 6, 7")
			fmt.Println("Q to leave the marchant")

			for {
				char, _, err := keyboard.GetKey()
				if err != nil {
					log.Fatal(err)
				}

				if char == 'v' || char == 'V' {
					fmt.Println("\nObjets que vous pouvez vendre :")
					fmt.Println("    ______________________________")
					fmt.Println("   / \\                             \\.")
					fmt.Println("  |   |         Revente            |.")
					fmt.Println("   \\_ |                            |.")
					for _, it := range ventes {
						fmt.Printf("     |   %s) %-25s %2d Gold |\n", it.key, it.nom, it.price)
					}
					fmt.Println("     |                            |.")
					fmt.Println("     |   _________________________|___")
					fmt.Println("     |  /                            /.")
					fmt.Println("     \\_/dc__________________________/.")

					fmt.Println("\nAppuyez sur la touche correspondant à l’objet pour le vendre, ou Q pour revenir.")
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
					return
				}
			}

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
