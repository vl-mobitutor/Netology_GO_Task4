package transfer

import (
	"fmt"
	"github.com/vl-mobitutor/Netology_GO_Task4/pkg/card"
	"math"
)

type Service struct {
	CardSvc *card.Service
	Fees []Fee
}

type Fee struct {
	Description string
	FeePercentage float64 //Размер комиссия в % - указывается в виде десятичной дроби - т.е. 1.50% => 0.0150
	FeeMinimum int64 //Минимальная комиссия - указывается в копейках
}

func NewService(cardSvc *card.Service, feeSet []Fee) *Service {
	return &Service {
		CardSvc: cardSvc,
		Fees: feeSet,
	}
}


//Функция перевода по номеру карты
func (s *Service) Card2Card(fromNumber, toNumber string, amount int64) (totalSum int64, ok bool) {
	var myFee int64

	//Проверка на совпадение номеров карты-источника и карты получателя
	if fromNumber == toNumber {
		fmt.Println("Номера карты-источника и карты-получателя совпадают!")
		return amount, false
	}

	//Проверка корректности суммы перевода
	if amount <= 0 {
		fmt.Println("Некорректная сумма перевода!")
		return amount, false
	}


	//Определяем чьи карты
	isFromOur, cardFromIndex := s.CardSvc.FindMyCard(fromNumber)
	isToOur, cardToIndex := s.CardSvc.FindMyCard(toNumber)


	//----------------------------------------------------------------------------------------------
	//Блок, если обе карты "наши"
	if (isFromOur == true) && (isToOur == true) {

		//Установка комиссии
		myFee = int64( math.Round(s.Fees[0].FeePercentage * float64(amount)))
		if myFee < s.Fees[0].FeeMinimum {
			myFee = s.Fees[0].FeeMinimum
		}

		totalSum = amount + myFee //Полная сумма списания с карты-источника

		//Т.к. карты наши, то сразу проводим списание и зачисление
		if totalSum < s.CardSvc.Cards[cardFromIndex].Balance {

			//Списание с карты источника суммы перевода  + комиссия
			s.CardSvc.Cards[cardFromIndex].Balance = s.CardSvc.Cards[cardFromIndex].Balance - totalSum

			//Зачисление на карту-получатель суммы перевода (без комиссии)
			s.CardSvc.Cards[cardToIndex].Balance = s.CardSvc.Cards[cardToIndex].Balance + amount

			fmt.Println("Тип перевода - внутрибанковский платеж")
			ok = true

		} else {
			fmt.Printf("На карте %s недостаточно средств для перевода! \n", fromNumber)
			ok = false
		}

	}


	//-----------------------------------------------------------------------------------------------
	//Блок, если с "нашей" карты на внешнюю карту
	if (isFromOur == true) && (isToOur == false) {

		//Установка комиссии
		myFee = int64( math.Round(s.Fees[1].FeePercentage * float64(amount)))
		if myFee < s.Fees[1].FeeMinimum {
			myFee = s.Fees[1].FeeMinimum
		}

		totalSum = amount + myFee //Полная сумма списания с карты-источника

		//Проводим только списание с нашей карты
		if totalSum < s.CardSvc.Cards[cardFromIndex].Balance {

			//Списание с карты источника суммы перевода  + комиссия
			s.CardSvc.Cards[cardFromIndex].Balance = s.CardSvc.Cards[cardFromIndex].Balance - totalSum

			fmt.Println("Тип перевода - с карты банка на внешнюю карту")
			ok = true

		} else {
			fmt.Printf("На карте %s недостаточно средств для перевода! \n", fromNumber)
			ok = false
		}

	}


	//----------------------------------------------------------------------------------------------
	//Блок, если с внешней карты на карту банка
	if (isFromOur == false) && (isToOur == true) {

		//Установка комиссии
		myFee = int64( math.Round(s.Fees[2].FeePercentage * float64(amount)))
		if myFee < s.Fees[2].FeeMinimum {
			myFee = s.Fees[2].FeeMinimum
		}

		totalSum = amount + myFee //Полная сумма списания с карты-источника

		//Зачисление на карту-получатель суммы перевода (без комиссии)
		s.CardSvc.Cards[cardToIndex].Balance = s.CardSvc.Cards[cardToIndex].Balance + amount

		fmt.Println("Тип перевода - с внешней карты на карту банка")
		ok = true

	}

	//----------------------------------------------------------------------------------------------
	//Блок, если с внешней карты на внешнюю банка
	if (isFromOur == false) && (isToOur == false) {

		//Установка комиссии
		myFee = int64( math.Round(s.Fees[3].FeePercentage * float64(amount)))
		if myFee < s.Fees[3].FeeMinimum {
			myFee = s.Fees[3].FeeMinimum
		}

		totalSum = amount + myFee //Полная сумма списания с карты-источника

		fmt.Println("Тип перевода - с внешней карты на внешнюю карту")
		ok = true

	}

	return totalSum, ok
}
