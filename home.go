// Jogo Primeira fase
//Você terá ervas, plantas, raizes, flores, madeiras e produtos alquimicos
// você pode misturar e testar
// Ira criar um poção quando a combinação for correta

// Se a quantidade for incorreta podera criar reações
// buscar receitas para aspoções

package main

import (
	"fmt"
)

func IniciarJogo() {
	jogador := criarJogadorInicial()

	fmt.Println("\nBem-vindo ao mundo da Alquimia!")
	exibirAjuda()

	for {
		fmt.Print("\nO que você quer fazer? ")
		var options int
		fmt.Scanln(&options)

		// O 'continue' reinicia o loop, pedindo a próxima ação.
		switch options {
		case 1:
			exibirInventario(jogador)
			continue
		case 2:
			jogador = misturarIngredientes(jogador)
			continue
		case 3:
			fmt.Println("Obrigado por jogar! Até a próxima.")
			return // Encerra a função IniciarJogo e o programa.
		case 4:
			exibirAjuda()
			continue
		default:
			fmt.Println("Comando inválido. Tente novamente.")
		}
	}
}

func exibirAjuda() {
	fmt.Println("\n--- COMANDOS DISPONÍVEIS ---")
	fmt.Println("1 - Inventário: Exibe os itens que você possui.")
	fmt.Println("2 - Misturar: Tenta misturar ingredientes para criar uma poção.")
	fmt.Println("3 - Sair: Sai do jogo.")
	fmt.Println("4 - Ajuda: Exibe esta lista de comandos.")
	fmt.Println("----------------------------")
}