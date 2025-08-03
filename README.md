# go-game

#### Como rodar
go run main.go

#### Rodar quando usar mais de um arquivo
go run *.go

## Structs em Go

No Go, uma struct é um tipo de dado que permite agrupar zero ou mais valores de tipos diferentes sob um único nome. Pense nela como uma "coleção de campos" ou uma "caixinha" onde você pode guardar informações relacionadas.

A principal finalidade de uma struct é criar tipos complexos e personalizados que representam entidades do mundo real.

### Exemplo Prático

Imagine que você precisa guardar as informações de um usuário. Em vez de ter várias variáveis separadas como `nome string`, `idade int` e `email string`, você pode criar uma struct `Usuario` para agrupar tudo isso:

```go
package main

import "fmt"

type Usuario struct {
    Nome  string
    Idade int
    Email string
}

func main() {
    // Criando uma nova variável do tipo Usuario
    usuario1 := Usuario{
        Nome:  "Maria",
        Idade: 30,
        Email: "maria@email.com",
    }

    fmt.Println(usuario1.Nome)  // Acessando um campo da struct
    fmt.Println(usuario1.Idade)
}
```

### Por que usar structs?

**Organização e Clareza**: As structs ajudam a organizar o código, tornando-o mais fácil de ler e entender. Se outra pessoa ler seu código, ela saberá imediatamente que `Nome`, `Idade` e `Email` estão relacionados a um `Usuario`.

**Reutilização**: Você pode usar a mesma struct para criar várias variáveis do mesmo tipo, como `usuario1`, `usuario2`, etc., sem precisar redigitar as informações dos campos toda vez.

**Abstração**: Permite modelar conceitos complexos. Você pode ter structs dentro de outras structs, como uma `Endereco` dentro de uma `Usuario`, para representar dados mais aninhados e complexos.

Em resumo, a struct é uma ferramenta fundamental em Go para construir programas mais robustos e bem estruturados, permitindo que você represente e manipule dados de forma lógica e organizada.

### Exemplo do Projeto

Este projeto demonstra o uso de structs com um exemplo simples de `Marcas`:

**struct.go** - Define a struct e funções relacionadas:
```go
type Marcas struct {
    nome string
}

func CriarMarca(nome string) Marcas {
    return Marcas{nome: nome}
}
```

**main.go** - Usa a struct:
```go
func main() {
    marca := CriarMarca("Nike")
    fmt.Println(marca.nome) 
}
```

Este é um exemplo prático de como **separar responsabilidades**: a definição da struct fica em um arquivo separado, e o uso dela fica no `main.go`.

