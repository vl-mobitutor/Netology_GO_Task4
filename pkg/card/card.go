package card

type Service struct {
	BankName string
	Cards []*Card
}

type Card struct {
	Id int64
	Issuer string
	Balance int64
	Currency string
	Number string
}

func NewService(bankName string) *Service {
	return &Service{BankName: bankName}
}

func (s *Service) IssueCard(id int64, issuer, currency, number string) *Card {
	card := &Card {
		Id: id,
		Issuer: issuer,
		Balance: 0,
		Currency: currency,
		Number: number,
	}
	s.Cards = append(s.Cards, card)
	return card
}

//Поиск номера карты в массиве собственных карт банка
func (s *Service) FindMyCard(cardNumber string) (result bool, cardIndex int){

	for index, value := range s.Cards {
		if value.Number == cardNumber {
			result = true
			cardIndex = index
			break
		}
		result = false
		cardIndex = 0
	}

	return result, cardIndex
}