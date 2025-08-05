package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	var options int
	exibirMenuPrincipal()
	fmt.Scanln(&options)
	opcaoMenu(options)

}
