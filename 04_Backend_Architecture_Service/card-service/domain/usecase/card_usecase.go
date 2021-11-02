package usecase

import (
	"card-service/domain/entity"
	domain "card-service/domain/repository"
	"errors"

	"github.com/jinzhu/gorm"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
type CardUsecase interface {
	CreateCard(*entity.Card) error
	GetAllCards() ([]entity.Card, error)
	InquiryCard(string) (*entity.Card, error)
	LinkCard(*entity.Card) error
	IsCardExist(string) (bool, *entity.Card, error)
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
type card struct {
	r domain.CardRepository
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func NewCardUsecase(r domain.CardRepository) CardUsecase {
	return &card{r}
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (u *card) CreateCard(card *entity.Card) error {

	if err := card.Validate(); err != nil {
		return err
	}

	isExist, _, err := u.IsCardExist(card.ID)
	if err != nil {
		return err
	}

	if isExist {
		return entity.ErrCardIsExist
	}

	if err := u.r.Create(card); err != nil {
		return err
	}

	return nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (u *card) GetAllCards() ([]entity.Card, error) {
	return u.r.FindAll()
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (u *card) InquiryCard(id string) (*entity.Card, error) {
	card, err := u.r.FindByID(id)
	if err != nil {
		return nil, err
	}

	return card, nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (u *card) LinkCard(card *entity.Card) error {

	return u.r.UpdateCard(card)
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (u *card) IsCardExist(id string) (bool, *entity.Card, error) {
	card, err := u.r.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil, nil
		}

		return false, nil, err
	}

	return true, card, nil
}
