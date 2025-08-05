package main

import (
	"database/sql"
	"fmt"
)

func IniciarJogo() {
	

	fmt.Println("\nBem-vindo ao mundo da Alquimia!")
	exibirAjuda()

	for {
		fmt.Print("\nO que você quer fazer? ")
		var options int
		fmt.Scanln(&options)

		
jogador := criarJogadorInicial()
		// O 'continue' reinicia o loop, pedindo a próxima ação.
		switch options {
		case 1:
			exibirInventario(jogador)
			continue
		case 2:
			misturarIngredientes(jogador)
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

func exibirMenuPrincipal() {
	fmt.Println("\n--- COMANDOS DISPONÍVEIS ---")
	fmt.Println("1 - Novo Jogo")
	fmt.Println("2 - Carregar Jogo")
	fmt.Println("3 - Sair do jogo.")
	fmt.Println("----------------------------")
}

func exibirAjuda() {
	fmt.Println("\n--- COMANDOS DISPONÍVEIS ---")
	fmt.Println("1 - Inventário: Exibe os itens que você possui.")
	fmt.Println("2 - Misturar: Tenta misturar ingredientes para criar uma poção.")
	fmt.Println("3 - Sair: Sai do jogo.")
	fmt.Println("4 - Ajuda: Exibe esta lista de comandos.")
	fmt.Println("----------------------------")
}

func opcaoMenu(options int) {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/go")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	if options == 1 {
		fmt.Print("\n Novo Jogo ")
		criarJogadorInicial()
		jogador := NewJogador("felipe_teste", criarJogadorInicial().Inventario, []Pocao{})
		err = insertJogador(db, jogador)
		if err != nil {
			panic(err)
		}
		fmt.Println("\nJogador inserido com sucesso no banco de dados!")
	} 
	if options == 2 {
		fmt.Print("\n Carregar Jogo ")
	} else {
		fmt.Print("\n Sair ")
		return
	}
}