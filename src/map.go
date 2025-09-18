package main

func ShowMap() {
	Clear()
	mapMenu := []string{

		"                       ,-.^._                 _",
		"                     .'      `-.            ,' ;",
		"          /`-.  ,----'         `-.   _  ,-.,'  `",
		"       _.'   `--'                 `-' '-'      ;",
		"      :                         o             ;    __,-.",
		"      ,'    o            Marchand            ;_,-',.__'--.",
		"     :    Forgeron                           ,--```    `--'",
		"     :                                      ;",
		"     :                                      :",
		"     ;                o                     :",
		"    (              Dungeon                  ;",
		"     `-.                                  ,'",
		"       ;                                  :",
		"     .'                             .-._,'",
		"   .'                               `.",
		"_.''                                .__;",
		"`._                               ;",
		"   `.                            :    ,---------------------.",
		"     `.               ,..__,---._;    |    Marchand(B)      |",
		"       `-.__         :                |    Forgeron(F)      |",
		"            `.--.____;                |    Dungeon lvl 1(T) |",
		"                                      |    Dungeon lvl 2(N) |",
		"                                      |    Dungeon lvl 3(M) |",
		"                                      |    Dungeon lvl 4(L) |",
		"                                      |    Menu(P)          |",
		"                                      `---------------------'",
	}
	lines := CombineColumnsToLines([][]string{mapMenu}, 4)
	FullScreenDrawCentered(lines)

}
