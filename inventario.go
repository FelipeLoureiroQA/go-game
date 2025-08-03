package main

import (
	"fmt"
	"strings"
)

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

func encontrarItemNoInventario(jogador Jogador, nomeItem string) (int, bool) {
	for i, item := range jogador.Inventario {
		if strings.EqualFold(item.Nome, nomeItem) {
			return i, true
		}
	}
	return -1, false
}

func exibirInventario(jogador Jogador) {
	fmt.Println("\n--- SEU INVENTÁRIO ---")
	if len(jogador.Inventario) == 0 {
		fmt.Println("Você não possui ingredientes.")
	} else {
		fmt.Println("Ingredientes:")
		for _, item := range jogador.Inventario {
			fmt.Printf("- %s\n", item.Nome)
		}
	}

	if len(jogador.Pocoes) == 0 {
		fmt.Println("\nVocê não possui poções.")
	} else {
		fmt.Println("\nPoções:")
		for _, pocao := range jogador.Pocoes {
			fmt.Printf("- %s: %s\n", pocao.Nome, pocao.Descricao)
		}
	}
	fmt.Println("----------------------")
}