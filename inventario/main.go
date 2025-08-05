package main

import "fmt"

// Item representa um item no inventário.
type Item struct {
	Nome  string
	Peso  float64
	Valor int
}

// Inventario representa a mochila do jogador.
type Inventario struct {
	Itens      []Item
	PesoMaximo float64
}

func main() {

	// Criando um novo inventário com peso máximo de 10.0
	mochila := Inventario{
		Itens:      []Item{},
		PesoMaximo: 10.0,
	}

	// Criando alguns itens
	espada := Item{Nome: "Espada de Aço", Peso: 5.5, Valor: 100}
	pocao := Item{Nome: "Poção de Cura", Peso: 0.5, Valor: 25}
	escudo := Item{Nome: "Escudo Grande", Peso: 6.0, Valor: 150}

	// Adicionando itens
	mochila.AdicionarItem(espada)
	mochila.AdicionarItem(pocao)
	mochila.AdicionarItem(escudo) // Este item não deve ser adicionado por ser muito pesado

	// Listando os itens no inventário
	mochila.ListarItens()
}

// PesoAtual calcula o peso total de todos os itens no inventário.
func (i *Inventario) PesoAtual() float64 {
	total := 0.0
	for _, item := range i.Itens {
		total += item.Peso
	}
	return total
}

// AdicionarItem tenta adicionar um item ao inventário.
func (i *Inventario) AdicionarItem(item Item) bool {
	if i.PesoAtual()+item.Peso > i.PesoMaximo {
		fmt.Printf("Erro: %s é muito pesado para o inventário.\n", item.Nome)
		return false
	}
	i.Itens = append(i.Itens, item)
	fmt.Printf("%s adicionado ao inventário.\n", item.Nome)
	return true
}

// ListarItens imprime todos os itens no inventário.
func (i *Inventario) ListarItens() {
	fmt.Println("--- INVENTÁRIO ---")
	if len(i.Itens) == 0 {
		fmt.Println("O inventário está vazio.")
		return
	}
	for _, item := range i.Itens {
		fmt.Printf(" - %s (Peso: %.2f, Valor: %d)\n", item.Nome, item.Peso, item.Valor)
	}
	fmt.Printf("------------------\n")
	fmt.Printf("Peso Total: %.2f / %.2f\n", i.PesoAtual(), i.PesoMaximo)
}
