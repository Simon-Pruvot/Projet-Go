package main

import (
	"fmt"
	"time"
)

func (c *Character) UsePoison() {
	for i := 0; i < len(c.inv); i++ {
		if c.inv[i].nom == "Potion de poison" && c.inv[i].quantity > 0 {
			c.inv[i].quantity--
			c.Hp -= 10
			fmt.Println("You make damage to your enemi")
			fmt.Println(``)
			time.Sleep(1 * time.Second)
			c.Hp -= 10
			fmt.Println("You make damage to your enemi")
			fmt.Println(``)
			time.Sleep(1 * time.Second)
			c.Hp -= 10
			fmt.Println("You make damage to your enemi")
			fmt.Println(``)
			if c.Hp > c.HpMax {
				c.Hp = c.HpMax
			}
			fmt.Println("End of poison")
			fmt.Println(``)

			// if no more left → remove it
			if c.inv[i].quantity == 0 {
				c.inv = append(c.inv[:i], c.inv[i+1:]...)
			}
			return
		}
	}
	fmt.Println("⚠️ No potion of poison left!")
}
