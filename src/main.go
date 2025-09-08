package main

import "fmt"

type Objects struct {
	nom      string
	quantity int
}

type Character struct {
	Nom    string
	Classe string
	Lvl    int
	HpMax  int
	Hp     int
	inv    []Objects
}

func main() {
	initCharacter("Simon", "Elfe", 1, 100, 40, []Objects{{"potion", 3}})
	Simon.displayInfo()
}

func initCharacter(nom string, classe string, lvl int, hpmax int, hp int, inv Objects) {
	character := Character{nom, classe, lvl, hpmax, hp, []Objects{}}
}

func (c Character) displayInfo() {
	fmt.Println("------------------------------")
	fmt.Println("\tNom : " + c.Nom)
	fmt.Println("\tClasse : " + c.Classe)
	fmt.Println("Lvl : %d", c.Lvl)
	fmt.Println("Hp Max : %d", c.HpMax)
	fmt.Println("Hp : %d", c.Hp)
}
