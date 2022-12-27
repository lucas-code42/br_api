package api

import (
	"br_api/models"
	"br_api/routers"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

// GetAllAvailable chama a rota https://brapi.dev/api/available e retorna
func GetAllAvailable() {
	client := &http.Client{}
	req, err := http.NewRequest(
		routers.UrlAvailable.Method,
		routers.UrlAvailable.Url,
		nil,
	)
	if err != nil {
		log.Fatal("Erro ao instanciar nova requisição -> ", err)
	}

	req.Header.Add("accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Erro ao executar a requisição -> ", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Erro ao converter http.body para bytes -> ", err)
	}

	var bodyJSON models.AvailableModels

	if err = json.Unmarshal(body, &bodyJSON); err != nil {
		log.Fatal("Erro ao converter bytes para model -> ", err)
	}

	for i, v := range bodyJSON.Stocks {
		fmt.Printf("Ação nº%d - Cód = %s\n", i, v)
	}

}

// GetCurrency chama uma das rotas da model UrlCurrency e retorna
func GetCurrency() {
	client := &http.Client{}

	req, err := http.NewRequest(
		routers.UrlCurrency.Method,
		routers.UrlCurrency.UrlUsdToBrl,
		nil,
	)
	if err != nil {
		log.Fatal("Erro (GetCurrency) ao instanciar nova requisição -> ", err)
	}

	req.Header.Add("accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Erro (GetCurrency) ao executar request -> ", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Erro (GetCurrency) ao parsear res.body para bytes -> ", err)
	}

	var bodyJSON models.Currency
	if err = json.Unmarshal(body, &bodyJSON); err != nil {
		log.Fatal("Erro (GetCurrency) ao parsear res.body para JSON -> ", err)
	}

	reponseObj := bodyJSON.Currency[0]
	price, err := strconv.ParseFloat(reponseObj.AskPrice, 32)
	if err != nil {
		log.Fatal("Erro (GetCurrency) ao parsear string to float -> ", err)
	}
	fmt.Printf("Cotação %s\nPreço R$ %.2f\n", reponseObj.Name, price)

}

// GetCheaperStocks busca as 100 ações mais baratas e retorna
func GetCheaperStocks() {
	client := &http.Client{}

	req, err := http.NewRequest(
		routers.UrlQuoteList.Method,
		routers.UrlQuoteList.Url,
		nil,
	)
	if err != nil {
		log.Fatal("Erro (GetCheaperStocks) ao instanciar nova requisição -> ", err)
	}

	req.Header.Add("content-type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Erro (GetCheaperStocks) ao executar a requisição ->", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Erro (GetCheaperStocks) ao parsear res.body para bytes ->", err)
	}

	var bodyJSON models.QuoteList
	if err = json.Unmarshal(body, &bodyJSON); err != nil {
		log.Fatal("Erro (GetCheaperStocks) ao parsear body para models.QuoteList ->", err)
	}

	// TODO -> CRIAR ALGORITIMO DE ORDENAÇÃO
	var aux [2]interface{}
	for i, v := range bodyJSON.Stocks {
		if i == 1 {
			aux[0] = v
		}
		if i == 0 {
			aux[1] = v
		}
		fmt.Println("Posição na lista", i, v.Close, v.Name)
	}
	fmt.Println(aux[:]...)

}
