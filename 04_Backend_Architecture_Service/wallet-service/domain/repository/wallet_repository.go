package domain

import "wallet-service/domain/entity"

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated 03/11/2021
type WalletRepository interface {
	Initialize()
	Create(*entity.Wallet) error
	FindByID(string) (*entity.Wallet, error)
	FindByAccountNumber(string) ([]entity.Wallet, error)
	FindAll() ([]entity.Wallet, error)
	UpdateBalance(*entity.Wallet, float64) error
	// FindByPhoneNumber(string) (*entity.User, error)
}
