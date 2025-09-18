package main

import "fmt"

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

// Fonction pour ajouter un objet à l'inventaire
func (c *Character) addInventory(obj Objects) {
	if c.canAddItem() {
		c.inv = append(c.inv, obj)
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
	return Character{Nom: nom, Classe: classe, Lvl: lvl, HpMax: hpmax, Hp: hp, inv: inv, Money: 100, Skills: skills, Equipment: equipment, MaxInventorySlots: 10, InventoryUpgrades: 0, Mana: 30, ManaMax: 30, bonusDMG: 0, Initiative: 5}
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

	// Handle replacement and assign new equipment
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
	case "weapon":
		replaced = c.Equipment.Weapon
		c.Equipment.Weapon = &item

	default:
		fmt.Println("⚠️ Slot inconnu :", slot)
		return
	}

	// Recalculate HP
	c.updateHpMax(baseHpMax)

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

	if c.Equipment.Chapeau != nil && c.Equipment.Chapeau.nom == "Chapeau A" {
		c.HpMax += 15
	} else if c.Equipment.Chapeau != nil && c.Equipment.Chapeau.nom == "Chapeau B" {
		c.HpMax += 10
	} else if c.Equipment.Chapeau != nil && c.Equipment.Chapeau.nom == "Chapeau C" {
		c.HpMax += 5
	}

	if c.Equipment.Tunique != nil && c.Equipment.Tunique.nom == "Tunique A" {
		c.HpMax += 35
	} else if c.Equipment.Tunique != nil && c.Equipment.Tunique.nom == "Tunique B" {
		c.HpMax += 25
	} else if c.Equipment.Tunique != nil && c.Equipment.Tunique.nom == "Tunique C" {
		c.HpMax += 15
	}

	if c.Equipment.Bottes != nil && c.Equipment.Bottes.nom == "Bottes A" {
		c.HpMax += 25
	} else if c.Equipment.Bottes != nil && c.Equipment.Bottes.nom == "Bottes B" {
		c.HpMax += 15
	} else if c.Equipment.Bottes != nil && c.Equipment.Bottes.nom == "Bottes C" {
		c.HpMax += 5
	}
	if c.Equipment.Weapon != nil {
		switch c.Equipment.Weapon.nom {
		case "Épée A":
			c.bonusDMG = 15
		case "Épée B":
			c.bonusDMG = 10
		case "Épée C":
			c.bonusDMG = 5
		}
	}

	// Adjust current HP if it exceeds new max
	if c.Hp > c.HpMax {
		c.Hp = c.HpMax
	}
}

func (c *Character) isDead() {
	if c.Hp <= 0 {
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
			} else if objName == "Livre : Coup de poing" {
				c.spellBook("Coup de poing")
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
