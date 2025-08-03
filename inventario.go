package main

// criarJogadorInicial inicializa o jogador com itens iniciais.
func criarJogadorInicial() Jogador {
	jogador := Jogador{
		Inventario: []Item{
			ItensDisponiveis[0], // Erva-do-Sono
			ItensDisponiveis[1], // Raiz-de-Cura
			ItensDisponiveis[4], // Sais-da-Terra
			ItensDisponiveis[5], // Agua-Pura
		},
		Pocoes: []Pocao{},
	}
	return jogador
}