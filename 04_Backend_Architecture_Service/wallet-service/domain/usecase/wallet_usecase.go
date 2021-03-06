package usecase

import (
	"errors"
	"wallet-service/domain/entity"
	domain "wallet-service/domain/repository"

	"github.com/jinzhu/gorm"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated 03/11/2021
type WalletUsecase interface {
	CreateWallet(*entity.Wallet) error
	GetAllWallets() ([]entity.Wallet, error)
	GetByAccountNumber(string) ([]entity.Wallet, error)
	InquiryBalance(string) (*entity.Wallet, error)
	UpdateBalance(string, string, float64) (*entity.Wallet, error)
	IsWalletExist(string) (bool, error)
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
type wallet struct {
	r domain.WalletRepository
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func NewWalletUsecase(r domain.WalletRepository) WalletUsecase {
	return &wallet{r}
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (u *wallet) CreateWallet(wallet *entity.Wallet) error {

	isExist, err := u.IsWalletExist(wallet.ID)
	if err != nil {
		return err
	}

	if isExist {
		return entity.ErrWalletIsExist
	}

	if err := u.r.Create(wallet); err != nil {
		return err
	}

	return nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
func (u *wallet) GetAllWallets() ([]entity.Wallet, error) {
	return u.r.FindAll()
}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
func (u *wallet) GetByAccountNumber(accountNumber string) ([]entity.Wallet, error) {
	return u.r.FindByAccountNumber(accountNumber)
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (u *wallet) InquiryBalance(id string) (*entity.Wallet, error) {
	wallet, err := u.r.FindByID(id)
	if err != nil {
		return nil, err
	}

	return wallet, err
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (u *wallet) UpdateBalance(id, userId string, amount float64) (*entity.Wallet, error) {
	wallet, err := u.InquiryBalance(id)
	if err != nil {
		return nil, err
	}

	// if wallet.UserID != userId {
	// 	return nil, entity.ErrInvalidUser
	// }

	newBalance := wallet.Balance + amount

	// if (newBalance) > wallet.MaxLimit {
	// 	return nil, entity.ErrWalletMaxLimit
	// }

	err = u.r.UpdateBalance(wallet, newBalance)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (u *wallet) IsWalletExist(id string) (bool, error) {
	_, err := u.r.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
