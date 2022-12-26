package cli

import "fmt"

// Greet cumprimenta o usuário
func Greet() {
	fmt.Println("cli app init")
}

// ShowMenu mostra o menu
func ShowMenu() {
	fmt.Println("Menu:")
	fmt.Println("1 - Mostrar todas as ações disponiveis")
	fmt.Println("2 - Cotação dólar x real.")
	fmt.Println("3 - TOP 100 ações mais baratas")
}
