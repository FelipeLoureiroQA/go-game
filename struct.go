package main

// Definindo uma struct simples para marcas
type Marcas struct {
    nome string
}

// Função para criar uma marca com nome
func CriarMarca(nome string) Marcas {
    return Marcas{nome: nome}
}

    // Use a função do struct.go lá em main.go assim:
    // marca := CriarMarca("Nike")


	