package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func (c *Character) initGoblin(tip string) Monster {
	rand.Seed(time.Now().UnixNano())

	switch tip {
	case "Gobelin":
		return Monster{
			Nom:        "Gobelin d'entraînement",
			HpMax:      40 * c.Lvl,
			Hp:         40 * c.Lvl,
			Atk:        5 * c.Lvl,
			Initiative: rand.Intn(10) + 1,
		}
	case "dead":
		return Monster{
			Nom:        "Death",
			HpMax:      80 * c.Lvl,
			Hp:         80 * c.Lvl,
			Atk:        7 * c.Lvl,
			Initiative: rand.Intn(10) + 1,
		}
	case "skeleton":
		return Monster{
			Nom:        "Skeleton",
			HpMax:      100 * c.Lvl,
			Hp:         100 * c.Lvl,
			Atk:        8 * c.Lvl,
			Initiative: rand.Intn(10) + 1,
		}
	case "dragon":
		return Monster{
			Nom:        "Dragon",
			HpMax:      150 * c.Lvl,
			Hp:         150 * c.Lvl,
			Atk:        10 * c.Lvl,
			Initiative: rand.Intn(10) + 1,
		}
	default:
		return Monster{
			Nom:        "Gobelin d'entraînement",
			HpMax:      40 * c.Lvl,
			Hp:         40 * c.Lvl,
			Atk:        5 * c.Lvl,
			Initiative: rand.Intn(10) + 1,
		}
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
		m.Hp -= (5 + c.bonusDMG)
		if m.Hp < 0 {
			m.Hp = 0
		}
		c.Initiative += 2 // gain d'initiative pour attaque simple
		if c.Initiative > 10 {
			c.Initiative = 10
		}
		fmt.Printf("%s utilise Attaque basique et inflige %d dégâts à %s\n", c.Nom, (5 + c.bonusDMG), m.Nom)

	case "s":
		// Spéciale : 12 dégâts, coût 20 mana
		if c.Mana < 20 {
			fmt.Println("Pas assez de mana pour attaque spéciale.")
			return false
		}
		c.Mana -= 20
		m.Hp -= (12 + c.bonusDMG)
		if m.Hp < 0 {
			m.Hp = 0
		}
		c.Initiative -= 1 // perte d'initiative pour spéciale
		if c.Initiative < 0 {
			c.Initiative = 0
		}
		fmt.Printf("%s utilise Attaque spéciale et inflige %d dégâts à %s\n", c.Nom, (12 + c.bonusDMG), m.Nom)

	case "m":
		fmt.Println("Coup de poing:1 | Boule de feu:2")
		// Magie : uniquement si on a appris au moins un sort
		if len(c.Skills) == 0 {
			fmt.Println("Vous n'avez appris aucun sort.")
			return false
		}
		if c.Mana < 30 {
			fmt.Println("Pas assez de mana pour lancer un sort.")
			return false
		}
		for {
			var inputspell string
			_, err := fmt.Scanln(&inputspell)
			if err != nil {
				fmt.Println("Entrée invalide. Réessayez.")
				continue
			}

			valid := false
			switch inputspell {
			case "é", "2":
				if len(c.Skills) < 2 || c.Mana < 30 {
					fmt.Println("Choix invalide ou pas assez de mana.(1/2)")
					continue
				}
				spell := c.Skills[1]
				c.Mana -= 30
				m.Hp -= (18 + c.bonusDMG)
				if m.Hp < 0 {
					m.Hp = 0
				}
				fmt.Printf("%s lance %s et inflige %d dégâts à %s\n", c.Nom, spell, (18 + c.bonusDMG), m.Nom)
				valid = true
				c.Initiative -= 2 // perte d'initiative pour sort

			case "&", "1":
				if len(c.Skills) < 1 || c.Mana < 30 {
					fmt.Println("Choix invalide ou pas assez de mana.(1/2)")
					continue
				}
				spell := c.Skills[0]
				c.Mana -= 15
				m.Hp -= (8 + c.bonusDMG)
				if m.Hp < 0 {
					m.Hp = 0
				}
				fmt.Printf("%s lance %s et inflige %d dégâts à %s\n", c.Nom, spell, (8 + c.bonusDMG), m.Nom)
				valid = true
				c.Initiative -= 2 // perte d'initiative pour sort

			default:
				fmt.Println("Option invalide, réessayez.")
			}

			if valid {
				if c.Initiative < 0 {
					c.Initiative = 0
				} else if c.Initiative > 10 {
					c.Initiative = 10
				}
				break
			}
		}

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
		switch item.nom {
		case "Potion de vie":
			c.Hp += 50
			if c.Hp > c.HpMax {
				c.Hp = c.HpMax
			}
			c.Initiative += 1 // gain d'initiative pour potions
		case "Potion de poison":
			c.UsePoison(m)
			fmt.Printf("Vous utilisez %s. PV : %d / %d\n", item.nom, c.Hp, c.HpMax)
			c.Initiative += 1
		case "Livre de sort":
			c.Skills = append(c.Skills, "Boule de feu")
			item.quantity--
			fmt.Println("Vous apprenez un nouveau sort : Boule de feu !")
			c.Initiative += 1
		default:
			fmt.Printf("Vous utilisez %s (effet non implémenté).\n", item.nom)
			item.quantity--
			c.Initiative += 1
		}

		if item.quantity <= 0 {
			c.inv = append(c.inv[:idx], c.inv[idx+1:]...)
		} else {
			c.inv[idx].quantity = item.quantity
		}

		if c.Initiative > 10 {
			c.Initiative = 10
		}

	case "k":
		// Skip
		fmt.Printf("%s choisit de concentrer son énergie et passe son tour.\n", c.Nom)
		c.Initiative += 3
		if c.Initiative > 10 {
			c.Initiative = 10
		}

	default:
		fmt.Println("Option inconnue, tour passé.")
	}

	// S'assurer que l'initiative est toujours entre 0 et 10
	if c.Initiative < 0 {
		c.Initiative = 0
	} else if c.Initiative > 10 {
		c.Initiative = 10
	}

	return m.Hp <= 0
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
