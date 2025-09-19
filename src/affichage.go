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

func Espace(esp int, chaine1 string, chaine2 string) string {
	esp -= (len(chaine1) + len(chaine2))
	for i := 0; i <= esp; i++ {
		chaine1 += " "
	}
	return chaine1 + chaine2 + string(rune(127))
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
		fmt.Sprintf("║ Money  : %-18d  ║", c.Money),
		"║                              ║",
		"╚══════════════════════════════╝",
	}

	FullScreenDrawCentered(lines)
}

func (c *Character) accessInventory(baseHpMax int) {
	slots := 10 * (c.InventoryUpgrades + 1)
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
				item = fmt.Sprintf("[%d] %s :%d", i, c.inv[i].nom, c.inv[i].quantity)
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

				switch {
				case strings.HasPrefix(item.nom, "Armor "):
					c.UseArmorSet(item.nom)
					// Remove the used set
					c.inv = append(c.inv[:index], c.inv[index+1:]...)

				case strings.HasPrefix(item.nom, "Livre "):
					c.useItem(item.nom)

				case item.nom == "Potion de vie":
					c.takePot()

				default:
					fmt.Println("⚠️ Cet objet ne peut pas être utilisé :", item.nom)
				}
			}

		case 'e', 'E':
			index := demanderIndex(len(c.inv))
			if index >= 0 && index < len(c.inv) {
				itemName := c.inv[index].nom

				switch {
				case len(itemName) >= 7 && itemName[:7] == "Chapeau":
					c.Equip(c.inv[index], "chapeau", baseHpMax)
					fmt.Println("✅ Équipé :", itemName)

				case len(itemName) >= 7 && itemName[:7] == "Tunique":
					c.Equip(c.inv[index], "tunique", baseHpMax)
					fmt.Println("✅ Équipé :", itemName)

				case len(itemName) >= 6 && itemName[:6] == "Bottes":
					c.Equip(c.inv[index], "bottes", baseHpMax)
					fmt.Println("✅ Équipé :", itemName)

				case strings.HasPrefix(itemName, "Épée"):
					c.Equip(c.inv[index], "weapon", baseHpMax)
					fmt.Println("✅ Équipé :", itemName)

				default:
					fmt.Println("⚠️ Choix invalide, vous ne pouvez équiper que Chapeau, Tunique, Bottes ou Épée.")
				}
			} else {
				fmt.Println("⚠️ Index invalide.")
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
	if c.Equipment.Weapon != nil {
		fmt.Println("  Weapon :", c.Equipment.Weapon.nom)
	} else {
		fmt.Println("  Bottes : Aucune")
	}
}

func (c Character) displaySkills() {
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
