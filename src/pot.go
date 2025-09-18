package main

import (
	"fmt"
	"time"
)

func (c *Character) UsePoison(target *Monster) {
	for i := 0; i < len(c.inv); i++ {
		if c.inv[i].nom == "Potion de poison" && c.inv[i].quantity > 0 {
			// Consommer la potion
			c.inv[i].quantity--

			// Dégâts sur 3 tours (10 PV chaque)
			for j := 0; j < 3; j++ {
				target.Hp -= 10
				if target.Hp < 0 {
					target.Hp = 0
				}
				fmt.Printf("☠️ Le poison inflige 10 dégâts à %s ! (PV restants : %d/%d)\n", target.Nom, target.Hp, target.HpMax)
				time.Sleep(1 * time.Second)
			}

			fmt.Println("☠️ Le poison s'est dissipé.")

			// Supprimer la potion si quantité = 0
			if c.inv[i].quantity == 0 {
				c.inv = append(c.inv[:i], c.inv[i+1:]...)
			}
			return
		}
	}
	fmt.Println("⚠️ Vous n’avez plus de potion de poison !")
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
