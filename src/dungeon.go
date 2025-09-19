package main

import (
	"fmt"
	"math/rand"
	"time"
)

func trainingFight(player *Character, goblin Monster) {
	turn := 1
	fmt.Printf("=== Début de l'entraînement contre %s ===\n", goblin.Nom)

	rand.Seed(time.Now().UnixNano())

	lootPool := []Objects{
		{"Fourrure de Loup", 1},
		{"Peau de Troll", 1},
		{"Cuir de Sanglier", 1},
		{"Plume de Corbeau", 1},
	}

	for player.Hp > 0 && goblin.Hp > 0 {
		fmt.Printf("\n---- Tour %d ----\n", turn)
		player.Mana += 10
		if player.Mana > player.ManaMax {
			player.Mana = player.ManaMax
		}
		combat0(player, &goblin)
		characterTurn(player, &goblin)
		goblinPattern(&goblin, player, turn)

		if goblin.Hp <= 0 {
			fmt.Printf("%s est vaincu !\n", goblin.Nom)

			// Récompenses fixes
			GainXP(player, 50*player.Lvl)
			player.Money += 20

			// Loot aléatoire
			randomLoot := lootPool[rand.Intn(len(lootPool))]
			player.inv = append(player.inv, Objects{"Cuir de Sanglier", 1}) // loot 100%
			player.inv = append(player.inv, randomLoot)

			fmt.Printf("🎁 Récompenses : 50 XP, 20 or, +1 %s\n", randomLoot.nom)
			fmt.Println("Entrée une lettre pour continuer...")
			var input string
			fmt.Scanln(&input)
		}
		turn++
	}
	if player.Hp <= 0 {
		mort()
		player.isDead()
		fmt.Println("Entrée une lettre pour continuer...")
		var input string
		fmt.Scanln(&input)
		return
	}
}

func Fight1(player *Character, goblin Monster) {
	turn := 1
	fmt.Printf("=== Début de l'entraînement contre %s ===\n", goblin.Nom)

	// Seed pour la RNG (évite d’avoir toujours le même résultat)
	rand.Seed(time.Now().UnixNano())

	lootPool := []Objects{
		{"Fourrure de Loup", 2},
		{"Peau de Troll", 3},
		{"Cuir de Sanglier", 1},
		{"Plume de Corbeau", 2},
	}

	for player.Hp > 0 && goblin.Hp > 0 {
		fmt.Printf("\n---- Tour %d ----\n", turn)
		player.Mana += 10
		if player.Mana > player.ManaMax {
			player.Mana = player.ManaMax
		}
		combat(player, &goblin)
		characterTurn(player, &goblin)
		goblinPattern(&goblin, player, turn)

		if goblin.Hp <= 0 {
			fmt.Printf("%s est vaincu !\n", goblin.Nom)

			// Récompenses fixes
			GainXP(player, 50*player.Lvl)
			player.Money += 20

			// Loot aléatoire
			randomLoot := lootPool[rand.Intn(len(lootPool))]
			player.inv = append(player.inv, Objects{"Cuir de Sanglier", 1}) // loot 100%
			player.inv = append(player.inv, randomLoot)

			fmt.Printf("🎁 Récompenses : 50 XP, 20 or, +1 %s\n", randomLoot.nom)
			fmt.Println("Entrée une lettre pour continuer...")
			var input string
			fmt.Scanln(&input)
		}

		turn++
	}

	if player.Hp <= 0 {
		mort()
		player.isDead()
		fmt.Println("Entrée une lettre pour continuer...")
		var input string
		fmt.Scanln(&input)
		return
	}
}

func Fight2(player *Character, goblin Monster) {
	turn := 1
	fmt.Printf("=== Début de l'entraînement contre %s ===\n", goblin.Nom)

	// Seed pour la RNG (évite d’avoir toujours le même résultat)
	rand.Seed(time.Now().UnixNano())

	lootPool := []Objects{
		{"Fourrure de Loup", 2},
		{"Peau de Troll", 3},
		{"Cuir de Sanglier", 1},
		{"Plume de Corbeau", 2},
	}

	for player.Hp > 0 && goblin.Hp > 0 {
		fmt.Printf("\n---- Tour %d ----\n", turn)
		player.Mana += 10
		if player.Mana > player.ManaMax {
			player.Mana = player.ManaMax
		}
		combat2(player, &goblin)
		characterTurn(player, &goblin)
		goblinPattern(&goblin, player, turn)

		if goblin.Hp <= 0 {
			fmt.Printf("%s est vaincu !\n", goblin.Nom)

			// Récompenses fixes
			GainXP(player, 50*player.Lvl)
			player.Money += 20

			// Loot aléatoire
			randomLoot := lootPool[rand.Intn(len(lootPool))]
			player.inv = append(player.inv, Objects{"Cuir de Sanglier", 1}) // loot 100%
			player.inv = append(player.inv, randomLoot)

			fmt.Printf("🎁 Récompenses : 50 XP, 20 or, +1 %s\n", randomLoot.nom)
			fmt.Println("Entrée une lettre pour continuer...")
			var input string
			fmt.Scanln(&input)
		}

		turn++
	}

	if player.Hp <= 0 {
		mort()
		player.isDead()
		fmt.Println("Entrée une lettre pour continuer...")
		var input string
		fmt.Scanln(&input)
		return
	}
}

func Fight3(player *Character, goblin Monster) {
	turn := 1
	fmt.Printf("=== Début de l'entraînement contre %s ===\n", goblin.Nom)

	// Seed pour la RNG (évite d’avoir toujours le même résultat)
	rand.Seed(time.Now().UnixNano())

	// Définition du loot possible
	lootPool := []Objects{
		{"Fourrure de Loup", 2},
		{"Peau de Troll", 3},
		{"Cuir de Sanglier", 1},
		{"Plume de Corbeau", 2},
	}

	for player.Hp > 0 && goblin.Hp > 0 {
		fmt.Printf("\n---- Tour %d ----\n", turn)
		player.Mana += 10
		if player.Mana > player.ManaMax {
			player.Mana = player.ManaMax
		}
		combat3(player, &goblin)
		characterTurn(player, &goblin)
		goblinPattern(&goblin, player, turn)

		if goblin.Hp <= 0 {
			fmt.Printf("%s est vaincu !\n", goblin.Nom)

			// Récompenses fixes
			GainXP(player, 50*player.Lvl)
			player.Money += 20

			// Loot aléatoire
			randomLoot := lootPool[rand.Intn(len(lootPool))]
			player.inv = append(player.inv, Objects{"Cuir de Sanglier", 1}) // loot 100%
			player.inv = append(player.inv, randomLoot)

			fmt.Printf("🎁 Récompenses : 50 XP, 20 or, +1 %s\n", randomLoot.nom)
			lorefin()

			fmt.Println("Appuyez sur Entrée pour continuer...")
			fmt.Scanln()
		}

		if player.Hp <= 0 {
			mort()
			player.isDead()
			fmt.Println("Appuyez sur Entrée pour continuer...")
			fmt.Scanln()
			return
		}

		turn++
	}
}
