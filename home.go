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
	fmt.Println("Comandos disponíveis:")
	fmt.Println("1 - Inventario': Exibe os itens que você possui.")
	fmt.Println("2 - Misturar': Tenta misturar ingredientes para criar uma poção.")
	fmt.Println("3 - Sair': Sai do jogo.")
	fmt.Println("4 - Ajuda': Exibe esta lista de comandos.")

	for {
		fmt.Print("\nO que você quer fazer? ")
		var options int
		fmt.Scanln(&options)

		switch options {
		case 1:
			exibirInventario(jogador)
		case 2:
			jogador = misturarIngredientes(jogador)
		case 3:
			fmt.Println("Obrigado por jogar!")
			return
		case 4:
			fmt.Println("Comandos disponíveis:")
			fmt.Println("1 - Inventario': Exibe os itens que você possui.")
			fmt.Println("2 - Misturar': Tenta misturar ingredientes para criar uma poção.")
			fmt.Println("3 - Sair': Sai do jogo.")
			fmt.Println("4 - Ajuda': Exibe esta lista de comandos.")
			
		default:
			fmt.Println("Comando inválido. Tente novamente.")
		}
	}
 }
