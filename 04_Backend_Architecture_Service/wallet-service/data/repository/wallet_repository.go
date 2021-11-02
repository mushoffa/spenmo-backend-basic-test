package repository

import (
	"time"
	"wallet-service/data/datasource"
	"wallet-service/data/model"
	"wallet-service/domain/entity"
	domain "wallet-service/domain/repository"

	"github.com/google/uuid"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
type wallet struct {
	db datasource.Database
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func NewWalletRepository(db datasource.Database) domain.WalletRepository {
	return &wallet{db}
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (r *wallet) Initialize() {
	r.db.InitializeDatabase()
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (r *wallet) Create(wallet *entity.Wallet) error {
	id := uuid.New().String()
	created := time.Now()

	walletDB := model.WalletDB{
		ID:      id,
		Created: created,
		Name:    wallet.Name,
		Balance: 0,
		UserID:  wallet.UserID,
	}

	if err := r.db.Create(&walletDB); err != nil {
		return err
	}

	wallet.ID = id
	wallet.Created = created

	return nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (r *wallet) FindByID(id string) (*entity.Wallet, error) {
	wallet := entity.Wallet{}

	if err := r.db.FindByID("id", id, &wallet); err != nil {
		return nil, err
	}

	return &wallet, nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (r *wallet) UpdateBalance(wallet *entity.Wallet, balance float64) error {
	// wallet := entity.Wallet{}

	// Need to improve the library
	// if err := r.db.UpdateByID("id", id, "balance", balance, &wallet); err != nil {
	// 	return nil, err
	// }

	updated := time.Now()

	if err := r.db.GetInstance().Model(&model.WalletDB{}).
		Where("id = ?", wallet.ID).
		Where("user_id = ?", wallet.UserID).
		Update("balance", balance).
		Update("updated", updated).Error; err != nil {
		return err
	}

	wallet.Updated = updated
	wallet.Balance = balance

	return nil
}
