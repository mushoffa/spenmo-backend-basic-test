package service

import (
	"card-service/domain/entity"
	"card-service/domain/usecase"
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/mushoffa/spenmo-proto/protos"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
type card struct {
	protos.UnimplementedCardServiceServer
	u usecase.CardUsecase
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func NewCardService(u usecase.CardUsecase) protos.CardServiceServer {
	return &card{u: u}
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (s *card) CreateCard(ctx context.Context, request *protos.CreateCardRequest) (*protos.CreateCardResponse, error) {
	card := entity.Card{
		PAN:          request.Pan,
		ExpiryDate:   request.ExpiryDate,
		CardType:     request.CardType,
		LimitDaily:   request.LimitDaily,
		LimitMonthly: request.LimitMonthly,
	}

	if err := s.u.CreateCard(&card); err != nil {
		return nil, err
	}

	response := protos.CreateCardResponse{
		Card: &protos.Card{
			Id:           card.ID,
			Created:      card.Created.String(),
			Pan:          card.PAN,
			ExpiryDate:   card.ExpiryDate,
			Type:         card.CardType,
			LimitDaily:   card.LimitDaily,
			LimitMonthly: card.LimitMonthly,
		},
	}

	return &response, nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (s *card) GetAllCards(ctx context.Context, request *empty.Empty) (*protos.GetAllCardsResponse, error) {
	_cards, err := s.u.GetAllCards()
	if err != nil {
		return nil, err
	}

	cards := []*protos.Card{}

	for _, _card := range _cards {

		// limitDaily, _ := fmt.Printf("%f", _card.LimitDaily)

		card := &protos.Card{
			Id:           _card.ID,
			Created:      _card.Created.String(),
			Updated:      _card.Updated.String(),
			Pan:          _card.PAN,
			Name:         _card.Name,
			ExpiryDate:   _card.ExpiryDate,
			LimitDaily:   _card.LimitDaily,
			LimitMonthly: _card.LimitMonthly,
			Type:         _card.CardType,
			Status:       _card.Status,
			WalletId:     _card.WalletID,
			UserId:       _card.UserID,
		}

		cards = append(cards, card)
	}

	response := protos.GetAllCardsResponse{
		Cards: cards,
	}

	return &response, nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (s *card) InquiryCard(ctx context.Context, request *protos.InquiryCardRequest) (*protos.InquiryCardResponse, error) {
	card, err := s.u.InquiryCard(request.Id)
	if err != nil {
		return nil, err
	}

	response := protos.InquiryCardResponse{
		Card: &protos.Card{
			Id:           card.ID,
			Created:      card.Created.String(),
			Updated:      card.Updated.String(),
			Pan:          card.PAN,
			Name:         card.Name,
			ExpiryDate:   card.ExpiryDate,
			LimitDaily:   card.LimitDaily,
			LimitMonthly: card.LimitMonthly,
			Type:         card.CardType,
			Status:       card.Status,
			WalletId:     card.WalletID,
			UserId:       card.UserID,
		},
	}

	return &response, nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (s *card) LinkCard(ctx context.Context, request *protos.LinkCardRequest) (*protos.LinkCardResponse, error) {

	response := protos.LinkCardResponse{}

	return &response, nil
}
