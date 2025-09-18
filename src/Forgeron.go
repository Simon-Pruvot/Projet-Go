package main

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

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

	lgb := []string{
		"         .",
		"        ('",
		"        '|",
		"        |'",
		"       [::]",
		"       [::]   _......_",
		"       [::].-'      _.-`.",
		"       [:.'    .-. '-._.-`.",
		"       [/ /\\   |  \\        `-..",
		"       / / |   `-.'      .-.   `-.",
		"      /  `-'            (   `.    `.",
		"     |           /\\      `-._/      \\",
		"     '    .'\\   /  `.           _.-'|",
		"    /    /  /   \\_.-'        _.':;:/",
		"  .'     \\_/             _.-':;_.-'",
		" /   .-.             _.-' \\;.-'",
		"/   (   \\       _..-'     |",
		"\\    `._/  _..-'    .--.  |",
		" `-.....-'/  _ _  .'    '.|",
		"          | |_|_| |      | \\  (o)",
		"     (o)  | |_|_| |      | | (\\'/)",
		"    (\\'/)/  ''''' |     o|  \\;:;",
		"     :;  |        |      |  |/)",
		" LGB  ;: `-.._    /__..--'\\.' ;:",
		"          :;  `--' :;   :;",
	}

	content := []string{
		"",
		"",
		"",
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

	lines := CombineColumnsToLines([][]string{lgb, content}, 4)
	FullScreenDrawCentered(lines)

ForgeronLoop:
	for {
		char, _, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		switch char {
		case '1', '&':
			fmt.Println("🔨 Ressources nécessaires pour Épée A :")
			for res, qty := range reqs[swordLegend.nom] {
				fmt.Printf("   - %s x%d\n", res, qty)
			}
			c.craftItem(swordLegend, reqs[swordLegend.nom])

		case '2', 'é':
			fmt.Println("🔨 Ressources nécessaires pour Armor A :")
			for res, qty := range reqs[armorLegend.nom] {
				fmt.Printf("   - %s x%d\n", res, qty)
			}
			c.craftItem(armorLegend, reqs[armorLegend.nom])

		case '3', '"':
			fmt.Println("🔨 Ressources nécessaires pour Potion Légendaire :")
			for res, qty := range reqs[potionLegend.nom] {
				fmt.Printf("   - %s x%d\n", res, qty)
			}
			c.craftItem(potionLegend, reqs[potionLegend.nom])

		case '4', '\'':
			fmt.Println("🔨 Ressources nécessaires pour Potion de vie :")
			for res, qty := range reqs["Potion de vie"] {
				fmt.Printf("   - %s x%d\n", res, qty)
			}
			c.craftItem(Objects{"Potion de vie", 1}, reqs["Potion de vie"])

		case 'q', 'Q':
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
