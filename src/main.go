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

	player := classe()

	if chosendif == "/start" {
	} else if chosendif == "/hard" {
		player.HpMax = player.HpMax / 2
		player.Hp = player.Hp / 2
	} else if chosendif == "/easy" {
	}
	fmt.Println("Press I to open inventory, H to drink potion, P to pause, D to display info, Q to quit.")

	for {
		char, _, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		switch char {
		case 'i', 'I':
			player.accessInventory()
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
					player.accessInventory()
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
		case 'f', 'F':
			player.Forgeron(swordLegend, armorLegend, potionLegend)
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
	lines := []string{
		"------------------------------",
		"------------------------------",
		fmt.Sprintf("Nom : %s", c.Nom),
		fmt.Sprintf("Classe : %s", c.Classe),
		fmt.Sprintf("Lvl : %d", c.Lvl),
		fmt.Sprintf("Hp Max : %d", c.HpMax),
		fmt.Sprintf("Hp : %d", c.Hp),
		"------------------------------",
		"------------------------------",
	}

	FullScreenDrawCentered(lines)
}

func (c Character) accessInventory() {
	slots := 10
	perRow := 5

	// Create rows (each row = []string of slots)
	var cols [][]string
	for row := 0; row < perRow; row++ {
		cols = append(cols, []string{})
	}

	for i := 0; i < slots; i++ {
		item := "[ empty ]"
		if i < len(c.inv) {
			item = c.inv[i].nom // or however you store item name
		}
		cols[i%perRow] = append(cols[i%perRow], item)
	}

	// Format inventory lines
	lines := []string{"Inventory:", strings.Repeat("-", 80)}
	grid := CombineColumnsToLines(cols, 4) // 4 spaces between slots
	lines = append(lines, grid...)
	lines = append(lines, strings.Repeat("-", 80))

	// Center display
	FullScreenDrawCentered(lines)
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
	if len(c.inv) >= 10 {
		fmt.Println("⚠️ Inventaire plein ! Impossible d’ajouter plus d’objets.")
		return false
	}
	return true
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
	}

	content := []string{
		"",
		"Bienvenue chez le Forgeron ⚒️",
		"Objets que vous pouvez fabriquer (5 Gold chacun) :",
		"  1 - Épée A",
		"  2 - Armor A",
		"  3 - Potion Légendaire",
		"",
		"Appuyez sur 1,2,3 pour fabriquer, ou Q pour quitter.",
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
