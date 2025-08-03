package main

// Item representa um ingrediente ou qualquer item no inventário.
type Item struct {
	Nome string
}

// Pocao representa uma poção que pode ser criada.
type Pocao struct {
	Nome       string
	Descricao  string
	Combinacao []string // Nomes dos itens necessários
}

// Jogador representa o estado do jogador no jogo.
type Jogador struct {
	Inventario []Item
	Pocoes     []Pocao
}
