package main

import (
	"fmt"
	"sort"
	"strings"
)

var Receitas = []Pocao{
	{
		Nome: "Poção de Cura",
		Combinacao: []string{
			"Raiz-de-Cura",
			"Agua-Pura",
		},
		Descricao: "Uma poção que restaura a saúde do aventureiro.",
	},
	{
		Nome: "Poção do Sono",
		Combinacao: []string{
			"Erva-do-Sono",
			"Sais-da-Terra",
		},
		Descricao: "Uma poção que faz com que criaturas adormeçam.",
	},
}

func verificarReceita(combinacao []string) (Pocao, bool) {
	sort.Strings(combinacao) // Ordena para a ordem não importar

	for _, receita := range Receitas {
		receitaOrdenada := make([]string, len(receita.Combinacao))
		copy(receitaOrdenada, receita.Combinacao)
		sort.Strings(receitaOrdenada)

		if strings.Join(combinacao, ",") == strings.Join(receitaOrdenada, ",") {
			return Pocao{Nome: receita.Nome, Descricao: receita.Descricao}, true
		}
	}
	return Pocao{}, false
}

func misturarIngredientes(jogador Jogador) Jogador {
	fmt.Println("\nQuais ingredientes você quer misturar?")
	fmt.Print("Ingrediente 1: ")
	var nomeItem1 string
	fmt.Scanln(&nomeItem1)

	fmt.Print("Ingrediente 2: ")
	var nomeItem2 string
	fmt.Scanln(&nomeItem2)

	// Verificar se o jogador possui os itens
	idx1, ok1 := encontrarItemNoInventario(jogador, nomeItem1)
	idx2, ok2 := encontrarItemNoInventario(jogador, nomeItem2)

	if !ok1 {
		fmt.Printf("Você não tem '%s' no inventário.\n", nomeItem1)
		return jogador
	}
	if !ok2 {
		fmt.Printf("Você não tem '%s' no inventário.\n", nomeItem2)
		return jogador
	}
	if strings.EqualFold(nomeItem1, nomeItem2) {
		fmt.Println("Você não pode misturar um item com ele mesmo.")
		return jogador
	}

	// Verificar se a combinação resulta em uma poção
	combinacao := []string{nomeItem1, nomeItem2}
	pocaoResultante, receitaEncontrada := verificarReceita(combinacao)

	if receitaEncontrada {
		fmt.Printf("\nSucesso! Você criou: %s!\n", pocaoResultante.Nome)

		// Adicionar poção ao inventário
		jogador.Pocoes = append(jogador.Pocoes, pocaoResultante)

		// Remover ingredientes usados do inventário
		// Remove o de maior índice primeiro para não invalidar o menor índice.
		if idx1 > idx2 {
			jogador.Inventario = append(jogador.Inventario[:idx1], jogador.Inventario[idx1+1:]...)
			jogador.Inventario = append(jogador.Inventario[:idx2], jogador.Inventario[idx2+1:]...)
		} else {
			jogador.Inventario = append(jogador.Inventario[:idx2], jogador.Inventario[idx2+1:]...)
			jogador.Inventario = append(jogador.Inventario[:idx1], jogador.Inventario[idx1+1:]...)
		}

	} else {
		fmt.Println("\nA mistura falhou. Nada aconteceu...")
		// Futuramente, pode haver um efeito negativo aqui!
	}

	return jogador
}
