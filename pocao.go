package main

// Pocao representa uma poção que pode ser criada.
type Pocao struct {
	Nome       string
	Descricao  string
	Combinacao []string // Nomes dos itens necessários
}