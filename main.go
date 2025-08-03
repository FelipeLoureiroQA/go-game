package main

import (
	"database/sql"

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
	ID 		   string
	Name       string
	Inventario []Item
	Pocoes     []Pocao
}

func NewJogador(name string, inventario []Item, pocoes []Pocao) *Jogador{
	return &Jogador{
		ID: uuid.New().String(),
		Name: name,
		Inventario: inventario,
		Pocoes: pocoes,
	}
}


func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306/go-game")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	jogador := NewJogador("felipe", []Item{}, []Pocao{})
	err = insertJogador(db, jogador)
	if err != nil {
		panic(err)
	}


	IniciarJogo()
}

func insertJogador(db *sql.DB, jogador *Jogador) error {
stmt, err := db.Prepare("INSERT INTO jogadores (id, name, inventario, pocoes) VALUES (?, ?, ?)")
if err != nil {
	return err
}
defer stmt.Close()
_, err = stmt.Exec(jogador.ID, jogador.Name, jogador.Inventario, jogador.Pocoes)
if err!= nil {
	return err
}
return nil

}








