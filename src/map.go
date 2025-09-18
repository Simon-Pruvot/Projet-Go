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
		"    (              Dangeon                  ;",
		"     `-.                                  ,'",
		"       ;                                  :",
		"     .'                             .-._,'",
		"   .'                               `.",
		"_.''                                .__;",
		"`._                               ;",
		"   `.                            :    ,---------------------.",
		"     `.               ,..__,---._;    |    Marchand(B)      |",
		"       `-.__         :                |    Forgeron(F)      |",
		"            `.--.____;                |    Dangeon lvl 1(T) |",
		"                                      |    Dangeon lvl 2(N) |",
		"                                      |    Dangeon lvl 3(M) |",
		"                                      |    Dangeon lvl 4(L) |",
		"                                      |    Menu(P)          |",
		"                                      `---------------------'",
	}
	lines := CombineColumnsToLines([][]string{mapMenu}, 4)
	FullScreenDrawCentered(lines)

}
