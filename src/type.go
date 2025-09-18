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
	Mana              int
	ManaMax           int
	XP                int
	bonusDMG          int
}

type Monster struct {
	Nom   string
	HpMax int
	Hp    int
	Atk   int
}

var objects = map[string]Objects{
	"Armor A": {nom: "Armor A", quantity: 1},
	"Armor B": {nom: "Armor B", quantity: 1},
	"Armor C": {nom: "Armor C", quantity: 1},

	"Chapeau A": {nom: "Chapeau A", quantity: 1},
	"Tunique A": {nom: "Tunique A", quantity: 1},
	"Bottes A":  {nom: "Bottes A", quantity: 1},

	"Chapeau B": {nom: "Chapeau B", quantity: 1},
	"Tunique B": {nom: "Tunique B", quantity: 1},
	"Bottes B":  {nom: "Bottes B", quantity: 1},

	"Chapeau C": {nom: "Chapeau C", quantity: 1},
	"Tunique C": {nom: "Tunique C", quantity: 1},
	"Bottes C":  {nom: "Bottes C", quantity: 1},
}
