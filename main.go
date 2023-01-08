package main

import (
	"br_api/api"
	"br_api/cli"
	"br_api/routers"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	routers.StartRouters()


	cli.Greet()
	cli.ShowMenu()
	fmt.Print("Escolha uma opção do menu:")

	var input string
	fmt.Scan(&input)

	fmt.Println("Sua opção digitada foi: ", input)

	switch input {
	case "1":
		api.GetAllAvailable()
	case "2":
		api.GetCurrency()
	case "3":
		api.GetCheaperStocks()
	case "4":
		api.SortStocksBySector()
	default:
		fmt.Println("Nenhuma foi opção digitada, encerrando...")
	}
	os.Exit(0)
}
