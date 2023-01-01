package controller

import "br_api/models"

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
	return newMap
}
