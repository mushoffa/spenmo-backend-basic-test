package domain

import "card-service/domain/entity"

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
type CardRepository interface {
	Initialize()
	Create(*entity.Card) error
	FindByID(string) (*entity.Card, error)
	FindAll() ([]entity.Card, error)
	UpdateCard(*entity.Card) error
	UpdateStatus(*entity.Card, string) error
}
