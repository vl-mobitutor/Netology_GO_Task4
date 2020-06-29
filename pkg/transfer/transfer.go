package transfer

import "github.com/vl-mobitutor/Netology_GO_Task4/pkg/card"

type Service struct {
	CardSvc *card.Service
	Fees []Fee
}

type Fee struct {
	Description string
	FeePercentage float32
	FeeMinimum int64
}

func NewService(cardSvc *card.Service, feeSet []Fee) *Service {
	return &Service {
		CardSvc: cardSvc,
		Fees: feeSet,
	}
}

func (s *Service) Card2Card(fromNumber, toNumber string, amount int) (total int, ok bool) {


	return total, ok
}
