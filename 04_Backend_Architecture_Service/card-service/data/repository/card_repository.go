package repository

import (
	"card-service/data/datasource"
	"card-service/data/model"
	"card-service/domain/entity"
	domain "card-service/domain/repository"
	"time"

	"github.com/google/uuid"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
type card struct {
	db datasource.Database
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func NewCardRepository(db datasource.Database) domain.CardRepository {
	return &card{db}
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (r *card) Initialize() {
	r.db.InitializeDatabase()
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (r *card) Create(card *entity.Card) error {
	id := uuid.New().String()
	created := time.Now()

	cardDB := model.CardDB{
		ID:           id,
		Created:      created,
		PAN:          card.PAN,
		ExpiryDate:   card.ExpiryDate,
		LimitDaily:   card.LimitDaily,
		LimitMonthly: card.LimitMonthly,
		CardType:     card.CardType,
		Status:       "UNLINKED",
	}

	if err := r.db.Create(&cardDB); err != nil {
		return err
	}

	card.ID = id
	card.Created = created

	return nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated 04/11/2021
func (r *card) FindByID(id string) (*entity.Card, error) {
	card := entity.Card{}

	if err := r.db.FindByID("pan", id, &card); err != nil {
		return nil, err
	}

	return &card, nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (r *card) FindAll() ([]entity.Card, error) {
	cards := []entity.Card{}

	if err := r.db.FindAll(&cards); err != nil {
		return nil, err
	}

	return cards, nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (r *card) UpdateCard(card *entity.Card) error {
	updated := time.Now()
	card.Status = "LINKED"

	if err := r.db.GetInstance().Model(&model.CardDB{}).
		Where("id = ?", card.ID).
		Update("status", card.Status).
		Update("name", card.Name).
		Update("updated", updated).
		Update("user_id", card.UserID).
		Update("wallet_id", card.WalletID).Error; err != nil {
		return err
	}

	card.Updated = updated

	return nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (r *card) UpdateStatus(card *entity.Card, status string) error {
	// updated := time.Now()

	if err := r.db.UpdateByID("id", card.ID, "status", status, &card); err != nil {
		return err
	}

	return nil
}
