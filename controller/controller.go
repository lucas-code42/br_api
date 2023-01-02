package controller

import (
	"br_api/models"
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

// SortStocks ordena as stocks do menor para o maior
func SortStocks(data []models.QuoteListData) []models.QuoteListData {
	cut := len(data) / 10
	data, s1 := DivideList(data, cut)

	cut = len(data) / 10
	data, s2 := DivideList(data, cut)

	cut = len(data) / 10
	data, s3 := DivideList(data, cut)

	cut = len(data) / 10
	data, s4 := DivideList(data, cut)

	c := make(chan []models.QuoteListData)
	go SortListChan(s1, c)
	go SortListChan(s2, c)
	go SortListChan(s3, c)
	go SortListChan(s4, c)

	sorted1 := <-c
	sorted2 := <-c
	sorted3 := <-c
	sorted4 := <-c

	var finalList []models.QuoteListData
	finalList = append(finalList, sorted1...)
	finalList = append(finalList, sorted2...)
	finalList = append(finalList, sorted3...)
	finalList = append(finalList, sorted4...)
	finalList = append(finalList, data...)

	finalList = SortList(finalList)

	return finalList
}

// SortListChan ordena uma lista e devolve a informação no canal
func SortListChan(data []models.QuoteListData, c chan []models.QuoteListData) {
	for j := 0; j < len(data); j++ {
		for i := 0; i < len(data); i++ {
			var temp models.QuoteListData
			currentPosition := i
			nextPostion := i + 1

			if nextPostion == len(data) {
				break
			}

			if data[currentPosition].Close > data[nextPostion].Close {
				temp = data[currentPosition]
				data[currentPosition] = data[nextPostion]
				data[nextPostion] = temp
			}
		}

	}
	c <- data
}

// SortList Ordena a lista e retorna
func SortList(data []models.QuoteListData) []models.QuoteListData {
	for j := 0; j < len(data); j++ {
		for i := 0; i < len(data); i++ {
			var temp models.QuoteListData
			currentPosition := i
			nextPostion := i + 1

			if nextPostion == len(data) {
				break
			}

			if data[currentPosition].Close > data[nextPostion].Close {
				temp = data[currentPosition]
				data[currentPosition] = data[nextPostion]
				data[nextPostion] = temp
			}
		}

	}
	return data
}

// DivideList divide a lista e retorna o recorte mais a sobra da lista.
func DivideList(list []models.QuoteListData, cut int) (tmp, slice []models.QuoteListData) {
	tmp = list[cut:]
	slice = list[:cut]
	return tmp, slice
}

// CreateGroupBySector mapeia cada setor com um array de objts
func CreateGroupBySector(data []models.QuoteListData) map[string][]models.QuoteListData {
	newMap := make(map[string][]models.QuoteListData)

	for _, v := range data {
		var listTemp []models.QuoteListData
		_, ok := newMap[v.Sector]
		if !ok {
			listTemp = append(listTemp, v)
			newMap[v.Sector] = listTemp
		} else {
			newMap[v.Sector] = append(newMap[v.Sector], v)
		}
	}

	newMap = SortStockInSectors(newMap)
	writeXlsx(newMap)
	return newMap
}

// SortStockInSectors ordena as ações da menor para maior dentro do map separado por setores
func SortStockInSectors(data map[string][]models.QuoteListData) map[string][]models.QuoteListData {
	for _, array := range data {
		for j := 0; j < len(array); j++ {
			for i := 0; i < len(array); i++ {
				var temp models.QuoteListData
				currentPosition := i
				nextPostion := i + 1

				if nextPostion == len(array) {
					break
				}

				if array[currentPosition].Close > array[nextPostion].Close {
					temp = array[currentPosition]
					array[currentPosition] = array[nextPostion]
					array[nextPostion] = temp
				}
			}

		}
	}
	return data
}

// writeXlsx escreve dados em um arquivo xlsx, retorna true caso operação de certo
func writeXlsx(data map[string][]models.QuoteListData) bool {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println("Erro ao fechar excelize", err)
		}
	}()

	var dataToWrite []models.QuoteListData
	for _, arr := range data {
		dataToWrite = append(dataToWrite, arr...)
	}

	for i := 0; i < len(dataToWrite); i++ {
		A_cellIndex := fmt.Sprintf("A%s", strconv.FormatInt(int64(i+1), 10))
		B_cellIndex := fmt.Sprintf("B%s", strconv.FormatInt(int64(i+1), 10))
		C_cellIndex := fmt.Sprintf("C%s", strconv.FormatInt(int64(i+1), 10))
		D_cellIndex := fmt.Sprintf("D%s", strconv.FormatInt(int64(i+1), 10))
		
		// fmt.Println(A_cellIndex)
		// fmt.Println(B_cellIndex)
		// fmt.Println(C_cellIndex)

		f.SetCellValue("Sheet1", A_cellIndex, dataToWrite[i].Sector)
		f.SetCellValue("Sheet1", B_cellIndex, dataToWrite[i].Stock)
		f.SetCellValue("Sheet1", C_cellIndex, dataToWrite[i].Name)
		f.SetCellValue("Sheet1", D_cellIndex, dataToWrite[i].Close)
	}

	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println("Erro ao salvar arquivo", err)
		return false
	}

	return true
}
