package main

import (
	//"fmt"
)

func inventario() Jogador {
	jogador := Jogador{
		Inventario: []Item{
			ItensDisponiveis[0],
			ItensDisponiveis[1],
			ItensDisponiveis[2],
			ItensDisponiveis[5],
		},
		Pocoes: []Pocao{},
	}
	return jogador
}