package main

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

func initCharacter(nom string, classe string, lvl int, hpmax int, hp int, inv Objects) {
	character := Character{"Simon", "Elfe", 1, 100, 40, []Objects{{"potion", 3}}}
}
