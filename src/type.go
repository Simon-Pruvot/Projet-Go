package main

type Equipment struct {
	Chapeau *Objects
	Tunique *Objects
	Bottes  *Objects
}

type Objects struct {
	nom      string
	quantity int
}

type Character struct {
	Nom               string
	Classe            string
	Lvl               int
	HpMax             int
	Hp                int
	inv               []Objects
	Money             int
	Skills            []string
	Equipment         Equipment
	MaxInventorySlots int
	InventoryUpgrades int
}
