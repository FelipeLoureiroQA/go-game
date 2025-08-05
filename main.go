package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/google/uuid"
)

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

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/go")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	jogador := NewJogador("felipe_teste", criarJogadorInicial().Inventario, []Pocao{})
	err = insertJogador(db, jogador)
	if err != nil {
		panic(err)
	}
	fmt.Println("Jogador de teste inserido com sucesso no banco de dados!")
}
