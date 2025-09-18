package main

import (
	"fmt"
	"math/rand"
	"time"
)

func trainingFight(player *Character, goblin Monster) {
	turn := 1
	fmt.Printf("=== Début de l'entraînement contre %s ===\n", goblin.Nom)

	// Seed pour la RNG (évite d’avoir toujours le même résultat)
	rand.Seed(time.Now().UnixNano())
	combat0(player, &goblin)
	// Définition du loot possible
	lootPool := []Objects{
		{"Fourrure de Loup", 1},
		{"Peau de Troll", 1},
		{"Cuir de Sanglier", 1},
		{"Plume de Corbeau", 1},
	}

	for player.Hp > 0 && goblin.Hp > 0 {
		fmt.Printf("\n---- Tour %d ----\n", turn)
		characterTurn(player, &goblin) // joueur joue
		combat0(player, &goblin)

		if goblin.Hp <= 0 {
			fmt.Printf("%s est vaincu !\n", goblin.Nom)

			// Récompenses fixes
			GainXP(player, 50)
			player.Money += 20

			// Loot aléatoire
			randomLoot := lootPool[rand.Intn(len(lootPool))]
			player.inv = append(player.inv, Objects{"Cuir de Sanglier", 1}) // loot 100%
			player.inv = append(player.inv, randomLoot)

			fmt.Printf("🎁 Récompenses : 50 XP, 20 or, +1 %s\n", randomLoot.nom)
			break
		}

		goblinPattern(&goblin, player, turn) // gobelin joue
		turn++
	}

	if player.Hp <= 0 {
		mort()
		player.isDead()
		fmt.Println("Vous avez été vaincu... retour au menu.")
		return
	}
}

func Fight1(player *Character, goblin Monster) {
	turn := 1
	fmt.Printf("=== Début de l'entraînement contre %s ===\n", goblin.Nom)

	// Seed pour la RNG (évite d’avoir toujours le même résultat)
	rand.Seed(time.Now().UnixNano())
	combat(player, &goblin)
	// Définition du loot possible
	lootPool := []Objects{
		{"Fourrure de Loup", 2},
		{"Peau de Troll", 3},
		{"Cuir de Sanglier", 1},
		{"Plume de Corbeau", 2},
	}

	for player.Hp > 0 && goblin.Hp > 0 {
		fmt.Printf("\n---- Tour %d ----\n", turn)
		characterTurn(player, &goblin) // joueur joue
		combat(player, &goblin)

		if goblin.Hp <= 0 {
			fmt.Printf("%s est vaincu !\n", goblin.Nom)

			// Récompenses fixes
			GainXP(player, 50)
			player.Money += 20

			// Loot aléatoire
			randomLoot := lootPool[rand.Intn(len(lootPool))]
			player.inv = append(player.inv, Objects{"Cuir de Sanglier", 1}) // loot 100%
			player.inv = append(player.inv, randomLoot)

			fmt.Printf("🎁 Récompenses : 50 XP, 20 or, +1 %s\n", randomLoot.nom)
			break
		}

		goblinPattern(&goblin, player, turn) // gobelin joue
		turn++
	}

	if player.Hp <= 0 {
		mort()
		player.isDead()
		fmt.Println("Vous avez été vaincu... retour au menu.")
		return
	}
}

func Fight2(player *Character, goblin Monster) {
	turn := 1
	fmt.Printf("=== Début de l'entraînement contre %s ===\n", goblin.Nom)

	// Seed pour la RNG (évite d’avoir toujours le même résultat)
	rand.Seed(time.Now().UnixNano())
	combat2(player, &goblin)
	// Définition du loot possible
	lootPool := []Objects{
		{"Fourrure de Loup", 2},
		{"Peau de Troll", 3},
		{"Cuir de Sanglier", 1},
		{"Plume de Corbeau", 2},
	}

	for player.Hp > 0 && goblin.Hp > 0 {
		fmt.Printf("\n---- Tour %d ----\n", turn)
		characterTurn(player, &goblin) // joueur joue
		combat2(player, &goblin)

		if goblin.Hp <= 0 {
			fmt.Printf("%s est vaincu !\n", goblin.Nom)

			// Récompenses fixes
			GainXP(player, 50)
			player.Money += 20

			// Loot aléatoire
			randomLoot := lootPool[rand.Intn(len(lootPool))]
			player.inv = append(player.inv, Objects{"Cuir de Sanglier", 1}) // loot 100%
			player.inv = append(player.inv, randomLoot)

			fmt.Printf("🎁 Récompenses : 50 XP, 20 or, +1 %s\n", randomLoot.nom)
			break
		}

		goblinPattern(&goblin, player, turn) // gobelin joue
		turn++
	}

	if player.Hp <= 0 {
		mort()
		player.isDead()
		fmt.Println("Vous avez été vaincu... retour au menu.")
		return
	}
}

func Fight3(player *Character, goblin Monster) {
	turn := 1
	fmt.Printf("=== Début de l'entraînement contre %s ===\n", goblin.Nom)

	// Seed pour la RNG (évite d’avoir toujours le même résultat)
	rand.Seed(time.Now().UnixNano())
	combat3(player, &goblin)
	// Définition du loot possible
	lootPool := []Objects{
		{"Fourrure de Loup", 2},
		{"Peau de Troll", 3},
		{"Cuir de Sanglier", 1},
		{"Plume de Corbeau", 2},
	}

	for player.Hp > 0 && goblin.Hp > 0 {
		fmt.Printf("\n---- Tour %d ----\n", turn)
		characterTurn(player, &goblin) // joueur joue
		combat3(player, &goblin)

		if goblin.Hp <= 0 {
			fmt.Printf("%s est vaincu !\n", goblin.Nom)

			// Récompenses fixes
			GainXP(player, 50)
			player.Money += 20

			// Loot aléatoire
			randomLoot := lootPool[rand.Intn(len(lootPool))]
			player.inv = append(player.inv, Objects{"Cuir de Sanglier", 1}) // loot 100%
			player.inv = append(player.inv, randomLoot)

			fmt.Printf("🎁 Récompenses : 50 XP, 20 or, +1 %s\n", randomLoot.nom)
			break
		}

		goblinPattern(&goblin, player, turn) // gobelin joue
		turn++
	}

	if player.Hp <= 0 {
		mort()
		player.isDead()
		fmt.Println("Vous avez été vaincu... retour au menu.")
		return
	}
}
