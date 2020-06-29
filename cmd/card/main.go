package main

import (
	"fmt"
	"github.com/vl-mobitutor/Netology_GO_Task4/pkg/card"
	"github.com/vl-mobitutor/Netology_GO_Task4/pkg/transfer"
)


func main() {

	//Массив карт к выпуску
	myCards := []card.Card {
		{
			Id: 1,
			Issuer: "MasterCard",
			Currency: "RUR",
			Number: "1111 1111 1111 0001",
		},
		{
			Id: 2,
			Issuer: "MasterCard",
			Currency: "RUR",
			Number: "1111 1111 1111 0002",
		},
		{
			Id: 3,
			Issuer: "MasterCard",
			Currency: "RUR",
			Number: "1111 1111 1111 0003",
		},
		{
			Id: 4,
			Issuer: "Visa",
			Currency: "RUR",
			Number: "1111 1111 1111 0004",
		},
		{
			Id: 5,
			Issuer: "Visa",
			Currency: "RUR",
			Number: "1111 1111 1111 0005",
		},
	}

	//Выпускаем карты и кладем на счет по 1000 рублей
	svc := card.NewService("Super Bank")
	for index, newCard := range myCards {
		svc.IssueCard(newCard.Id, newCard.Issuer, newCard.Currency, newCard.Number)
		svc.Cards[index].Balance = 1000_00
	}

	//Показываем карты до перевода
	for _, value := range svc.Cards {
		fmt.Println(*value)
	}


	//Инициируем и настраиваем перевод
	feeSet := []transfer.Fee {
		{
			Description: "С карты банка на карту банка",
			FeePercentage: 0.00,
			FeeMinimum: 0,
		},
		{
			Description: "С карты банка на внешнюю карту",
			FeePercentage: 0.005,
			FeeMinimum: 10_00,
		},
		{
			Description: "С карты банка на внешнюю карту",
			FeePercentage: 0.015,
			FeeMinimum: 30_00,
		},
	}
	trf := transfer.NewService(svc, feeSet)

	for _, value := range trf.Fees {
		fmt.Println(*value)
	}

	//trf.Card2Card("1111 1111 1111 0001", "1111 1111 1111 0002", 300_00)


}
