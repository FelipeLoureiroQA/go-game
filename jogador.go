package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/google/uuid"
)

type Jogador struct {
	ID         string
	Name       string
	Inventario []Item
	Pocoes     []Pocao
}

func NewJogador(name string, inventario []Item, pocoes []Pocao) *Jogador {
	return &Jogador{
		ID:         uuid.New().String(),
		Name:       name,
		Inventario: inventario,
		Pocoes:     pocoes,
	}
}

func insertJogador(db *sql.DB, jogador *Jogador) error {
	// Serializa os slices para JSON antes de inserir no banco
	inventarioJSON, err := json.Marshal(jogador.Inventario)
	if err != nil {
		return fmt.Errorf("falha ao serializar inventario para JSON: %w", err)
	}

	pocoesJSON, err := json.Marshal(jogador.Pocoes)
	if err != nil {
		return fmt.Errorf("falha ao serializar pocoes para JSON: %w", err)
	}

	stmt, err := db.Prepare("insert into jogador (id, name, inventario, pocoes) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(jogador.ID, jogador.Name, inventarioJSON, pocoesJSON)
	if err != nil {
		return err
	}
	return nil

}
